package model

import (
	"errors"

	"github.com/jinzhu/copier"
	"github.com/nooocode/pkg/model"
	apipb "github.com/nooocode/usercenter/api"
	"gorm.io/gorm"
)

type Tenant struct {
	model.Model
	// 租户名称
	// required: true
	// @inject_tag: validate:"required"
	Name string `json:"name" validate:"required" gorm:"index;size:200"`
	// 联系人
	// required: true
	// @inject_tag: validate:"required"
	Contact string `json:"contact" validate:"required" gorm:"size:100"`
	// 联系人电话
	// required: true
	// @inject_tag: validate:"required"
	CellPhone string `json:"cellPhone" validate:"required" gorm:"size:50"`
	// 地址
	// required: true
	// @inject_tag: validate:"required"
	Address string `json:"address" validate:"required" gorm:"size:200"`
	// 业务范围
	BusinessScope string `json:"businessScope" gorm:"size:200"`
	// 占地面积
	AreaCovered string `json:"areaCovered" gorm:"size:100"`
	// 人员规模
	StaffSize int  `json:"staffSize"`
	Enable    bool `json:"enable" gorm:"index"`
}

func CreateTenant(m *Tenant) error {
	duplication, err := dbClient.CreateWithCheckDuplication(m, " name =? ", m.Name)
	if err != nil {
		return err
	}
	if duplication {
		return errors.New("存在相同租户")
	}
	return nil
}

func UpdateTenant(m *Tenant) error {
	return dbClient.DB().Transaction(func(tx *gorm.DB) error {
		duplication, err := dbClient.UpdateWithCheckDuplication2(tx.Session(&gorm.Session{FullSaveAssociations: true}), m, "id != ?  and  name =? ", m.ID, m.Name)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同租户")
		}

		return nil
	})
}

type QueryTenantRequest struct {
	model.CommonRequest
	Name string `json:"name" form:"name" uri:"name"`
}

type QueryTenantResponse struct {
	model.CommonResponse
	Data []*Tenant `json:"data"`
}

func QueryTenant(req *apipb.QueryTenantRequest, resp *apipb.QueryTenantResponse) {
	db := dbClient.DB().Model(&Tenant{})
	if req.Name != "" {
		db = db.Where("name LIKE ?", "%"+req.Name+"%")
	}

	OrderStr := "`created_at`"
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

func GetAllTenant() (list []*Tenant, err error) {
	err = dbClient.DB().Find(&list).Error
	return
}

func GetTenantByID(id string) (*Tenant, error) {
	m := &Tenant{}
	err := dbClient.DB().Where("id = ?", id).First(m).Error
	return m, err
}

func DeleteTenant(id string) (err error) {
	return dbClient.DB().Unscoped().Delete(&Tenant{}, "id=?", id).Error
}

func CopyTenant(id string) error {
	from, err := GetTenantByID(id)
	if err != nil {
		return err
	}
	to := &Tenant{}
	err = copier.CopyWithOption(to, from, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if err != nil {
		return err
	}
	to.Name += " Copy"
	return CreateTenant(to)
}

func EnableTenant(id string, enable bool) error {
	err := dbClient.DB().Model(&Tenant{}).Where("id=?", id).Update("enable", enable).Error
	if err != nil {
		return err
	}
	return nil
}
