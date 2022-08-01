package models

import (
	"time"
)

type Product struct {
	ID         string    `json:"id" xml:"id" query:"id" gorm:"primaryKey;type:varchar(255);NOT NULL"`
	Name       string    `json:"name" xml:"name" query:"name" gorm:"type:varchar(255);NOT NULL"`
	SKU        uint64    `json:"sku" xml:"sku" query:"sku" gorm:"type:varchar(255)"`
	CategoryId string    `json:"category_id" xml:"category_id" query:"category_id" gorm:"type:varchar(255);NOT NULL;index"`
	CreatedAt  time.Time `json:"created_at" xml:"created_at" query:"created_at" gorm:"NOT NULL"`
	UpdatedAt  time.Time `json:"updated_at" xml:"updated_at" query:"updated_at" gorm:"NOT NULL"`
	DeletedAt  time.Time `json:"deleted_at" xml:"deleted_at" query:"deleted_at"`
	CreatedBy  string    `json:"created_by" xml:"created_by" query:"created_by" gorm:"type:varchar(255)"`
	UpdatedBy  string    `json:"updated_by" xml:"updated_by" query:"updated_by" gorm:"type:varchar(255)"`
	DeletedBy  string    `json:"deleted_by" xml:"deleted_by" query:"deleted_by" gorm:"type:varchar(255)"`
}
