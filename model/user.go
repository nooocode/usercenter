package model

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/dgrijalva/jwt-go/v4"
	scrypt "github.com/elithrar/simple-scrypt"
	"github.com/nooocode/pkg/model"
	"github.com/nooocode/pkg/utils/log"
	apipb "github.com/nooocode/usercenter/api"
	"github.com/nooocode/usercenter/model/token"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	model.TenantModel
	//账号信息
	UserName      string      `json:"userName"  validate:"required" gorm:"size:50;index;comment:用户登录名"`
	Password      string      `json:"password"  gorm:"size:200;comment:用户登录密码"`
	Nickname      string      `json:"nickname"  validate:"required" gorm:"size:100;index;default:未设置;comment:用户昵称" `
	UserRoles     []*UserRole `json:"userRoles"`
	RoleIDs       []string    `json:"roleIDs" gorm:"-"`
	WechatUnionID string      `json:"wechatUnionID" gorm:"size:36;index;comment:微信UionID"`
	WechatOpenID  string      `json:"wechatOpenID" gorm:"size:36;index;comment:微信OpenID"`
	//启用账号后可以登录系统
	Enable bool `json:"enable" gorm:"index"`
	// 用户名/密码错误次数
	ErrNumber int32 `json:"errNumber"`
	// 随着密码/用户名错误次数增加，用户被锁定的时间也就越长，一天最多错误5次
	LockedExpired int64 `json:"lockedExpired"`
	// 强制修改密码
	ForceChangePwd bool `json:"forceChangePwd"`
	// 用户过期时间
	Expired                int64 `json:"expired"`
	ChangePwdErrNum        int32 `json:"changePwdErrNum"`
	ChangePwdLockedExpired int64 `json:"changePwdLockedExpired"`

	CanDel bool `json:"canDel"`

	//基本信息
	Email       string `json:"email" gorm:"size:100;"`
	Mobile      string `json:"mobile" gorm:"size:20;index;comment:手机号"`
	IDCard      string `json:"idCard"  gorm:"size:18;index;comment:身份证号"`
	Avatar      string `json:"avatar" gorm:"size:200;comment:用户头像"`
	EID         string `json:"eid" gorm:"size:50;"`
	Title       string `json:"title" gorm:"size:100;comment:职位"`
	Description string `json:"description" gorm:"size:200;"`
	RealName    string `json:"realName" gorm:"index;size:50;"`
	Gender      bool   `json:"gender"`

	Country  string `json:"country" gorm:"size:100;"`  //国家
	Province string `json:"province" gorm:"size:100;"` //省份
	City     string `json:"city" gorm:"size:100;"`     //城市
	County   string `json:"county" gorm:"size:100;"`   //区县
	Birthday int64  `json:"birthday"`                  //公历出生日期包含时分
}

func (u User) GetRoleIDs() []string {
	var roleIDs []string
	for _, role := range u.UserRoles {
		roleIDs = append(roleIDs, role.RoleID)
	}
	return roleIDs
}

type UserRole struct {
	model.Model
	UserID string `json:"userID" gorm:"index"`
	RoleID string `json:"roleID" gorm:"index"`
	Role   *Role
}

//@author: [guoxf](https://github.com/guoxf)
//@function: CreateUser
//@description: 新增User
//@param: user User
//@return: err error
func CreateUser(user *User, isCreateFromWechat bool) error {
	//密码加密
	if user.Password != "" {
		//校验密码强度
		if !ValidPasswdStrength(user.Password) {
			return errors.New("密码强度不够")
		}
		var err error
		user.Password, err = EncryptedPassword(user.Password)
		if err != nil {
			return err
		}
	}
	user.CanDel = true
	user.UserName = strings.ToLower(user.UserName)
	var duplication bool
	var err error
	if isCreateFromWechat {
		duplication, err = dbClient.CreateWithCheckDuplication(user, "wechat_union_id = ? and wechat_open_id = ?", user.WechatUnionID, user.WechatOpenID)
	} else {
		duplication, err = dbClient.CreateWithCheckDuplication(user, "user_name = ? or mobile = ?", user.UserName, user.Mobile)
	}

	if err != nil {
		return err
	}
	if duplication {
		return errors.New("存在相同用户名或者手机号")
	}
	return nil
}

