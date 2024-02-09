package service

import (
	"encoding/json"
	"fmt"

	"github.com/just-arun/micro-auth/model"
	"github.com/just-arun/micro-auth/util"
	"gorm.io/gorm"
)

type general struct{}

func General() general {
	return general{}
}

func (r general) Create(db *gorm.DB, data *model.General) error {
	return db.Save(data).Error
}

func (r general) Get(db *gorm.DB) (data *model.General, err error) {
	err = db.
		Model(&model.General{Active: true}).
		Preload("UpdatedBy").
		Preload("UpdatedBy.Roles").
		Last(&data).
		Error
	return
}

func (r general) GetMany(db *gorm.DB, pagination *util.Pagination) (data []model.General, err error) {

	// tnx := db.Model(&model.General{}).
	// 	Order("id DESC").
	// 	Preload("UpdatedBy").
	// 	Preload("UpdatedBy.Roles")

	// if pagination != nil {
	// 	tnx.Count(&pagination.Total)
	// }

	// if pagination != nil {
	// 	if pagination != nil {
	// 		tnx = tnx.Offset(int(pagination.Skip))
	// 		tnx = tnx.Limit(int(pagination.Limit))
	// 	}
	// }

	query := util.NamedStringFormate(`
	SELECT 
			u.id as "author.id", 
			u.email as "author.email", 
			u.user_name as "author.userName",
			
			g.id as "id",
			g.name as "name",
			g.can_login as canLogin,
			g.can_register as canRegister,
			g.http_only_cookie as httpOnlyCookie,
			g.access_token_expiry_time as accessTokenExpireTime,
			g.refresh_token_expiry_time as refreshTokenExpireTime,
			g.organization_email_domain as organizationEmailDomain,
			g.token_placement as tokenPlacement,
			g.updated_description as updatedDescription,
			g.active
		FROM general_updated_user gu
		JOIN users u on u.id = gu.user_id
		JOIN generals g on g.id = gu.general_id
	`, map[string]interface{}{

	})
	
	
	query += `ORDER BY g.active DESC
	`

	query += `;`

	tnx := db.Raw(query)


	mData := []map[string]interface{}{}

	if err = tnx.Scan(&mData).Error; err != nil {
		return
	}

	// fmt.Println(mData)

	data = []model.General{}

	for _, obj := range mData {
		val := util.ParseDotedKeyToNestedMap(obj).(map[string]interface{})
		bData, err := json.Marshal(val)
		if err != nil {
			return nil, err
		}

		fmt.Println(string(bData), "\n\n")
		da := model.General{}
		if err := util.ConvertMapToStruct(val, &da); err != nil {
			return nil, err
		}
		data = append(data, da)
	}

	return
}

func (r general) Update(db *gorm.DB, data *model.General) (err error) {
	tx := db.Model(&model.General{}).Where("active = ?", true).Update("active", false)
	if err = tx.Error; err != nil {
		return
	}
	data.ID = 0
	return tx.Save(&data).Error
}
