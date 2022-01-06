package model

import (
	"errors"
	"fmt"
	"strings"

	"github.com/nooocode/pkg/model"
	apipb "github.com/nooocode/usercenter/api"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Role struct {
	model.TenantModel
	Name          string      `json:"name" gorm:"size:100;comment:角色名"`
	ParentID      string      `json:"parentID" gorm:"comment:父角色ID"`
	Children      []*Role     `json:"children" gorm:"-"`
	RoleMenus     []*RoleMenu `json:"roleMenus"`
	DefaultRouter string      `json:"defaultRouter" gorm:"size:100;comment:默认菜单;default:dashboard"`
	Description   string      `json:"description" gorm:"size:200;"`
	CanDel        bool        `json:"canDel" gorm:"default:1"`
	Tenant        *Tenant     `json:"tenant"`
}

type RoleMenu struct {
	model.Model
	RoleID string `json:"roleID" gorm:"index;comment:角色ID"`
	MenuID string `json:"menuID" gorm:"index;comment:菜单ID"`
	Funcs  string `json:"funcs" gorm:"size:500;comment:功能名称,多个以逗号隔开"`
	Menu   *Menu  `json:"menu"`
}

//@author: [guoxf](https://github.com/guoxf)
//@function: CreateRole
//@description: 创建一个角色
//@param: newRole Role
//@return: err error, role Role
func CreateRole(newRole *Role) error {
	newRole.CanDel = true
	duplication, err := dbClient.CreateWithCheckDuplication(newRole, "id = ?", newRole.ID)
	if err != nil {
		return err
	}
	if duplication {
		return errors.New("存在相同角色id")
	}
	err = updateRoleAuth(newRole.ID)
	if err != nil {
		return err
	}

	return nil
}

func updateRoleAuth(id string) error {
	roleDetail, err := GetFullRoleByID(id)
	if err != nil {
		return err
	}
	roleID := roleDetail.ID
	var newRules = make(map[string]*CasbinRule)
	for _, menu := range roleDetail.RoleMenus {
		funcs := strings.Split(menu.Funcs, ",")
		for _, fn := range menu.Menu.MenuFuncs {
			flag := false
			for _, f := range funcs {
				if fn.Name == f {
					flag = true
					break
				}
			}
			if !flag {
				continue
			}
			for _, api := range fn.MenuFuncApis {
				if !api.API.Enable {
					continue
				}
				checkAuth := "true"
				if !api.API.CheckAuth {
					checkAuth = "false"
				}
				key := fmt.Sprintf("p-%v-%v-%v-%v", roleID, api.API.Path, api.API.Method, checkAuth)
				_, ok := newRules[key]
				if ok {
					continue
				}
				newRules[key] = &CasbinRule{
					Ptype:     "p",
					RoleID:    roleID,
					Path:      api.API.Path,
					Method:    api.API.Method,
					CheckAuth: checkAuth,
				}
			}
		}
	}
	var list []*CasbinRule
	for _, r := range newRules {
		list = append(list, r)
	}
	_, err = ClearCasbin(0, roleID)
	if err != nil {
		fmt.Println("ClearCasbin error:", err)
	}
	fmt.Println("=======>newRules", list)
	if len(list) > 0 {
		err = UpdateCasbin(roleID, list)
		if err != nil {
			return err
		}
	}

	return nil
}

//@author: [guoxf](https://github.com/guoxf)
//@function: CopyRole
//@description: 复制一个角色
//@param: copyInfo RoleCopyResponse
//@return: err error, role Role
func CopyRole(copyInfo RoleCopyResponse) (*Role, error) {
	err := dbClient.DB().Transaction(func(tx *gorm.DB) error {
		duplication, err := dbClient.CheckDuplication(tx.Model(&Role{}), "id = ?", copyInfo.Role.ID)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同角色id")
		}
		copyInfo.Role.Children = []*Role{}
		menus, err := GetMenuRole(copyInfo.OldRoleID)
		if err != nil {
			return err
		}
		var roleMenus []*RoleMenu
		for _, v := range menus {
			roleMenus = append(roleMenus, &RoleMenu{
				MenuID: v.MenuID,
				RoleID: copyInfo.Role.ID,
				Funcs:  v.Funcs,
			})
		}
		copyInfo.Role.RoleMenus = roleMenus

		err = dbClient.DB().Create(&copyInfo.Role).Error
		if err != nil {
			return err
		}

		roleID := copyInfo.Role.ID
		rules := GetPolicyPathByRoleID(copyInfo.OldRoleID)
		for i := range rules {
			rules[i].RoleID = roleID
		}
		return UpdateCasbin(roleID, rules)

	})

	return &copyInfo.Role, err
}

