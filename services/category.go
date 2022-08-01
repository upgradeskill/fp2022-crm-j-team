package services

import (
	"github.com/google/uuid"
	"github.com/upgradeskill/fp2022-crm-j-team/models"
	"github.com/upgradeskill/fp2022-crm-j-team/ports"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

type serviceCategory struct {
	categoryRepository ports.Category
}

func Category(categoryRepository ports.Category) *serviceCategory {
	return &serviceCategory{categoryRepository: categoryRepository}
}

/**
* ==========================================
* Service Create New Category
*===========================================
 */

func (s *serviceCategory) Create(category *schemas.Category) (*models.Category, schemas.DatabaseError) {
	category.ID = uuid.NewString()
	res, err := s.categoryRepository.Create(category)
	return res, err
}

/**
* ==========================================
* Service Update Category By ID Category
*===========================================
 */

func (s *serviceCategory) Update(category *schemas.Category) (*models.Category, schemas.DatabaseError) {
	res, err := s.categoryRepository.Update(category)
	return res, err
}

/**
* ==========================================
* Service Result Category By ID Category
*===========================================
 */

func (s *serviceCategory) Get(input *schemas.Category) (*models.Category, schemas.DatabaseError) {
	var category schemas.Category
	category.ID = input.ID

	res, err := s.categoryRepository.Get(&category)
	return res, err
}

/**
* ==========================================
* Service Delete Category By ID Category
*===========================================
 */

func (s *serviceCategory) Delete(category *schemas.Category) (*models.Category, schemas.DatabaseError) {
	res, err := s.categoryRepository.Delete(category)
	return res, err
}
