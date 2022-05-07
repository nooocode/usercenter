package model

import (
	"errors"

	"github.com/nooocode/pkg/model"
	apipb "github.com/nooocode/usercenter/api"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Menu struct {
	model.Model
	Level       uint32           `json:"level"`
	ParentID    string           `json:"parentID" gorm:"comment:父菜单ID"`
	Path        string           `json:"path" gorm:"size:200;comment:路由path"`
	Name        string           `json:"name" gorm:"size:100;comment:路由name"`
	Hidden      bool             `json:"hidden" gorm:"comment:是否在列表隐藏"`
	Component   string           `json:"component" gorm:"size:200;comment:对应前端文件路径"`
	Sort        int32            `json:"sort" gorm:"comment:排序标记"`
	Cache       bool             `json:"cache" gorm:"comment:是否缓存"`
	DefaultMenu bool             `json:"defaultMenu" gorm:"comment:是否是基础路由（开发中）"`
	Title       string           `json:"title" gorm:"size:100;comment:菜单名"`
	Icon        string           `json:"icon" gorm:"size:100;comment:菜单图标"`
	CloseTab    bool             `json:"closeTab" gorm:"comment:自动关闭tab"`
	Children    []*Menu          `json:"children" gorm:"-"`
	Parameters  []*MenuParameter `json:"parameters"`
	MenuFuncs   []*MenuFunc      `json:"menuFuncs"`
}

type MenuParameter struct {
	model.Model
	MenuID string `json:"menuID" gorm:"index"`
	Type   string `json:"type" gorm:"size:50;comment:地址栏携带参数为params还是query"`
	Key    string `json:"key" gorm:"size:100;comment:地址栏携带参数的key"`
	Value  string `json:"value" gorm:"size:200;comment:地址栏携带参数的值"`
}

type MenuFunc struct {
	model.Model
	MenuID       string        `json:"menuID" gorm:"index"`
	Name         string        `json:"name" gorm:"size:100;comment:功能名称"`
	Title        string        `json:"title" gorm:"size:100;comment:显示名称"`
	Hidden       bool          `json:"hidden" gorm:"comment:是否隐藏"`
	MenuFuncApis []MenuFuncApi `json:"menuFuncApis"`
}

type MenuFuncApi struct {
	model.Model
	MenuFuncID string `json:"menuFuncID" gorm:"index"`
	APIID      string `json:"apiID" gorm:"column:api_id"`
	API        *API   `json:"apiInfo"`
}

//@author: [guoxf](https://github.com/guoxf)
//@function: AddMenu
//@description: 添加菜单
//@param: menu Menu
//@return: err error
func AddMenu(menu *Menu) error {
	if menu.ParentID != "" {
		parent := &Menu{}
		err := dbClient.DB().Where("id=?", menu.ParentID).First(&parent).Error
		if err != nil {
			return err
		}
		menu.Level = parent.Level + 1
	}

	return dbClient.DB().Create(menu).Error
	// duplication, err := dbClient.CreateWithCheckDuplication(menu, "name = ?", menu.Name)
	// if err != nil {
	// 	return err
	// }
	// if duplication {
	// 	return errors.New("存在重复name，请修改name")
	// }
	// return nil
}

//@author: [guoxf](https://github.com/guoxf)
//@function: DeleteMenu
//@description: 删除菜单
//@param: id uint
//@return: err error
func DeleteMenu(id string) (err error) {
	return dbClient.DB().Transaction(func(tx *gorm.DB) error {
		//判断次菜单是否存在子菜单
		duplication, err := dbClient.CheckDuplication(tx.Model(&Menu{}), "parent_id = ?", id)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("此菜单存在子菜单不可删除")
		}
		menu, err := GetMenuByID(id)
		if err != nil {
			return err
		}

		// 删除和角色关联
		err = tx.Unscoped().Delete(&RoleMenu{}, "menu_id=?", id).Error
		if err != nil {
			return err
		}

		err = tx.Unscoped().Delete(&MenuParameter{}, "menu_id = ?", id).Error
		if err != nil {
			return err
		}

		var menuFuncIDs []string
		for _, menuFunc := range menu.MenuFuncs {
			menuFuncIDs = append(menuFuncIDs, menuFunc.ID)
		}
		if len(menuFuncIDs) > 0 {
			err = tx.Unscoped().Delete(&MenuFuncApi{}, "menu_func_id in ?", menuFuncIDs).Error
			if err != nil {
				return err
			}
		}
		err = tx.Unscoped().Delete(&MenuFunc{}, "menu_id = ?", id).Error
		if err != nil {
			return err
		}

		err = tx.Unscoped().Delete(&Menu{}, "id = ?", id).Error
		if err != nil {
			return err
		}

		return err
	})

}