//@author: [guoxf](https://github.com/guoxf)
//@function: DeleteUser
//@description: 删除User
//@param: user User
//@return: err error
func DeleteUser(id string) (err error) {
	return dbClient.DB().Transaction(func(tx *gorm.DB) error {
		oldUser, err := GetUserById(id)
		if err != nil {
			return err
		}
		if !oldUser.CanDel {
			return errors.New("此用户不允许删除")
		}
		err = dbClient.DB().Unscoped().Delete(&UserRole{}, "user_id=?", id).Error
		if err != nil {
			return err
		}
		return dbClient.DB().Delete(&User{}, "id=?", id).Error
	})
}

type QueryUserRequest struct {
	model.CommonRequest
	UserName  string `json:"userName" form:"userName" uri:"userName"`
	Nickname  string `json:"nickname" form:"nickname" uri:"nickname"`
	IDCard    string `json:"idCard" form:"idCard" uri:"idCard"`
	Mobile    string `json:"mobile" form:"mobile" uri:"mobile"`
	WechatID  string `json:"wechatID" form:"wechatID" uri:"wechatID"`
	Title     string `json:"title" form:"title" uri:"title"`
	UserNames string `json:"userNames" form:"userNames" uri:"userNames"`
}

type QueryUserResponse struct {
	model.CommonResponse
	Data []*User `json:"data"`
}

//@author: [guoxf](https://github.com/guoxf)
//@function: GetUserInfoList
//@description: 分页查询User
//@param: user User, info PageInfo, order string, desc bool
//@return: list []*User, total int64 , err error
func QueryUser(req *apipb.QueryUserRequest, resp *apipb.QueryUserResponse) {
	db := dbClient.DB().Model(&User{})

	if req.UserName != "" {
		db = db.Where("user_name LIKE ?", "%"+req.UserName+"%")
	} else if len(req.UserNames) > 0 {
		db = db.Where("user_name in ?", strings.Split(req.UserNames, ","))
	}

	if req.Nickname != "" {
		db = db.Where("nickname LIKE ?", "%"+req.Nickname+"%")
	}

	if req.IdCard != "" {
		db = db.Where("id_card = ?", req.IdCard)
	}

	if req.Mobile != "" {
		db = db.Where("mobile LIKE ?", "%"+req.Mobile+"%")
	}

	if req.WechatID != "" {
		db = db.Where("wechat_id = ?", req.WechatID)
	}

	if req.Title != "" {
		db = db.Where("title LIKE ?", "%"+req.Title+"%")
	}

	OrderStr := "`user_name`"
	if req.OrderField != "" {
		if req.Desc {
			OrderStr = req.OrderField + " desc"
		} else {
			OrderStr = req.OrderField
		}
	}
	var err error
	var users []*User
	resp.Records, resp.Pages, err = dbClient.PageQuery(db, req.PageSize, req.PageIndex, OrderStr, &users)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = UsersToPB(users)
	}
	resp.Total = resp.Records
}

//@author: [guoxf](https://github.com/guoxf)
//@function: GetAllUsers
//@description: 获取所有User
//@return: users []User , err error
func GetAllUsers() (users []*User, err error) {
	err = dbClient.DB().Find(&users).Error
	return
}

//@author: [guoxf](https://github.com/guoxf)
//@function: GetUserById
//@description: 根据id获取user
//@param: id uint
//@return: user User , err error
func GetUserById(id string) (user User, err error) {
	err = dbClient.DB().Preload("UserRoles.Role").Preload(clause.Associations).Where("id = ?", id).First(&user).Error
	user.Password = ""
	user.WechatUnionID = ""
	user.WechatOpenID = ""
	user.RoleIDs = user.GetRoleIDs()
	return
}

