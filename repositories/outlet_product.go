package repositories

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/upgradeskill/fp2022-crm-j-team/models"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

type repositoryOutletProduct struct {
	db *gorm.DB
}

func OutletProduct(db *gorm.DB) *repositoryOutletProduct {
	return &repositoryOutletProduct{db: db}
}

func (r *repositoryOutletProduct) Create(input *schemas.OutletProduct) (*models.OutletProduct, schemas.DatabaseError) {
	var outletProduct models.OutletProduct
	outletProduct.OutletId = input.OutletId
	outletProduct.ProductId = input.ProductId
	outletProduct.Price = input.Price
	outletProduct.Stock = input.Stock

	res := r.db.Debug().Create(&outletProduct)
	if res.RowsAffected < 1 {
		err := schemas.DatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &outletProduct, err
	}

	return &outletProduct, schemas.DatabaseError{}
}
