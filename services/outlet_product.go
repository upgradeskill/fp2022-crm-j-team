package services

import (
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
	res, err := s.outletProductRepository.Create(outletProduct)
	return res, err
}
