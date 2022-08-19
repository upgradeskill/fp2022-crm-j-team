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
	res, err := s.user.Login(input.Email)
	if res.ID == "" {
		return res, schemas.DatabaseError{
			Code: http.StatusUnprocessableEntity,
			Type: "user not found",
		}
	}

	checkIsPasswordVerified := pkg.CheckPasswordHash(input.Password, res.Password)
	if !checkIsPasswordVerified {
		return res, schemas.DatabaseError{
			Code: http.StatusUnprocessableEntity,
			Type: "password wrong",
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
func (s *serviceUser) Create(input *schemas.User) (*models.User, schemas.DatabaseError) {
	result, err := s.user.CheckEmailExistOnCreate(input.Email)
	if result.ID != "" {
		return result, schemas.DatabaseError{
			Code: http.StatusUnprocessableEntity,
			Type: "user already exists",
		}
	}

	if input.Role != "owner" && input.Role != "staff" {
		return result, schemas.DatabaseError{
			Code: http.StatusUnprocessableEntity,
			Type: "role must between 'owner' or 'staff'",
		}
	}

	user := models.User{
		Name:      input.Name,
		Email:     input.Email,
		Password:  pkg.HashPassword(input.Password),
		Role:      input.Role,
		OutletId:  input.OutletId,
		CreatedBy: helpers.SessionUser().ID,
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

	if input.Password != "" {
		input.Password = pkg.HashPassword(input.Password)
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
		DeletedBy: helpers.SessionUser().ID,
	}

	res, err := s.user.Delete(&user)
	return res, err
}

/**
* ==========================================
* Service Get User by user id
*===========================================
 */
func (s *serviceUser) Get(input *schemas.User) (*models.User, schemas.DatabaseError) {
	res, err := s.user.Get(input.ID)
	if res.ID == "" {
		return res, schemas.DatabaseError{
			Code: http.StatusNotFound,
			Type: "user not found",
		}
	}

	return res, err
}

/**
* ==========================================
* Service Get all user
*===========================================
 */
func (s *serviceUser) GetAll(input *schemas.User) (*[]models.User, schemas.DatabaseError) {
	res, err := s.user.GetAll(input)
	return res, err
}
