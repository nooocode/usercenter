package model

import (
	"database/sql"

	commonmodel "github.com/nooocode/pkg/model"
	apipb "github.com/nooocode/usercenter/api"
)

func PBToUser(in *apipb.UserInfo) *User {
	if in == nil {
		return nil
	}
	return &User{
		TenantModel: commonmodel.TenantModel{
			Model: commonmodel.Model{
				ID: in.Id,
			},
			TenantID: in.TenantID,
		},
		UserName:    in.UserName,
		Nickname:    in.Nickname,
		UserRoles:   PBToUserRoles(in.UserRoles),
		RoleIDs:     in.RoleIDs,
		Enable:      in.Enable,
		Email:       in.Email,
		Mobile:      in.Mobile,
		IDCard:      in.IdCard,
		Avatar:      in.Avatar,
		EID:         in.Eid,
		Title:       in.Title,
		Description: in.Description,
		RealName:    in.RealName,
		Gender:      in.Gender,
		Type:        in.Type,
		Group:       in.Group,
	}
}

func UserToPB(in *User) *apipb.UserInfo {
	if in == nil {
		return nil
	}
	user := &apipb.UserInfo{
		Id:          in.ID,
		TenantID:    in.TenantID,
		UserName:    in.UserName,
		Nickname:    in.Nickname,
		UserRoles:   UserRolesToPB(in.UserRoles),
		RoleIDs:     in.RoleIDs,
		Enable:      in.Enable,
		Email:       in.Email,
		Mobile:      in.Mobile,
		IdCard:      in.IDCard,
		Avatar:      in.Avatar,
		Eid:         in.EID,
		Title:       in.Title,
		Description: in.Description,
		RealName:    in.RealName,
		Gender:      in.Gender,
		Type:        in.Type,
		Group:       in.Group,
	}
	if in.Tenant != nil {
		user.TenantName = in.Tenant.Name
	}
	return user
}

func UsersToPB(in []*User) []*apipb.UserInfo {
	var list []*apipb.UserInfo
	for _, user := range in {
		list = append(list, UserToPB(user))
	}
	return list
}

func PBToUserRoles(userRoles []*apipb.UserRole) []*UserRole {
	var list []*UserRole
	for _, userRole := range userRoles {
		list = append(list, &UserRole{
			Model: commonmodel.Model{
				ID: userRole.Id,
			},
			UserID: userRole.UserID,
			RoleID: userRole.RoleID,
		})
	}
	return list
}

func UserRolesToPB(userRoles []*UserRole) []*apipb.UserRole {
	var list []*apipb.UserRole
	for _, userRole := range userRoles {
		list = append(list, &apipb.UserRole{
			Id:     userRole.ID,
			UserID: userRole.UserID,
			RoleID: userRole.RoleID,
		})
	}
	return list
}

func PBToAPI(in *apipb.APIInfo) *API {
	if in == nil {
		return nil
	}
	return &API{
		Model: commonmodel.Model{
			ID: in.Id,
		},
		Path:        in.Path,
		Group:       in.Group,
		Method:      in.Method,
		Enable:      in.Enable,
		Description: in.Description,
		CheckAuth:   in.CheckAuth,
		CheckLogin:  in.CheckLogin,
	}
}

func APIToPB(in *API) *apipb.APIInfo {
	if in == nil {
		return nil
	}
	return &apipb.APIInfo{
		Id:          in.ID,
		Path:        in.Path,
		Group:       in.Group,
		Method:      in.Method,
		Enable:      in.Enable,
		Description: in.Description,
		CheckAuth:   in.CheckAuth,
		CheckLogin:  in.CheckLogin,
	}
}

func APIsToPB(in []API) []*apipb.APIInfo {
	var list []*apipb.APIInfo
	for _, api := range in {
		list = append(list, APIToPB(&api))
	}
	return list
}

func PBToMenu(in *apipb.MenuInfo) *Menu {
	if in == nil {
		return nil
	}
	return &Menu{
		Model: commonmodel.Model{
			ID: in.Id,
		},
		Level:       in.Level,
		ParentID:    in.ParentID,
		Path:        in.Path,
		Name:        in.Name,
		Hidden:      in.Hidden,
		Component:   in.Component,
		Sort:        in.Sort,
		Cache:       in.Cache,
		DefaultMenu: in.DefaultMenu,
		Title:       in.Title,
		Icon:        in.Icon,
		Parameters:  PBToMenuParameters(in.Parameters),
		MenuFuncs:   PBToMenuFuncs(in.MenuFuncs),
	}
}

