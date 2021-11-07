package model

import (
	"context"
	"errors"

	"codeup.aliyun.com/atali/pkg/model"
	"codeup.aliyun.com/atali/pkg/utils/log"
)

type API struct {
	model.Model
	Path        string `json:"path" gorm:"size:200;index;comment:路径"`
	Group       string `json:"group" gorm:"size:50;index;comment:分组"`
	Method      string `json:"method" gorm:"size:50;index;default:POST;comment:方法"`
	Description string `json:"description" gorm:"size:200;index;comment:中文描述"`
	Enable      bool   `json:"enable" gorm:"index;comment:是否启用API"`
	//如果不开启权限校验，那么在每个角色都加上casbin rule
	//1. 不需要登录就可以直接访问
	//2. 需要登录但不需要校验权限，也就是所有人都可以访问
	//3. 需要登录并且校验权限
	CheckAuth  bool `json:"checkAuth" gorm:"index;comment:是否校验权限"`
	CheckLogin bool `json:"checkLogin" gorm:"index;comment:是否校验登录"`
}

func (API) TableName() string {
	return "api"
}

//@author: [guoxf](https://github.com/guoxf)
//@function: CreateAPI
//@description: 新增API
//@param: api API
//@return: err error
func CreateAPI(api *API) error {
	duplication, err := dbClient.CreateWithCheckDuplication(api, "path = ? AND method = ?", api.Path, api.Method)
	if err != nil {
		return err
	}
	if duplication {
		return errors.New("存在相同api")
	}
	updateNotCheckAuthRule()
	updateNotCheckLoginRule()
	return nil
}

//@author: [guoxf](https://github.com/guoxf)
//@function: DeleteApi
//@description: 删除API
//@param: api API
//@return: err error
func DeleteApi(id string) (err error) {
	api := &API{}
	err = dbClient.DB().Where("id = ?", id).First(&api).Error
	if err != nil {
		return err
	}
	err = dbClient.DB().Delete(api).Error
	if err != nil {
		return err
	}
	ClearCasbin(1, api.Path, api.Method)
	return nil
}

type QueryAPIRequest struct {
	model.CommonRequest
	Path       string `json:"path" form:"path" path:"path"`
	Method     string `json:"method" form:"method" path:"method"`
	Group      string `json:"group" form:"group" path:"group"`
	CheckAuth  int    `json:"checkAuth" form:"checkAuth" path:"checkAuth"`
	CheckLogin int    `json:"checkLogin" form:"checkLogin" path:"checkLogin"`
}

type QueryAPIResponse struct {
	model.CommonResponse
	Data []API `json:"data"`
}

//@author: [guoxf](https://github.com/guoxf)
//@function: GetAPIInfoList
//@description: 分页查询API
//@param: api API, info PageInfo, order string, desc bool
//@return: list []*API, total int64 , err error
func QueryAPI(req *QueryAPIRequest, resp *QueryAPIResponse) {
	db := dbClient.DB().Model(&API{})

	if req.Path != "" {
		db = db.Where("path LIKE ?", "%"+req.Path+"%")
	}

	if req.Method != "" {
		db = db.Where("method = ?", req.Method)
	}

	if req.Group != "" {
		db = db.Where("`group` = ?", req.Group)
	}

	if req.CheckAuth > 0 {
		db = db.Where("check_auth = ?", req.CheckAuth == 1)
	}

	if req.CheckLogin > 0 {
		db = db.Where("check_login = ?", req.CheckLogin == 1)
	}

	OrderStr := "`path`"
	if req.OrderField != "" {
		if req.Desc {
			OrderStr = req.OrderField + " desc"
		} else {
			OrderStr = req.OrderField
		}
	}
	var err error
	resp.Records, resp.Pages, err = dbClient.PageQuery(db, req.PageSize, req.PageIndex, OrderStr, &resp.Data)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	resp.Total = resp.Records
}

