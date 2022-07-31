package repository

import (
	"net/http"

	"github.com/upgradeskill/fp2022-crm-j-team/entities"
	"github.com/upgradeskill/fp2022-crm-j-team/models"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
	"gorm.io/gorm"
)

type repositoryUser struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) entities.UserInterface {
	return &repositoryUser{db: db}
}

func (r *repositoryUser) Create(input *schemas.User) (*models.User, schemas.DatabaseError) {
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Role:     input.Role,
		OutletId: input.OutletId,
	}

	checkEmailIsExist := r.db.Model(user).First(&user, "email = ?", input.Email)
	if checkEmailIsExist.RowsAffected > 0 {
		return &user, schemas.DatabaseError{
			Code: http.StatusUnprocessableEntity,
			Type: "user already exists",
		}
	}

	result := r.db.Model(user).Create(&user)
	if result.Error != nil {
		return &models.User{}, schemas.DatabaseError{
			Code: http.StatusUnprocessableEntity,
			Type: "error database",
		}
	}

	return &user, schemas.DatabaseError{}
}

func (r *repositoryUser) Get(userId string) (*models.User, schemas.DatabaseError) {
	var user models.User

	r.db.Raw("SELECT * FROM users WHERE id = ?", userId).Scan(&user)
	if user.ID == "" {
		return &user, schemas.DatabaseError{
			Code: http.StatusUnprocessableEntity,
			Type: "user already exists",
		}
	}

	return &user, schemas.DatabaseError{}
}

func (r *repositoryUser) GetAll() (*[]models.User, schemas.DatabaseError) {
	return &[]models.User{}, schemas.DatabaseError{}
}

func (r *repositoryUser) Delete(userId string) (bool, schemas.DatabaseError) {
	return true, schemas.DatabaseError{}
}

func (r *repositoryUser) Update(input *schemas.User) (*models.User, schemas.DatabaseError) {
	return &models.User{}, schemas.DatabaseError{}
}
