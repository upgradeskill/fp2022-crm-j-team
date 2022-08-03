package models

import (
	"time"
)

type Transaction struct {
	ID              string    `json:"id" xml:"id" query:"id" gorm:"primaryKey;type:varchar(255);NOT NULL"`
	Qty             uint64    `json:"qty" xml:"qty" query:"qty" gorm:"type:int"`
	OutletProductId string    `json:"outlet_product_id" xml:"outlet_product_id" query:"outlet_product_id" gorm:"type:varchar(255);NOT NULL;index"`
	StaffId         string    `json:"staff_id" xml:"staff_id" query:"staff_id" gorm:"type:varchar(255);NOT NULL;index"`
	CreatedAt       time.Time `json:"created_at" xml:"created_at" query:"created_at" gorm:"NOT NULL"`
	CreatedBy       string    `json:"created_by" xml:"created_by" query:"created_by" gorm:"type:varchar(255);NOT NULL"`
	UpdatedAt       time.Time `json:"updated_at" xml:"updated_at" query:"updated_at" gorm:"NOT NULL"`
	UpdatedBy       string    `json:"updated_by" xml:"updated_by" query:"updated_by" gorm:"type:varchar(255)"`
	DeletedAt       time.Time `json:"deleted_at" xml:"deleted_at" query:"deleted_at"`
	DeletedBy       string    `json:"deleted_by" xml:"deleted_by" query:"deleted_by" gorm:"type:varchar(255)"`
}
