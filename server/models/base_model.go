package models

import (
	"deltaboard-server/internal/library/random"
	"deltaboard-server/internal/library/reflect_model"
	"strings"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	Id        string `json:"id" gorm:"type:varchar(191);primary_key" description:"Primary key"`
	CreatedAt int64  `json:"created_at" gorm:"type:int(11)" description:"创建时间"`
	UpdatedAt int64  `json:"updated_at" gorm:"type:int(11)" description:"更新时间"`
}

func (model *BaseModel) BeforeCreate(*gorm.DB) error {
	model.CreatedAt = time.Now().Unix()
	model.UpdatedAt = model.CreatedAt

	return nil
}

func (model *BaseModel) BeforeUpdate(*gorm.DB) error {
	model.UpdatedAt = time.Now().Unix()
	return nil
}

func (model *BaseModel) Create(value interface{}, db *gorm.DB) error {
	for {
		model.Id = random.GenerateRandomString(32)
		err := db.Create(value).Error

		if err != nil && strings.Contains(err.Error(), "Duplicate") && strings.Contains(err.Error(), "PRIMARY") {
			continue
		} else {
			return err
		}
	}
}

// TryUpdateStatus 数据库排他锁模式下的状态更新
func TryUpdateStatus(model, statusToUpdate interface{}, statusFieldName string, db *gorm.DB) (updated bool, err error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else if err != nil {
			tx.Rollback()
		}
	}()
	if err = tx.Error; err != nil {
		return
	}
	id, err := reflect_model.GetValueByFieldName("Id", model)
	if err != nil {
		return
	}
	if err = tx.Set("gorm:query_option", "FOR UPDATE").First(model, "id = ?", id.Interface()).Error; err != nil {
		return
	}
	statusValue, err := reflect_model.GetValueByFieldName(statusFieldName, model)
	if err != nil {
		return
	}
	if statusToUpdate == statusValue.Interface() {
		tx.Rollback()
		return
	}
	if err = reflect_model.SetValueByFiledName(statusFieldName, model, statusToUpdate); err != nil {
		return
	}
	if err = tx.Save(model).Error; err != nil {
		return
	}
	// commit事务，释放锁
	err = tx.Commit().Error
	updated = true
	return
}
