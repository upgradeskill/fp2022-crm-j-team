package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/upgradeskill/fp2022-crm-j-team/pkg"
	"gorm.io/gorm"
)

type User struct {
	ID        string         `json:"id" xml:"id" form:"id" query:"id" gorm:"primaryKey;type:varchar(36);NOT NULL"`
	Name      string         `json:"name" xml:"name" form:"name" query:"name" gorm:"type:varchar(150);NOT NULL"`
	Email     string         `json:"email" xml:"email" form:"email" query:"email" gorm:"type:varchar(50)"`
	Password  string         `json:"-" xml:"password" form:"password" query:"password" gorm:"type:varchar(255)"`
	Role      string         `json:"role" xml:"role" form:"role" query:"role" gorm:"type:varchar(15)"`
	OutletId  string         `json:"outlet_id" xml:"outlet_id" form:"outlet_id" query:"outlet_id" gorm:"type:varchar(255)"`
	CreatedAt time.Time      `json:"-" xml:"created_at" query:"created_at" gorm:"NOT NULL;"`
	UpdatedAt time.Time      `json:"-" xml:"updated_at" query:"updated_at" gorm:"NOT NULL;"`
	DeletedAt gorm.DeletedAt `json:"-" xml:"deleted_at" query:"deleted_at"`
	CreatedBy string         `json:"-" xml:"created_by" query:"created_by" gorm:"type:varchar(255)"`
	UpdatedBy string         `json:"-" xml:"updated_by" query:"updated_by" gorm:"type:varchar(255)"`
	DeletedBy string         `json:"-" xml:"deleted_by" query:"deleted_by" gorm:"type:varchar(255)"`
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

func CreateUserSeeder(db *gorm.DB) {
	if err := db.AutoMigrate(&User{}); err == nil && db.Migrator().HasTable(&User{}) {
		if err := db.First(&User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			var users = []User{
				{
					Name:     "Test owner",
					Role:     "owner",
					Password: pkg.HashPassword("owner"),
					Email:    "owner@majoo.id",
				},
				{
					Name:     "Test staff",
					Role:     "staff",
					Password: pkg.HashPassword("staff"),
					Email:    "staff@majoo.id",
				},
			}
			db.Create(users)
		}
	}
}
