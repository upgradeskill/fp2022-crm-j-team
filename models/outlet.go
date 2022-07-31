package models

import (
	"time"
)

type Outlet struct {
	ID        string    `json:"id" xml:"id" form:"id" query:"id" gorm:"primaryKey;type:varchar(255);NOT NULL"`
	Name      string    `json:"name" xml:"name" form:"name" query:"name" gorm:"type:varchar(255);NOT NULL"`
	Phone     string    `json:"phone" xml:"phone" form:"phone" query:"phone" gorm:"type:bigint"`
	Address   string    `json:"address" xml:"address" form:"address" query:"address" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at" xml:"created_at" query:"created_at" gorm:"NOT NULL;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" xml:"updated_at" query:"updated_at" gorm:"NOT NULL;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt time.Time `json:"deleted_at" xml:"deleted_at" query:"deleted_at"`
	CreatedBy string    `json:"created_by" xml:"created_by" query:"created_by" gorm:"type:varchar(255)"`
	UpdatedBy string    `json:"updated_by" xml:"updated_by" query:"updated_by" gorm:"type:varchar(255)"`
	DeletedBy string    `json:"deleted_by" xml:"deleted_by" query:"deleted_by" gorm:"type:varchar(255)"`
}
