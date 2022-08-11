package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          string         `json:"id" xml:"id" form:"id" query:"id" gorm:"primaryKey;type:varchar(255);NOT NULL"`
	Name        string         `json:"name" xml:"name" form:"name" query:"name" gorm:"type:varchar(255);NOT NULL"`
	Description string         `json:"description" xml:"description" form:"description" query:"description" gorm:"type:text"`
	CreatedAt   time.Time      `json:"-" xml:"created_at" query:"created_at" gorm:"NOT NULL"`
	UpdatedAt   time.Time      `json:"-" xml:"updated_at" query:"updated_at" gorm:"NOT NULL"`
	DeletedAt   gorm.DeletedAt `json:"-" xml:"deleted_at" query:"deleted_at"`
	CreatedBy   string         `json:"-" xml:"created_by" query:"created_by" gorm:"type:varchar(255)"`
	UpdatedBy   string         `json:"-" xml:"updated_by" query:"updated_by" gorm:"type:varchar(255)"`
	DeletedBy   string         `json:"-" xml:"deleted_by" query:"deleted_by" gorm:"type:varchar(255)"`
}