//@author: [guoxf](https://github.com/guoxf)
//@function: UpdateRole
//@description: 更改一个角色
//@param: newRole Role
//@return:err error, role Role
func UpdateRole(newRole *Role) error {
	return dbClient.DB().Transaction(func(tx *gorm.DB) error {
		oldRole := &Role{}
		err := tx.Preload("RoleMenus").Preload(clause.Associations).Where("id = ?", newRole.ID).First(oldRole).Error
		if err != nil {
			return err
		}
		var deleteRoleMenu []string
		for _, oldRoleMenu := range oldRole.RoleMenus {
			flag := false
			for _, newRoleMenu := range newRole.RoleMenus {
				if newRoleMenu.ID == oldRoleMenu.ID {
					flag = true
				}
			}
			if !flag {
				deleteRoleMenu = append(deleteRoleMenu, oldRoleMenu.ID)
			}
		}
		if len(deleteRoleMenu) > 0 {
			err = tx.Unscoped().Delete(&RoleMenu{}, "id in ?", deleteRoleMenu).Error
			if err != nil {
				return err
			}
		}
		for _, m := range newRole.RoleMenus {
			m.RoleID = newRole.ID
			err = tx.Save(m).Error
			if err != nil {
				return err
			}
		}

		err = tx.Session(&gorm.Session{FullSaveAssociations: false}).Omit("can_del", "created_at").Updates(newRole).Error
		if err != nil {
			return err
		}

		err = updateRoleAuth(newRole.ID)
		return err
	})
}

//@author: [guoxf](https://github.com/guoxf)
//@function: DeleteRole
//@description: 删除角色
//@param: auth *Role
//@return: err error
func DeleteRole(roleID string) (err error) {
	return dbClient.DB().Transaction(func(tx *gorm.DB) error {
		duplication, err := dbClient.CheckDuplication(tx.Model(&UserRole{}), "role_id = ?", roleID)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("此角色有用户正在使用禁止删除")
		}

		duplication, err = dbClient.CheckDuplication(tx.Model(&Role{}), "parent_id = ?", roleID)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("此角色存在子角色不允许删除")
		}

		oldRole, err := GetRoleByID(roleID)
		if err != nil {
			return err
		}
		if !oldRole.CanDel {
			return errors.New("此角色不允许删除")
		}

		err = tx.Unscoped().Delete(&RoleMenu{}, "role_id=?", roleID).Error
		if err != nil {
			return err
		}
		err = tx.Unscoped().Delete(&Role{}, "id=?", roleID).Error
		if err != nil {
			return err
		}

		ClearCasbin(0, fmt.Sprint(roleID))
		return err
	})
}

func QueryRole(req *apipb.QueryRoleRequest, resp *apipb.QueryRoleResponse) {
	db := dbClient.DB().Model(&Role{})

	if req.Name != "" {
		db = db.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.TenantID != "" {
		db = db.Where("tenant_id = ?", req.TenantID)
	}

	OrderStr := "`name`"
	if req.OrderField != "" {
		if req.Desc {
			OrderStr = req.OrderField + " desc"
		} else {
			OrderStr = req.OrderField
		}
	}
	var err error
	var roles []*Role
	resp.Records, resp.Pages, err = dbClient.PageQueryWithPreload(db, req.PageSize, req.PageIndex, OrderStr, []string{"Tenant"}, &roles)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = RolesToPB(roles)
	}
}

//@author: [guoxf](https://github.com/guoxf)
//@function: findChildrenRole
//@description: 查询子角色
//@param: authority *Role
//@return: err error
func findChildrenRole(authority *Role) (err error) {
	err = dbClient.DB().Preload("RoleMenus").Where("parent_id = ?", authority.ID).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = findChildrenRole(authority.Children[k])
		}
	}
	return err
}

//@author: [guoxf](https://github.com/guoxf)
//@function: getMenuTreeMap
//@description: 获取路由总树map
//@param: roleID string
//@return: err error, treeMap map[string][]Menu
func getMenuTreeMap(roleID string) (treeMap map[string][]*Menu, err error) {
	var allMenus []*Menu
	treeMap = make(map[string][]*Menu)
	err = dbClient.DB().Where("role_id = ?", roleID).Order("sort").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentID] = append(treeMap[v.ParentID], v)
	}
	return treeMap, err
}

