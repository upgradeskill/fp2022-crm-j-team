package services

import (
	"github.com/google/uuid"
	"github.com/upgradeskill/fp2022-crm-j-team/models"
	"github.com/upgradeskill/fp2022-crm-j-team/ports"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

type serviceProduct struct {
	productRepository ports.Product
}

func Product(productRepository ports.Product) *serviceProduct {
	return &serviceProduct{productRepository: productRepository}
}

/**
* ==========================================
* Service Create New Product
*===========================================
 */

func (s *serviceProduct) Create(product *schemas.Product) (*models.Product, schemas.DatabaseError) {
	product.ID = uuid.NewString()
	res, err := s.productRepository.Create(product)
	return res, err
}
