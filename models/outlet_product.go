package models

import (
	"time"
)

type OutletProduct struct {
	ID        string    `json:"id" xml:"id" form:"id" query:"id" gorm:"primaryKey;type:varchar(255);NOT NULL"`
	OutletId  string    `json:"outlet_id" xml:"outlet_id" form:"outlet_id" query:"outlet_id" gorm:"type:varchar(255);NOT NULL;index"`
	ProductId string    `json:"product_id" xml:"product_id" form:"product_id" query:"product_id" gorm:"type:varchar(255);NOT NULL;index"`
	Price     float64   `json:"price" xml:"price" form:"price" query:"price" gorm:"type:double;NOT NULL;default:0"`
	Stock     int       `json:"stock" xml:"stock" form:"stock" query:"stock" gorm:"type:int(11);NOT NULL;default:0"`
	CreatedAt time.Time `json:"created_at" xml:"created_at" query:"created_at" gorm:"NOT NULL;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" xml:"updated_at" query:"updated_at" gorm:"NOT NULL;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt time.Time `json:"deleted_at" xml:"deleted_at" query:"deleted_at"`
	CreatedBy string    `json:"created_by" xml:"created_by" query:"created_by" gorm:"type:varchar(255)"`
	UpdatedBy string    `json:"updated_by" xml:"updated_by" query:"updated_by" gorm:"type:varchar(255)"`
	DeletedBy string    `json:"deleted_by" xml:"deleted_by" query:"deleted_by" gorm:"type:varchar(255)"`
}
