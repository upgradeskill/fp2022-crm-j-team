package services

import (
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
	res, err := s.productRepository.Create(product)
	return res, err
}

/**
* ==========================================
* Service Get Multiple Product
*===========================================
 */

func (s *serviceProduct) Get(product *schemas.Product) (*models.Product, schemas.DatabaseError) {
	res, err := s.productRepository.Get(product)
	return res, err
}

/**
* ==========================================
* Service Get Multiple Product
*===========================================
 */

func (s *serviceProduct) GetAll(product *schemas.Product) (*[]models.Product, schemas.DatabaseError) {
	res, err := s.productRepository.GetAll(product)
	return res, err
}

func (s *serviceProduct) Update(product *schemas.Product) (*models.Product, schemas.DatabaseError) {
	res, err := s.productRepository.Update(product)
	return res, err
}

func (s *serviceProduct) Delete(product *schemas.Product) (*models.Product, schemas.DatabaseError) {
	res, err := s.productRepository.Delete(product)
	return res, err
}
