package services

import (
	"net/http"
	"time"

	"github.com/upgradeskill/fp2022-crm-j-team/entities"
	"github.com/upgradeskill/fp2022-crm-j-team/models"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

type serviceUser struct {
	user entities.UserInterface
}

func NewServiceUser(user entities.UserInterface) *serviceUser {
	return &serviceUser{user: user}
}

/**
* ==========================================
* Service Create User
*===========================================
 */
func (s *serviceUser) Create(input *schemas.User) (*models.User, schemas.DatabaseError) {
	result, err := s.user.CheckEmailExistOnCreate(input.Email)
	if result.ID != "" {
		return result, schemas.DatabaseError{
			Code: http.StatusUnprocessableEntity,
			Type: "user already exists",
		}
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Role:     input.Role,
		OutletId: input.OutletId,
	}

	res, err := s.user.Create(&user)
	return res, err
}

/**
* ==========================================
* Service Update User
*===========================================
 */
func (s *serviceUser) Update(input *schemas.User) (*models.User, schemas.DatabaseError) {
	notfound, err := s.user.Get(input.ID)
	if notfound.ID == "" {
		return notfound, schemas.DatabaseError{
			Code: http.StatusNotFound,
			Type: "user not found",
		}
	}

	exist, err := s.user.CheckEmailExistOnUpdate(input.Email, input.ID)
	if exist.ID != "" {
		return exist, schemas.DatabaseError{
			Code: http.StatusUnprocessableEntity,
			Type: "user already exists",
		}
	}

	user := models.User{
		ID:       input.ID,
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Role:     input.Role,
		OutletId: input.OutletId,
	}

	res, err := s.user.Update(&user)
	return res, err
}

/**
* ==========================================
* Service Soft Delete User
*===========================================
 */
func (s *serviceUser) Delete(input *schemas.User) (*models.User, schemas.DatabaseError) {
	found, err := s.user.Get(input.ID)
	if found.ID == "" {
		return found, schemas.DatabaseError{
			Code: http.StatusNotFound,
			Type: "user not found",
		}
	}

	user := models.User{
		ID:        input.ID,
		DeletedAt: time.Now(),
	}

	res, err := s.user.Update(&user)
	return res, err
}

/**
* ==========================================
* Service Get User by user id
*===========================================
 */
func (s *serviceUser) Get(input *schemas.User) (*models.User, schemas.DatabaseError) {
	res, err := s.user.Get(input.ID)
	return res, err
}

/**
* ==========================================
* Service Get all user
*===========================================
 */
func (s *serviceUser) GetAll() (*[]models.User, schemas.DatabaseError) {
	res, err := s.user.GetAll()
	return res, err
}
