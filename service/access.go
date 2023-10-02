package service

import (
	"fmt"
	"os"
	"strings"

	"github.com/just-arun/micro-auth/model"
	"github.com/just-arun/micro-auth/util"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type access struct{}

func Access() access {
	return access{}
}

func (r access) AddOne(db *gorm.DB, access model.Access) error {
	return db.Save(&access).Error
}

func (r access) GetAll(db *gorm.DB) (accesses []model.Access, err error) {
	err = db.Model(&model.Access{}).Scan(&accesses).Error
	return
}

func (r access) GetMany(db *gorm.DB, filter []string) (accesses []model.Access, err error) {
	err = db.Model(&model.Access{}).Where("key IN ?", filter).Scan(&accesses).Error
	return
}

func (r access) UpdateOneName(db *gorm.DB, id uint, name string) (err error) {
	err = db.Model(&model.Access{ID: id}).
		Update("name", name).
		Error
	return
}

func (r access) GetSitemapAcl(e *echo.Echo, db *gorm.DB) {

	// data, _ := json.MarshalIndent(e.Routes(), "", "  ")
	// fmt.Println(string(data))

	var data []model.Access
	aclData := `package acl

	type ACL string

	const (
	`
	for _, v := range util.GetPath(e) {
		da := model.Access{
			Key:  v.Key,
			Name: strings.ReplaceAll(v.Value, ".", " "),
		}
		val := fmt.Sprintf(` ACL%v ACL = "%v"
	    `,
			strings.ReplaceAll(strings.ReplaceAll(v.Value, "auth.", ""), ".", ""),
			v.Key,
		)
		includes := false
		for _, v := range data {
			if v.Key == da.Key {
				includes = true
				break
			}
		}
		if includes {
			continue
		}
		aclData += val
		data = append(data, da)
	}
	aclData += `)
	`
	os.WriteFile("acl/acl.go", []byte(aclData), 0644)

	for _, v := range data {
		db.Save(&v)
	}
}
