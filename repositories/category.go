package repositories

import (
	"net/http"
	"time"

	"github.com/upgradeskill/fp2022-crm-j-team/helpers"
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

	db.Debug().Create(&category)

	return &category, schemas.DatabaseError{}
}

/**
* ==========================================
* Repository Get All Category
*===========================================
 */

func (r *repositoryCategory) GetAll(input *schemas.Category) (*[]models.Category, schemas.DatabaseError) {
	var category []models.Category

	db := r.db.Model(&category)

	if input.Name != "" {
		db.Where("name like ?", "%"+input.Name+"%")
	}

	db.Scopes(helpers.Paginate(input.Page, input.PageSize))
	db.Debug().Find(&category)

	if db.RowsAffected < 1 {
		err := schemas.DatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &category, err
	}

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

	if db.RowsAffected < 1 {
		err := schemas.DatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &category, err
	}

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
	category.UpdatedAt = time.Now()
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
