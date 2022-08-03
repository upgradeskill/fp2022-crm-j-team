package services

import (
	"github.com/google/uuid"
	"github.com/upgradeskill/fp2022-crm-j-team/models"
	"github.com/upgradeskill/fp2022-crm-j-team/ports"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

type serviceOutletProduct struct {
	outletProductRepository ports.OutletProduct
}

func OutletProduct(outletProductRepository ports.OutletProduct) *serviceOutletProduct {
	return &serviceOutletProduct{outletProductRepository: outletProductRepository}
}

func (s *serviceOutletProduct) Create(outletProduct *schemas.OutletProduct) (*models.OutletProduct, schemas.DatabaseError) {
	outletProduct.ID = uuid.NewString()
	res, err := s.outletProductRepository.Create(outletProduct)
	return res, err
}
