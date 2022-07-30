package models

import (
	"time"
)

type ModelProduct struct {
	ID         string    `json:"id" xml:"id" form:"id" query:"id" gorm:"type:varchar; primary_key"`
	Name       string    `json:"name" xml:"name" form:"name" query:"name" gorm:"type:varchar; not null"`
	SKU        uint64    `json:"sku" xml:"sku" form:"sku" query:"sku" gorm:"type:bigint; index"`
	CategoryId string    `json:"category_id" xml:"category_id" form:"category_id" query:"category_id" gorm:"type:varchar; index"`
	CreatedAt  time.Time `json:"created_at" xml:"created_at" query:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" xml:"updated_at" query:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at" xml:"deleted_at" query:"deleted_at"`
	CreatedBy  string    `json:"created_by" xml:"created_by" query:"created_by"`
	UpdatedBy  string    `json:"updated_by" xml:"updated_by" query:"updated_by"`
	DeletedBy  string    `json:"deleted_by" xml:"deleted_by" query:"deleted_by"`
}
