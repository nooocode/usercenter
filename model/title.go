package model

import (
	"errors"

	"github.com/nooocode/pkg/model"
)

type Title struct {
	model.TenantModel
	Name        string `json:"name" gorm:"size:100;"`
	DisplayName string `json:"displayName" gorm:"size:100;"`
	Description string `json:"description" gorm:"size:200;"`
}

func CreateTitle(title *Title) error {
	duplication, err := dbClient.CreateWithCheckDuplication(title, "name = ?", title.Name)
	if err != nil {
		return err
	}
	if duplication {
		return errors.New("存在相同职位")
	}
	return nil
}

func DeleteTitle(id string) (err error) {
	return dbClient.DB().Unscoped().Delete(&Title{}, "id=?", id).Error
}

type QueryTitleRequest struct {
	model.CommonRequest
	Name        string `json:"name" form:"name" uri:"name"`
	DisplayName string `json:"displayName" form:"displayName" uri:"displayName"`
}

type QueryTitleResponse struct {
	model.CommonResponse
	Data []Title `json:"data"`
}

func QueryTitle(req *QueryTitleRequest, resp *QueryTitleResponse) {
	db := dbClient.DB().Model(&Title{})

	if req.Name != "" {
		db = db.Where("name LIKE ?", "%"+req.Name+"%")
	}

	if req.DisplayName != "" {
		db = db.Where("display_name LIKE ?", "%"+req.DisplayName+"%")
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
	resp.Records, resp.Pages, err = dbClient.PageQuery(db, req.PageSize, req.PageIndex, OrderStr, &resp.Data)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	resp.Total = resp.Records
}

func GetAllTitles() (titles []Title, err error) {
	err = dbClient.DB().Find(&titles).Error
	return
}

func GetTitleById(id string) (title Title, err error) {
	err = dbClient.DB().Where("id = ?", id).First(&title).Error
	return
}

func UpdateTitle(title *Title) error {
	duplication, err := dbClient.UpdateWithCheckDuplication(title, "id <> ? and name = ?", title.ID, title.Name)
	if err != nil {
		return err
	}
	if duplication {
		return errors.New("存在相同职位")
	}

	return nil
}
