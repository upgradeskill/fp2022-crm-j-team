package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string    `json:"id" xml:"id" form:"id" query:"id" gorm:"primaryKey;type:varchar(36);NOT NULL"`
	Name      string    `json:"name" xml:"name" form:"name" query:"name" gorm:"type:varchar(150);NOT NULL"`
	Email     string    `json:"email" xml:"email" form:"email" query:"email" gorm:"type:varchar(50)"`
	Password  string    `json:"password" xml:"password" form:"password" query:"password" gorm:"type:varchar(255)"`
	Role      string    `json:"role" xml:"role" form:"role" query:"role" gorm:"type:varchar(15)"`
	OutletId  string    `json:"outlet_id" xml:"outlet_id" form:"outlet_id" query:"outlet_id" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at" xml:"created_at" query:"created_at" gorm:"NOT NULL;"`
	UpdatedAt time.Time `json:"updated_at" xml:"updated_at" query:"updated_at" gorm:"NOT NULL;"`
	DeletedAt time.Time `json:"deleted_at" xml:"deleted_at" query:"deleted_at"`
	CreatedBy string    `json:"created_by" xml:"created_by" query:"created_by" gorm:"type:varchar(255)"`
	UpdatedBy string    `json:"updated_by" xml:"updated_by" query:"updated_by" gorm:"type:varchar(255)"`
	DeletedBy string    `json:"deleted_by" xml:"deleted_by" query:"deleted_by" gorm:"type:varchar(255)"`
}

func (m *User) BeforeCreate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.CreatedAt = time.Now()
	return nil
}

func (m *User) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}
