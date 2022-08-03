package services

import (
	"github.com/google/uuid"
	"github.com/upgradeskill/fp2022-crm-j-team/models"
	"github.com/upgradeskill/fp2022-crm-j-team/ports"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

type serviceOutlet struct {
	outletRepository ports.Outlet
}

func Outlet(outletRepository ports.Outlet) *serviceOutlet {
	return &serviceOutlet{outletRepository: outletRepository}
}

/**
* ==========================================
* Service Create New Outlet
*===========================================
 */

func (s *serviceOutlet) Create(outlet *schemas.Outlet) (*models.Outlet, schemas.DatabaseError) {
	outlet.ID = uuid.NewString()
	res, err := s.outletRepository.Create(outlet)
	return res, err
}

/**
* ==========================================
* Service Update Outlet By ID Outlet
*===========================================
 */

func (s *serviceOutlet) Update(outlet *schemas.Outlet) (*models.Outlet, schemas.DatabaseError) {
	res, err := s.outletRepository.Update(outlet)
	return res, err
}

/**
* ==========================================
* Service Result Outlet By ID Outlet
*===========================================
 */

func (s *serviceOutlet) Get(input *schemas.Outlet) (*models.Outlet, schemas.DatabaseError) {
	var outlet schemas.Outlet
	outlet.ID = input.ID

	res, err := s.outletRepository.Get(&outlet)
	return res, err
}

/**
* ==========================================
* Service Result All Outlet
*===========================================
 */

func (s *serviceOutlet) GetAll(outlet *schemas.Outlet) (*[]models.Outlet, schemas.DatabaseError) {
	res, err := s.outletRepository.GetAll(outlet)
	return res, err
}

/**
* ==========================================
* Service Delete Outlet By ID Outlet
*===========================================
 */

func (s *serviceOutlet) Delete(outlet *schemas.Outlet) (*models.Outlet, schemas.DatabaseError) {
	res, err := s.outletRepository.Delete(outlet)
	return res, err
}
