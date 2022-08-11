package models

import (
	"database/sql"
	"time"
)

type Transaction struct {
	ID              string         `json:"id" query:"id" gorm:"primaryKey;type:varchar(255);NOT NULL"`
	Qty             uint64         `json:"qty" query:"qty" gorm:"type:int"`
	OutletProductId string         `json:"outlet_product_id" query:"outlet_product_id" gorm:"type:varchar(255);NOT NULL;index"`
	StaffId         string         `json:"staff_id" query:"staff_id" gorm:"type:varchar(255);NOT NULL;index"`
	CreatedAt       time.Time      `json:"created_at" query:"created_at" gorm:"NOT NULL"`
	CreatedBy       string         `json:"created_by" query:"created_by" gorm:"type:varchar(255);NOT NULL"`
	UpdatedAt       sql.NullTime   `json:"updated_at" query:"updated_at" gorm:"NOT NULL"`
	UpdatedBy       sql.NullString `json:"updated_by" query:"updated_by" gorm:"type:varchar(255)"`
	DeletedAt       sql.NullTime   `json:"deleted_at" query:"deleted_at"`
	DeletedBy       sql.NullString `json:"deleted_by" query:"deleted_by" gorm:"type:varchar(255)"`
}