//@author: [guoxf](https://github.com/guoxf)
//@function: UpdateMenu
//@description: 更新菜单
//@param: menu Menu
//@return:err error
func UpdateMenu(menu *Menu) (err error) {

	err = dbClient.DB().Transaction(func(tx *gorm.DB) error {
		// duplication, err := dbClient.CheckDuplication(tx.Model(&Menu{}), "id <> ? and name=?", menu.ID, menu.Name)
		// if err != nil {
		// 	return err
		// }
		// if duplication {
		// 	return errors.New("存在重复name，请修改name")
		// }
		if menu.ParentID != "" {
			parent := &Menu{}
			err := tx.Where("id=?", menu.ParentID).First(&parent).Error
			if err != nil {
				return err
			}
			menu.Level = parent.Level + 1
		}
		oldMenu := &Menu{}
		err = tx.Preload("MenuFuncs.MenuFuncApis.API").Preload(clause.Associations).Where("id = ?", menu.ID).First(oldMenu).Error
		if err != nil {
			return err
		}
		var deleteAPIs []string
		var deleteMenuFuncs []string
		var deleteParams []string
		for _, oldMenuFunc := range oldMenu.MenuFuncs {
			flag := false
			for _, newMenuFunc := range menu.MenuFuncs {
				if newMenuFunc.ID == oldMenuFunc.ID {
					flag = true
					for _, oldAPI := range oldMenuFunc.MenuFuncApis {
						apiFlag := false
						for _, newAPI := range newMenuFunc.MenuFuncApis {
							if newAPI.ID == oldAPI.ID {
								apiFlag = true
							}
						}
						if !apiFlag {
							deleteAPIs = append(deleteAPIs, oldAPI.ID)
						}
					}
				}
			}
			if !flag {
				deleteMenuFuncs = append(deleteMenuFuncs, oldMenuFunc.ID)
			}
		}

		for _, oldParam := range oldMenu.Parameters {
			flag := false
			for _, newParam := range menu.Parameters {
				if newParam.ID == oldParam.ID {
					flag = true
				}
			}
			if !flag {
				deleteParams = append(deleteParams, oldParam.ID)
			}
		}
		if len(deleteParams) > 0 {
			err = tx.Unscoped().Delete(&MenuParameter{}, "id in ?", deleteParams).Error
			if err != nil {
				return err
			}
		}
		if len(deleteAPIs) > 0 {
			err = tx.Unscoped().Delete(&MenuFuncApi{}, "id in ?", deleteAPIs).Error
			if err != nil {
				return err
			}
		}
		if len(deleteMenuFuncs) > 0 {
			err = tx.Unscoped().Delete(&MenuFuncApi{}, "menu_func_id in ?", deleteMenuFuncs).Error
			if err != nil {
				return err
			}
			err = tx.Unscoped().Delete(&MenuFunc{}, "id in ?", deleteMenuFuncs).Error
			if err != nil {
				return err
			}
		}
		return tx.Session(&gorm.Session{FullSaveAssociations: true}).Omit("created_at").Save(menu).Error
	})
	return err
}

//@author: [guoxf](https://github.com/guoxf)
//@function: GetMenuByID
//@description: 返回当前选中menu
//@param: id uint
//@return: err error, menu Menu
func GetMenuByID(id string) (*Menu, error) {
	menu := &Menu{}
	err := dbClient.DB().Preload("MenuFuncs.MenuFuncApis.API").Preload(clause.Associations).Where("id = ?", id).First(menu).Error
	return menu, err
}

//@author: [guoxf](https://github.com/guoxf)
//@function: GetMenuInfoList
//@description: 分页查询Menu
//@param: api Menu, info PageInfo, order string, desc bool
//@return: list []*Menu, total int64 , err error
func QueryMenu(req *apipb.QueryMenuRequest, resp *apipb.QueryMenuResponse, preload bool) {
	db := dbClient.DB().Model(&Menu{})

	if req.Name != "" {
		db = db.Where("name LIKE ?", "%"+req.Name+"%")
	}

	if req.Path != "" {
		db = db.Where("path LIKE ?", "%"+req.Path+"%")
	}

	if req.Title != "" {
		db = db.Where("title LIKE ?", "%"+req.Title+"%")
	}

	if req.ParentID != "" {
		db = db.Where("`parent_id` = ?", req.ParentID)
	}

	if req.Level != 0 {
		db = db.Where("`level` = ?", req.Level)
	}

	if len(req.Ids) > 0 {
		db = db.Where("id in ?", req.Ids)
	}

	OrderStr := "`sort`"
	if req.OrderField != "" {
		if req.Desc {
			OrderStr = req.OrderField + " desc"
		} else {
			OrderStr = req.OrderField
		}
	}
	var err error
	var list []*Menu
	if preload {
		resp.Records, resp.Pages, err = dbClient.PageQueryWithPreload(db, req.PageSize, req.PageIndex, OrderStr, []string{"MenuFuncs.MenuFuncApis", "Parameters", clause.Associations}, &list)
	} else {
		resp.Records, resp.Pages, err = dbClient.PageQuery(db, req.PageSize, req.PageIndex, OrderStr, &list)
	}
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = MenusToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllMenus() (menus []*Menu, err error) {
	err = dbClient.DB().Find(&menus).Error
	return
}
