package repositories

import (
	"github.com/upgradeskill/fp2022-crm-j-team/models"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
	"gorm.io/gorm"
)

type repositoryCategory struct {
	db *gorm.DB
}

func Category(db *gorm.DB) *repositoryCategory {
	return &repositoryCategory{db: db}
}

func (r *repositoryCategory) Create(input *schemas.Category) (*models.Category, schemas.DatabaseError) {
	var category models.Category
	category.ID = input.ID
	category.Name = input.Name
	category.Description = input.Description
	category.CreatedBy = "system"

	db := r.db.Model(&category)

	db.Create(&category)

	return &category, schemas.DatabaseError{}
}

func (r *repositoryCategory) Get(input *schemas.Category) (*models.Category, schemas.DatabaseError) {
	var category models.Category
	category.ID = input.ID

	db := r.db.Model(&category)

	db.Debug().First(&category)

	return &category, schemas.DatabaseError{}
}
