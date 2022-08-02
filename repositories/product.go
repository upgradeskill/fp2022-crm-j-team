package repositories

import (
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
	product.CreatedBy = "system" // temporary hardcode
	db := r.db.Model(&product)

	db.Create(&product)

	return &product, schemas.DatabaseError{}
}

func (r *repositoryProduct) Get(input *schemas.Product) (*models.Product, schemas.DatabaseError) {
	var product models.Product
	r.db.Where("id = ?", input.ID).First(&product)

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
