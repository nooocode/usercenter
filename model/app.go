package model

import (
	"errors"

	"github.com/nooocode/pkg/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type APP struct {
	model.Model
	Name        string    `json:"name" gorm:"size:200;index"`
	Entry       string    `json:"entry" gorm:"size:200;"`
	DevEntry    string    `json:"devEntry" gorm:"size:200;"`
	TestEntry   string    `json:"testEntry" gorm:"size:200;"`
	PreEntry    string    `json:"preEntry" gorm:"size:200;"`
	DisplayName string    `json:"displayName" gorm:"size:100;comment:显示名字"`
	Credentials bool      `json:"credentials"`
	Description string    `json:"description" gorm:"size:200;"`
	Props       []APPProp `json:"props"`
	Container   string    `json:"container" gorm:"size:100"`
	ActiveRule  string    `json:"activeRule" gorm:"size:200"`
	Enable      bool      `json:"enable"`
}

type APPProp struct {
	gorm.Model
	APPID uint   `json:"appID"`
	Key   string `json:"key" gorm:"size:50;"`
	Value string `json:"value" gorm:"size:200;"`
}

func CreateAPP(md *APP) error {
	duplication, err := dbClient.CreateWithCheckDuplication(md, "name = ?", md.Name)
	if err != nil {
		return err
	}
	if duplication {
		return errors.New("存在相同APP")
	}
	return nil
}

func DeleteAPP(id string) (err error) {
	return dbClient.DB().Unscoped().Delete(&APP{}, "id=?", id).Error
}

type QueryAPPRequest struct {
	model.CommonRequest
	Name   string `json:"name" form:"name" uri:"name"`
	Enable int    `json:"enable" form:"enable" uri:"enable"`
}

type QueryAPPResponse struct {
	model.CommonResponse
	Data []*APP `json:"data"`
}

func QueryAPP(req *QueryAPPRequest, resp *QueryAPPResponse) {
	db := dbClient.DB().Model(&APP{})

	if req.Name != "" {
		db = db.Where("name LIKE ?", "%"+req.Name+"%")
	}

	if req.Enable != -1 {
		db = db.Where("enable = ?", req.Enable == 1)
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

func GetAllAPPs() (mds []*APP, err error) {
	err = dbClient.DB().Preload("Props").Where("enable=1").Find(&mds).Error
	return
}

func GetAPPById(id string) (md APP, err error) {
	err = dbClient.DB().Preload("Props").Where("id = ?", id).First(&md).Error
	return
}

func UpdateAPP(md *APP) error {
	return dbClient.DB().Transaction(func(tx *gorm.DB) error {
		oldAPP := &APP{}
		err := tx.Preload("Props").Preload(clause.Associations).Where("id = ?", md.ID).First(oldAPP).Error
		if err != nil {
			return err
		}
		var deleteFile []uint
		for _, oldFile := range oldAPP.Props {
			flag := false
			for _, newFile := range md.Props {
				if newFile.ID == oldFile.ID {
					flag = true
				}
			}
			if !flag {
				deleteFile = append(deleteFile, oldFile.ID)
			}
		}
		if len(deleteFile) > 0 {
			err = tx.Unscoped().Delete(&APPProp{}, "id in ?", deleteFile).Error
			if err != nil {
				return err
			}
		}

		duplication, err := dbClient.UpdateWithCheckDuplication2(tx.Session(&gorm.Session{FullSaveAssociations: true}), md, "id <> ? and name = ?", md.ID, md.Name)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同APP")
		}

		return nil
	})
}