//@author: [guoxf](https://github.com/guoxf)
//@function: GetAllAPIs
//@description: 获取所有API
//@return: apis []API , err error
func GetAllAPIs() (apis []API, err error) {
	err = dbClient.DB().Find(&apis).Error
	return
}

//@author: [guoxf](https://github.com/guoxf)
//@function: GetAPIById
//@description: 根据id获取api
//@param: id string
//@return: api API , err error
func GetAPIById(id string) (api API, err error) {
	err = dbClient.DB().Where("id = ?", id).First(&api).Error
	return
}

//@author: [guoxf](https://github.com/guoxf)
//@function: UpdateAPI
//@description: 根据id更新api
//@param: api API
//@return: err error
func UpdateAPI(api *API) error {
	var oldA API
	err := dbClient.DB().Where("id = ?", api.ID).First(&oldA).Error
	if err != nil {
		return err
	}

	if oldA.Path != api.Path || oldA.Method != api.Method {
		err = UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method)
		if err != nil {
			return err
		}
	}

	duplication, err := dbClient.UpdateWithCheckDuplication(api, "id <> ? and path = ? AND method = ?", api.ID, api.Path, api.Method)
	if err != nil {
		return err
	}
	if duplication {
		return errors.New("存在相同api")
	}
	updateNotCheckAuthRule()
	updateNotCheckLoginRule()
	return nil
}

func EnableAPI(id string, enable bool) error {
	err := dbClient.DB().Table("api").Where("id=?", id).Update("enable", enable).Error
	if err != nil {
		return err
	}
	updateNotCheckAuthRule()
	updateNotCheckLoginRule()
	return nil
}

func updateNotCheckAuthRule() {
	resp := &QueryAPIResponse{}
	QueryAPI(&QueryAPIRequest{
		CheckAuth:  2,
		CheckLogin: 1,
		CommonRequest: model.CommonRequest{
			PageInfo: model.PageInfo{
				PageSize: 1000,
			},
		},
	}, resp)
	if resp.Code != model.Success {
		log.Errorf(context.Background(), "updateNotCheckAuthRule error:%s", resp.Message)
		return
	}
	var newRules []*CasbinRule
	for _, api := range resp.Data {
		newRules = append(newRules, &CasbinRule{
			Ptype:     "p",
			RoleID:    "0",
			Path:      api.Path,
			Method:    api.Method,
			CheckAuth: "false",
		})
	}
	_, err := ClearCasbin(0, "0")
	if err != nil {
		log.Errorf(context.Background(), "updateNotCheckAuthRule error:%v", err)
	}
	if len(newRules) > 0 {
		err = UpdateCasbin("0", newRules)
		if err != nil {
			log.Errorf(context.Background(), "updateNotCheckAuthRule error:%v", err)
		}
	}
}

func updateNotCheckLoginRule() {
	resp := &QueryAPIResponse{}
	QueryAPI(&QueryAPIRequest{
		CheckLogin: 2,
		CommonRequest: model.CommonRequest{
			PageInfo: model.PageInfo{
				PageSize: 1000,
			},
		},
	}, resp)
	if resp.Code != model.Success {
		log.Errorf(context.Background(), "updateNotCheckLoginRule error:%s", resp.Message)
		return
	}
	var newRules []*CasbinRule
	for _, api := range resp.Data {
		newRules = append(newRules, &CasbinRule{
			Ptype:     "p",
			RoleID:    "-1",
			Path:      api.Path,
			Method:    api.Method,
			CheckAuth: "false",
		})
	}
	_, err := ClearCasbin(0, "-1")
	if err != nil {
		log.Errorf(context.Background(), "updateNotCheckLoginRule error:%v", err)
	}
	if len(newRules) > 0 {
		err = UpdateCasbin("-1", newRules)
		if err != nil {
			log.Errorf(context.Background(), "updateNotCheckLoginRule error:%v", err)
		}
	}
}
