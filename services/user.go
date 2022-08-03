package services

import (
	"net/http"

	"github.com/upgradeskill/fp2022-crm-j-team/helpers"
	"github.com/upgradeskill/fp2022-crm-j-team/models"
	"github.com/upgradeskill/fp2022-crm-j-team/pkg"
	"github.com/upgradeskill/fp2022-crm-j-team/ports"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

type serviceUser struct {
	user ports.UserRepositoryInterface
}

func NewServiceUser(user ports.UserRepositoryInterface) *serviceUser {
	return &serviceUser{user: user}
}

/**
* ==========================================
* Service Login
*===========================================
 */
func (s *serviceUser) Login(input *schemas.UserLogin) (interface{}, schemas.DatabaseError) {
	res, err := s.user.Login(input.Email, input.Password)
	if res.ID == "" {
		return res, schemas.DatabaseError{
			Code: http.StatusUnprocessableEntity,
			Type: "user not found",
		}
	}

	token, _ := pkg.GenerateTokenJwt(res.ID, res.Email, res.Role, res.OutletId)
	return map[string]string{
		"id":        res.ID,
		"email":     res.Email,
		"role":      res.Role,
		"outlet_id": res.OutletId,
		"token":     token,
	}, err
}

/**
* ==========================================
* Service Create User
*===========================================
 */
func (s *serviceUser) Create(input *schemas.User) (*schemas.User, schemas.DatabaseError) {
	result, err := s.user.CheckEmailExistOnCreate(input.Email)
	if result.ID != "" {
		return nil, schemas.DatabaseError{
			Code: http.StatusUnprocessableEntity,
			Type: "user already exists",
		}
	}

	user := models.User{
		Name:      input.Name,
		Email:     input.Email,
		Password:  input.Password,
		Role:      input.Role,
		OutletId:  input.OutletId,
		CreatedBy: helpers.SessionUser().ID,
	}

	res, err := s.user.Create(&user)
	input.ID = res.ID
	return input, err
}

/**
* ==========================================
* Service Update User
*===========================================
 */
func (s *serviceUser) Update(input *schemas.User) (*schemas.User, schemas.DatabaseError) {
	notfound, _ := s.user.Get(input.ID)
	if notfound.ID == "" {
		return nil, schemas.DatabaseError{
			Code: http.StatusNotFound,
			Type: "user not found",
		}
	}

	exist, _ := s.user.CheckEmailExistOnUpdate(input.Email, input.ID)
	if exist.ID != "" {
		return nil, schemas.DatabaseError{
			Code: http.StatusUnprocessableEntity,
			Type: "user already exists",
		}
	}

	user := models.User{
		ID:        input.ID,
		Name:      input.Name,
		Email:     input.Email,
		Password:  input.Password,
		Role:      input.Role,
		OutletId:  input.OutletId,
		UpdatedBy: helpers.SessionUser().ID,
	}

	_, err := s.user.Update(&user)
	return input, err
}

/**
* ==========================================
* Service Soft Delete User
*===========================================
 */
func (s *serviceUser) Delete(input *schemas.User) (*schemas.User, schemas.DatabaseError) {
	found, _ := s.user.Get(input.ID)
	if found.ID == "" {
		return nil, schemas.DatabaseError{
			Code: http.StatusNotFound,
			Type: "user not found",
		}
	}

	user := models.User{
		ID:        input.ID,
		DeletedBy: helpers.SessionUser().ID,
	}

	res, err := s.user.Delete(&user)
	input.ID = res.ID
	input.Email = res.Email
	input.OutletId = res.OutletId
	input.Role = res.Role
	input.Name = res.Name
	return input, err
}

/**
* ==========================================
* Service Get User by user id
*===========================================
 */
func (s *serviceUser) Get(input *schemas.User) (*schemas.User, schemas.DatabaseError) {
	res, err := s.user.Get(input.ID)
	if res.ID == "" {
		return nil, schemas.DatabaseError{
			Code: http.StatusNotFound,
			Type: "user not found",
		}
	}

	input.ID = res.ID
	input.Email = res.Email
	input.OutletId = res.OutletId
	input.Role = res.Role
	input.Name = res.Name
	return input, err
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
