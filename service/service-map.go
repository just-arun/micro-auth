package service

import (
	"fmt"

	"github.com/just-arun/micro-auth/model"
	"github.com/just-arun/micro-auth/pubsub"
	"github.com/just-arun/micro-auth/util"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

type serviceMap struct{}

func ServiceMap() serviceMap {
	return serviceMap{}
}

func (st serviceMap) Add(db *gorm.DB, data *model.ServiceMap) error {
	return db.Save(&data).Error
}

func (st serviceMap) GetOne(db *gorm.DB, id uint) (data *model.ServiceMap, err error) {
	err = db.Model(&model.ServiceMap{ID: id}).Scan(&data).Error
	return
}

func (st serviceMap) GetMany(db *gorm.DB, searchQuery string, pagination *util.Pagination) (data []model.ServiceMap, err error) {
	tnx := db.Model(&model.ServiceMap{})

	if len(searchQuery) > 0 {
		tnx = tnx.Where("key LIKE ?", "%"+searchQuery+"%")
		if pagination != nil {
			tnx.Count(&pagination.Total)
		}
	} else {
		if pagination != nil {
			tnx.Count(&pagination.Total)
		}
	}

	if pagination != nil {
		tnx = tnx.Offset(int(pagination.Skip))
		tnx = tnx.Limit(int(pagination.Limit))
	}

	err = tnx.Scan(&data).Error
	return
}

func (st serviceMap) UpdateOne(db *gorm.DB, id uint, data *model.ServiceMap) (err error) {
	data.ID = id
	fmt.Println(data)
	return db.Save(&data).Error
}

func (st serviceMap) UpdateMany(db *gorm.DB, data []model.ServiceMap) (err error) {
	return db.Save(&data).Error
}

func (st serviceMap) DeleteOne(db *gorm.DB, id uint) (err error) {
	return db.Delete(&model.ServiceMap{ID: id}).Error
}

func (st serviceMap) PublishSitemap(db *gorm.DB, con *nats.EncodedConn) error {
	data, err := st.GetMany(db, "", nil)
	if err != nil {
		return err
	}
	return pubsub.Publisher().ChangeServiceMap(con, data)
}