func MenuToPB(in *Menu) *apipb.MenuInfo {
	if in == nil {
		return nil
	}
	var children []*apipb.MenuInfo
	//递归退出条件
	if len(in.Children) > 0 {
		children = MenusToPB(in.Children)
	}
	return &apipb.MenuInfo{
		Id:          in.ID,
		Level:       in.Level,
		ParentID:    in.ParentID,
		Path:        in.Path,
		Name:        in.Name,
		Hidden:      in.Hidden,
		Component:   in.Component,
		Sort:        in.Sort,
		Cache:       in.Cache,
		DefaultMenu: in.DefaultMenu,
		Title:       in.Title,
		Icon:        in.Icon,
		Parameters:  MenuParametersToPB(in.Parameters),
		MenuFuncs:   MenuFuncsToPB(in.MenuFuncs),
		Children:    children,
	}
}

func MenusToPB(in []*Menu) []*apipb.MenuInfo {
	var list []*apipb.MenuInfo
	for _, menu := range in {
		list = append(list, MenuToPB(menu))
	}
	return list
}

func PBToMenuParameters(params []*apipb.MenuParameter) []*MenuParameter {
	var list []*MenuParameter
	for _, param := range params {
		list = append(list, &MenuParameter{
			Model: commonmodel.Model{
				ID: param.Id,
			},
			MenuID: param.MenuID,
			Type:   param.Type,
			Key:    param.Key,
			Value:  param.Value,
		})
	}
	return list
}

func MenuParametersToPB(params []*MenuParameter) []*apipb.MenuParameter {
	var list []*apipb.MenuParameter
	for _, param := range params {
		list = append(list, &apipb.MenuParameter{
			Id:     param.ID,
			MenuID: param.MenuID,
			Type:   param.Type,
			Key:    param.Key,
			Value:  param.Value,
		})
	}
	return list
}

func PBToMenuFuncs(params []*apipb.MenuFunc) []*MenuFunc {
	var list []*MenuFunc
	for _, param := range params {
		list = append(list, &MenuFunc{
			Model: commonmodel.Model{
				ID: param.Id,
			},
			MenuID:       param.MenuID,
			Name:         param.Name,
			Title:        param.Title,
			Hidden:       param.Hidden,
			MenuFuncApis: PBToMenuFuncApis(param.MenuFuncApis),
		})
	}
	return list
}

func MenuFuncsToPB(params []*MenuFunc) []*apipb.MenuFunc {
	var list []*apipb.MenuFunc
	for _, param := range params {
		list = append(list, &apipb.MenuFunc{
			Id:           param.ID,
			MenuID:       param.MenuID,
			Name:         param.Name,
			Title:        param.Title,
			Hidden:       param.Hidden,
			MenuFuncApis: MenuFuncApisToPB(param.MenuFuncApis),
		})
	}
	return list
}

func PBToMenuFuncApis(params []*apipb.MenuFuncApi) []MenuFuncApi {
	var list []MenuFuncApi
	for _, param := range params {
		list = append(list, MenuFuncApi{
			Model: commonmodel.Model{
				ID: param.Id,
			},
			MenuFuncID: param.MenuFuncID,
			APIID:      param.ApiID,
			API:        *PBToAPI(param.ApiInfo),
		})
	}
	return list
}

func MenuFuncApisToPB(params []MenuFuncApi) []*apipb.MenuFuncApi {
	var list []*apipb.MenuFuncApi
	for _, param := range params {
		list = append(list, &apipb.MenuFuncApi{
			Id:         param.ID,
			MenuFuncID: param.MenuFuncID,
			ApiID:      param.APIID,
			ApiInfo:    APIToPB(&param.API),
		})
	}
	return list
}

func PBToRole(in *apipb.RoleInfo) *Role {
	if in == nil {
		return nil
	}
	var tenantID *sql.NullString
	if in.TenantID != "" {
		tenantID = &sql.NullString{
			String: in.TenantID,
			Valid:  in.TenantID != "",
		}
	}
	return &Role{
		Model: commonmodel.Model{
			ID: in.Id,
		},
		TenantID:      tenantID,
		Name:          in.Name,
		ParentID:      in.ParentID,
		DefaultRouter: in.DefaultRouter,
		Description:   in.Description,
		CanDel:        in.CanDel,
		RoleMenus:     PBToRoleMenus(in.RoleMenus),
	}
}