//@author: [guoxf](https://github.com/guoxf)
//@function: UpdateUser
//@description: 根据id更新user
//@param: user User
//@return: err error
func UpdateUser(user *User) error {
	user.UserName = strings.ToLower(user.UserName)
	return dbClient.DB().Transaction(func(tx *gorm.DB) error {
		oldUser := &User{}
		err := tx.Preload("UserRoles").Preload(clause.Associations).Where("id = ?", user.ID).First(oldUser).Error
		if err != nil {
			return err
		}

		var deleteUserRole []string
		for _, oldUserRole := range oldUser.UserRoles {
			flag := false
			for _, newUserRole := range user.UserRoles {
				if newUserRole.ID == oldUserRole.ID {
					flag = true
				}
			}
			if !flag {
				deleteUserRole = append(deleteUserRole, oldUserRole.ID)
			}
		}
		if len(deleteUserRole) > 0 {
			err = tx.Unscoped().Delete(&UserRole{}, "id in ?", deleteUserRole).Error
			if err != nil {
				return err
			}
		}

		duplication, err := dbClient.UpdateWithCheckDuplicationAndOmit(user, []string{"password", "can_del"}, "id <> ? and (user_name = ? or mobile = ?)", user.ID, user.UserName, user.Mobile)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同用户名或者手机号")
		}

		return nil
	})
}

func EnableUser(id string, enable bool) error {
	return dbClient.DB().Model(&User{}).Where("id=?", id).Update("enable", enable).Error
}

func ResetPwd(id string, pwd string) error {
	password, err := EncryptedPassword(pwd)
	if err != nil {
		return err
	}
	return dbClient.DB().Model(&User{}).Where("id=?", id).Update("password", password).Error
}

func UpdatePwd(id string, oldPwd, newPwd string) error {
	var user User
	err := dbClient.DB().Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}
	err = scrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPwd))
	if err != nil {
		return err
	}
	return ResetPwd(id, newPwd)
}

//UpdateProfile 更新个人信息
func UpdateProfile(m *User) error {
	return dbClient.DB().Model(m).Select("gender", "country",
		"province", "city", "county", "birthday", "nickname", "description", "eid", "avatar").Where("id=?", m.ID).Updates(m).Error
}

func Login(req *apipb.LoginRequest, resp *apipb.LoginResponse) {
	req.UserName = strings.ToLower(req.UserName)
	user := &User{}
	err := dbClient.DB().Model(user).Preload("UserRoles").First(user, "user_name=?", req.UserName).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
		return
	}
	if err == gorm.ErrRecordNotFound {
		resp.Code = model.UserIsNotExist
		resp.Message = "用户不存在!"
		return
	}

	if !user.Enable {
		resp.Code = model.UserDisabled
		resp.Message = "用户已禁用!"
		return
	}

	password := []byte(user.Password)
	err = scrypt.CompareHashAndPassword(password, []byte(req.Password))
	if err != nil && err == scrypt.ErrMismatchedHashAndPassword {
		resp.Code = model.UserNameOrPasswordIsWrong
		return
	} else if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	currentUser := &apipb.CurrentUser{
		Id:       user.ID,
		UserName: user.UserName,
		Gender:   user.Gender,
		RoleIDs:  user.GetRoleIDs(),
		TenantID: user.TenantID,
	}
	t, err := token.EncodeToken(currentUser)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
		return
	}
	resp.Data = t
}

func LoginByWechat(register bool, req *User, resp *apipb.LoginResponse) {
	user := &User{}
	err := dbClient.DB().Model(user).Preload("UserRoles").First(user, "wechat_union_id=? and wechat_open_id=?", req.WechatUnionID, req.WechatOpenID).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
		return
	}
	if err == gorm.ErrRecordNotFound {
		if register {
			//TODO 随机生成默认密码
			req.Password = DefaultPwd
			err = CreateUser(req, true)
			if err != nil && err != gorm.ErrRecordNotFound {
				resp.Code = model.InternalServerError
				resp.Message = err.Error()
				return
			}
			user = req
		} else {
			resp.Code = model.UserIsNotExist
			return
		}
	}

	if !user.Enable {
		resp.Code = model.UserDisabled
		resp.Message = "用户已禁用!"
		return
	}

	currentUser := &apipb.CurrentUser{
		Id:       user.ID,
		UserName: user.UserName,
		Gender:   user.Gender,
		RoleIDs:  user.GetRoleIDs(),
		TenantID: user.TenantID,
	}
	t, err := token.EncodeToken(currentUser)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
		return
	}
	resp.Data = t
}