//@author: [guoxf](https://github.com/guoxf)
//@function: GetMenuTree
//@description: 获取动态菜单树
//@param: roleID string
//@return: err error, menus []RoleMenu
func GetMenuTree(roleID string) (menus []*Menu, err error) {
	menuTree, err := getMenuTreeMap(roleID)
	menus = menuTree[""]
	for i := 0; i < len(menus); i++ {
		err = getChildrenList(menus[i], menuTree)
	}
	return menus, err
}

//@author: [guoxf](https://github.com/guoxf)
//@function: getChildrenList
//@description: 获取子菜单
//@param: menu *Menu, treeMap map[string][]Menu
//@return: err error
func getChildrenList(menu *Menu, treeMap map[string][]*Menu) (err error) {
	menu.Children = treeMap[menu.ID]
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(menu.Children[i], treeMap)
	}
	return err
}

//@author: [guoxf](https://github.com/guoxf)
//@function: GetInfoList
//@description: 获取路由分页
//@return: err error, list interface{}, total int64
func GetInfoList() (list []*Menu, total int64, err error) {
	var menuList []*Menu
	treeMap, err := getBaseMenuTreeMap()
	menuList = treeMap[""]
	for i := 0; i < len(menuList); i++ {
		err = getBaseChildrenList(menuList[i], treeMap)
	}
	return menuList, total, err
}

//@author: [guoxf](https://github.com/guoxf)
//@function: getBaseChildrenList
//@description: 获取菜单的子菜单
//@param: menu *Menu, treeMap map[string][]Menu
//@return: err error
func getBaseChildrenList(menu *Menu, treeMap map[string][]*Menu) (err error) {
	menu.Children = treeMap[menu.ID]
	for i := 0; i < len(menu.Children); i++ {
		err = getBaseChildrenList(menu.Children[i], treeMap)
	}
	return err
}

//@author: [guoxf](https://github.com/guoxf)
//@function: getBaseMenuTreeMap
//@description: 获取路由总树map
//@return: err error, treeMap map[string][]Menu
func getBaseMenuTreeMap() (treeMap map[string][]*Menu, err error) {
	var allMenus []*Menu
	treeMap = make(map[string][]*Menu)
	err = dbClient.DB().Order("sort").Preload("MenuFuncs").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentID] = append(treeMap[v.ParentID], v)
	}
	return treeMap, err
}

//@author: [guoxf](https://github.com/guoxf)
//@function: GetBaseMenuTree
//@description: 获取基础路由树
//@return: err error, menus []Menu
func GetBaseMenuTree() (menus []*Menu, err error) {
	treeMap, err := getBaseMenuTreeMap()
	menus = treeMap[""]
	for i := 0; i < len(menus); i++ {
		err = getBaseChildrenList(menus[i], treeMap)
	}
	return menus, err
}

//@author: [guoxf](https://github.com/guoxf)
//@function: GetMenuRole
//@description: 查看当前角色权限树
//@param: info *request.GetAuthorityId
//@return: err error, menus []RoleMenu
func GetMenuRole(roleID string) (menus []RoleMenu, err error) {
	err = dbClient.DB().Where("role_id = ? ", roleID).Order("sort").Find(&menus).Error
	return menus, err
}

type RoleResponse struct {
	Role Role `json:"role"`
}

type RoleCopyResponse struct {
	Role      Role   `json:"role"`
	OldRoleID string `json:"oldRoleID"`
}

func GetRoleByID(id string) (*Role, error) {
	role := &Role{}
	err := dbClient.DB().Preload("RoleMenus").Preload(clause.Associations).Where("id = ?", id).First(role).Error
	return role, err
}

func GetFullRoleByID(id string) (*Role, error) {
	role := &Role{}
	err := dbClient.DB().Preload("RoleMenus.Menu.MenuFuncs.MenuFuncApis.API").Preload(clause.Associations).Where("id = ?", id).First(role).Error
	return role, err
}

func GetAllRole(tenantID string) (roles []*Role, err error) {
	if tenantID != "" {
		err = dbClient.DB().Find(&roles, "tenant_id=?", tenantID).Error
	} else {
		err = dbClient.DB().Find(&roles).Error
	}
	return
}

func StatisticRoleCount(tenantID string) (int64, error) {
	db := dbClient.DB().Model(&Role{})
	if tenantID != "" {
		db = db.Where("tenant_id = ?", tenantID)
	}
	var count int64
	err := db.Count(&count).Error
	return count, err
}
