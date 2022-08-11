package repositories

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/upgradeskill/fp2022-crm-j-team/helpers"
	"github.com/upgradeskill/fp2022-crm-j-team/models"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

type repositoryProduct struct {
	db *gorm.DB
}

func Product(db *gorm.DB) *repositoryProduct {
	return &repositoryProduct{db: db}
}

func (r *repositoryProduct) Create(input *schemas.Product) (*models.Product, schemas.DatabaseError) {
	var product models.Product
	product.ID = input.ID
	product.Name = input.Name
	product.SKU = input.SKU
	product.CategoryId = input.CategoryId
	product.CreatedBy = helpers.SessionUser().ID

	res := r.db.Debug().Create(&product)
	if res.RowsAffected < 1 {
		err := schemas.DatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &product, err
	}

	return &product, schemas.DatabaseError{}
}

func (r *repositoryProduct) Get(input *schemas.Product) (*models.Product, schemas.DatabaseError) {
	product := models.Product{}

	res := r.db.Debug().Where("id = ?", input.ID).First(&product)

	if res.RowsAffected < 1 {
		err := schemas.DatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &product, err
	}

	return &product, schemas.DatabaseError{}
}

func (r *repositoryProduct) GetAll(input *schemas.Product) (*[]models.Product, schemas.DatabaseError) {
	var products []models.Product

	query := r.db.Scopes(helpers.Paginate(input.Page, input.PageSize))

	if input.ID != "" {
		query.Where("id like ?", "%"+input.ID+"%")
	}

	if input.Name != "" {
		query.Where("name like ?", "%"+input.Name+"%")
	}

	if input.CategoryId != "" {
		query.Where("category_id like ?", "%"+input.CategoryId+"%")
	}

	query.Find(&products)

	return &products, schemas.DatabaseError{}
}

func (r *repositoryProduct) Update(input *schemas.Product) (*models.Product, schemas.DatabaseError) {
	var product models.Product
	product.ID = input.ID
	product.Name = input.Name
	product.SKU = input.SKU
	product.CategoryId = input.CategoryId
	product.UpdatedBy = helpers.SessionUser().ID

	res := r.db.Debug().Updates(&product)

	if res.RowsAffected < 1 {
		err := schemas.DatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &product, err
	}

	return &product, schemas.DatabaseError{}
}

func (r *repositoryProduct) Delete(input *schemas.Product) (*models.Product, schemas.DatabaseError) {
	product := models.Product{}

	res := r.db.Debug().Where("id = ?", input.ID).Delete(&product)

	if res.RowsAffected < 1 {
		err := schemas.DatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &product, err
	}

	return &product, schemas.DatabaseError{}
}