func RoleToPB(in *Role) *apipb.RoleInfo {
	if in == nil {
		return nil
	}
	var children []*apipb.RoleInfo
	//递归退出条件
	if len(in.Children) > 0 {
		children = RolesToPB(in.Children)
	}
	tenantID := ""
	if in.TenantID != nil && in.TenantID.Valid {
		tenantID = in.TenantID.String
	}
	role := &apipb.RoleInfo{
		Id:            in.ID,
		TenantID:      tenantID,
		Name:          in.Name,
		ParentID:      in.ParentID,
		DefaultRouter: in.DefaultRouter,
		Description:   in.Description,
		CanDel:        in.CanDel,
		RoleMenus:     RoleMenusToPB(in.RoleMenus),
		Children:      children,
	}
	if in.Tenant != nil {
		role.TenantName = in.Tenant.Name
	}
	return role
}

func RolesToPB(in []*Role) []*apipb.RoleInfo {
	var list []*apipb.RoleInfo
	for _, role := range in {
		list = append(list, RoleToPB(role))
	}
	return list
}

func PBToRoleMenus(roleMenus []*apipb.RoleMenu) []*RoleMenu {
	var list []*RoleMenu
	for _, roleMenu := range roleMenus {
		list = append(list, &RoleMenu{
			Model: commonmodel.Model{
				ID: roleMenu.Id,
			},
			RoleID: roleMenu.RoleID,
			MenuID: roleMenu.MenuID,
			Funcs:  roleMenu.Funcs,
			Menu:   PBToMenu(roleMenu.Menu),
		})
	}
	return list
}

func RoleMenusToPB(roleMenus []*RoleMenu) []*apipb.RoleMenu {
	var list []*apipb.RoleMenu
	for _, roleMenu := range roleMenus {
		list = append(list, &apipb.RoleMenu{
			Id:     roleMenu.ID,
			RoleID: roleMenu.RoleID,
			MenuID: roleMenu.MenuID,
			Funcs:  roleMenu.Funcs,
			Menu:   MenuToPB(roleMenu.Menu),
		})
	}
	return list
}

func PBToTenant(in *apipb.TenantInfo) *Tenant {
	if in == nil {
		return nil
	}
	return &Tenant{
		Model: commonmodel.Model{
			ID: in.Id,
		},
		Name:          in.Name,
		Contact:       in.Contact,
		CellPhone:     in.CellPhone,
		Address:       in.Address,
		BusinessScope: in.BusinessScope,
		AreaCovered:   in.AreaCovered,
		StaffSize:     in.StaffSize,
		Enable:        in.Enable,
	}
}

func TenantToPB(in *Tenant) *apipb.TenantInfo {
	if in == nil {
		return nil
	}
	return &apipb.TenantInfo{
		Id:            in.ID,
		Name:          in.Name,
		Contact:       in.Contact,
		CellPhone:     in.CellPhone,
		Address:       in.Address,
		BusinessScope: in.BusinessScope,
		AreaCovered:   in.AreaCovered,
		StaffSize:     in.StaffSize,
		Enable:        in.Enable,
	}
}

func TenantsToPB(in []*Tenant) []*apipb.TenantInfo {
	var list []*apipb.TenantInfo
	for _, api := range in {
		list = append(list, TenantToPB(api))
	}
	return list
}

func UserProfileToUser(in *apipb.UserProfile) *User {
	if in == nil {
		return nil
	}
	return &User{
		TenantModel: commonmodel.TenantModel{
			Model: commonmodel.Model{
				ID: in.Id,
			},
		},
		Nickname:    in.Nickname,
		Email:       in.Email,
		Mobile:      in.Mobile,
		IDCard:      in.IdCard,
		Avatar:      in.Avatar,
		RealName:    in.RealName,
		Gender:      in.Gender,
		Country:     in.Country,
		Province:    in.Province,
		City:        in.City,
		County:      in.County,
		EID:         in.Eid,
		Birthday:    in.Birthday,
		Description: in.Description,
		UserName:    in.UserName,
	}
}
