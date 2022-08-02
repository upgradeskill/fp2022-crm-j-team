package repositories

import (
	"time"

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

/**
* ==========================================
* Repository Create New Category
*===========================================
 */

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

/**
* ==========================================
* Repository Get All Category
*===========================================
 */

func (r *repositoryCategory) GetAll() (*[]models.Category, schemas.DatabaseError) {
	var category []models.Category

	db := r.db.Model(&category)

	db.Debug().Find(&category)

	return &category, schemas.DatabaseError{}
}

/**
* ==========================================
* Repository Get Category By ID Category
*===========================================
 */

func (r *repositoryCategory) Get(input *schemas.Category) (*models.Category, schemas.DatabaseError) {
	var category models.Category
	category.ID = input.ID

	db := r.db.Model(&category)

	db.Debug().First(&category)

	return &category, schemas.DatabaseError{}
}

/**
* ==========================================
* Repository Update Category By ID Category
*===========================================
 */

func (r *repositoryCategory) Update(input *schemas.Category) (*models.Category, schemas.DatabaseError) {
	var category models.Category
	category.ID = input.ID
	category.Name = input.Name
	category.Description = input.Description
	category.CreatedAt = time.Now()
	category.UpdatedBy = "system"

	db := r.db.Model(&category)

	db.Debug().Updates(&category)

	return &category, schemas.DatabaseError{}
}

/**
* ==========================================
* Repository Delete Category By ID Category
*===========================================
 */

func (r *repositoryCategory) Delete(input *schemas.Category) (*models.Category, schemas.DatabaseError) {
	var category models.Category
	category.ID = input.ID
	category.DeletedBy = "system"

	db := r.db.Model(&category)

	db.Debug().Delete(&category)

	return &category, schemas.DatabaseError{}
}
