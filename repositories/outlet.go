package repositories

import (
	"time"

	"github.com/upgradeskill/fp2022-crm-j-team/helpers"
	"github.com/upgradeskill/fp2022-crm-j-team/models"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
	"gorm.io/gorm"
)

type repositoryOutlet struct {
	db *gorm.DB
}

func Outlet(db *gorm.DB) *repositoryOutlet {
	return &repositoryOutlet{db: db}
}

/**
* ==========================================
* Repository Create New Outlet
*===========================================
 */

func (r *repositoryOutlet) Create(input *schemas.Outlet) (*models.Outlet, schemas.DatabaseError) {
	var outlet models.Outlet
	outlet.ID = input.ID
	outlet.Name = input.Name
	outlet.Phone = input.Phone
	outlet.Address = input.Address
	outlet.CreatedBy = "system"

	db := r.db.Model(&outlet)

	db.Debug().Create(&outlet)

	return &outlet, schemas.DatabaseError{}
}

/**
* ==========================================
* Repository Get All Outlet
*===========================================
 */

func (r *repositoryOutlet) GetAll(input *schemas.Outlet) (*[]models.Outlet, schemas.DatabaseError) {
	var outlet []models.Outlet

	db := r.db.Model(&outlet)

	if input.Name != "" {
		db.Where("name like ?", "%"+input.Name+"%")
	}

	db.Scopes(helpers.Paginate(input.Page, input.PageSize))
	db.Debug().Find(&outlet)

	return &outlet, schemas.DatabaseError{}
}

/**
* ==========================================
* Repository Get Outlet By ID Outlet
*===========================================
 */

func (r *repositoryOutlet) Get(input *schemas.Outlet) (*models.Outlet, schemas.DatabaseError) {
	var outlet models.Outlet
	outlet.ID = input.ID

	db := r.db.Model(&outlet)

	db.Debug().First(&outlet)

	return &outlet, schemas.DatabaseError{}
}

/**
* ==========================================
* Repository Update Outlet By ID Outlet
*===========================================
 */

func (r *repositoryOutlet) Update(input *schemas.Outlet) (*models.Outlet, schemas.DatabaseError) {
	var outlet models.Outlet
	outlet.ID = input.ID
	outlet.Name = input.Name
	outlet.Phone = input.Phone
	outlet.Address = input.Address
	outlet.UpdatedAt = time.Now()
	outlet.UpdatedBy = "system"

	db := r.db.Model(&outlet)

	db.Debug().Updates(&outlet)

	return &outlet, schemas.DatabaseError{}
}

/**
* ==========================================
* Repository Delete Outlet By ID Outlet
*===========================================
 */

func (r *repositoryOutlet) Delete(input *schemas.Outlet) (*models.Outlet, schemas.DatabaseError) {
	var outlet models.Outlet
	outlet.ID = input.ID
	outlet.DeletedBy = "system"

	db := r.db.Model(&outlet)

	db.Debug().Delete(&outlet)

	return &outlet, schemas.DatabaseError{}
}