type CheckRegisterWithWechatResp struct {
	model.CommonResponse
	Data bool `json:"data"`
}

func CheckRegisterWithWechat(unionID, openID string, resp *CheckRegisterWithWechatResp) {
	user := &User{}
	err := dbClient.DB().Model(user).First(user, "wechat_union_id=? and wechat_open_id=?", unionID, openID).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
		return
	}
	if err == gorm.ErrRecordNotFound {
		resp.Data = false
	} else {
		resp.Data = true
	}
}

func Logout(t string) error {
	currentUser, err := token.DecodeToken(t)
	if err != nil {
		var expErr *jwt.TokenExpiredError
		if errors.As(err, &expErr) {
			log.Error(context.Background(), err)
			return nil
		}
	}
	if currentUser == nil {
		return nil
	}
	err = token.DefaultTokenCache.Del(fmt.Sprint(currentUser.Id), t)
	if err != nil {
		log.Error(context.Background(), err)
	}
	return nil
}

type UserProfile struct {
	*User
	Menus []*Menu `json:"menus"`
}

func GetUserProfile(id string) (*UserProfile, error) {
	var user = &User{}
	err := dbClient.DB().Preload("UserRoles.Role.RoleMenus").Preload(clause.Associations).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	user.Password = ""
	user.WechatUnionID = ""
	user.WechatOpenID = ""

	parents := make(map[string]*Menu)
	//去重
	haveMenu := make(map[string]*RoleMenu)
	var allMenus []*Menu
	treeMap := make(map[string]*Menu)
	err = dbClient.DB().Order("sort").Preload("MenuFuncs").Find(&allMenus).Error
	if err != nil {
		return nil, err
	}
	for _, v := range allMenus {
		treeMap[v.ID] = v
	}
	for _, userRole := range user.UserRoles {
		for _, roleMenu := range userRole.Role.RoleMenus {
			oldMenu, ok := haveMenu[roleMenu.MenuID]
			if ok {
				if oldMenu.Funcs == "" {
					oldMenu.Funcs = roleMenu.Funcs
				} else if roleMenu.Funcs != "" {
					oldMenu.Funcs += "," + roleMenu.Funcs
				}
				continue
			}
			haveMenu[roleMenu.MenuID] = roleMenu
		}
	}
	for _, roleMenu := range haveMenu {
		menu := treeMap[roleMenu.MenuID]
		//顶级menu不在菜单中显示
		if menu.Level == 0 {
			continue
		}

		funcs := strings.Split(roleMenu.Funcs, ",")
		if len(funcs) == 0 {
			menu.MenuFuncs = []*MenuFunc{}
		} else {
			var menuFuncs []*MenuFunc
			for _, fn := range menu.MenuFuncs {
				for _, f := range funcs {
					if fn.Name == f {
						menuFuncs = append(menuFuncs, fn)
					}
				}
			}
			menu.MenuFuncs = menuFuncs
		}

		if menu.Level == 1 {
			parents[menu.ID] = menu
		} else {
			parent := treeMap[menu.ParentID]
			parent.Children = append(parent.Children, menu)
		}
	}
	var result []*Menu
	for _, menu := range parents {
		if len(menu.Children) > 0 {
			sort.Slice(menu.Children, func(i, j int) bool {
				return menu.Children[i].Sort < menu.Children[j].Sort
			})
			for _, child := range menu.Children {
				if len(child.Children) > 0 {
					sort.Slice(child.Children, func(i, j int) bool {
						return child.Children[i].Sort < child.Children[j].Sort
					})
				}
			}
		}
		result = append(result, menu)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Sort < result[j].Sort
	})

	return &UserProfile{
		User:  user,
		Menus: result,
	}, nil
}
