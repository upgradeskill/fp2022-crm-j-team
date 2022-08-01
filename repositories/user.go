package repositories

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

func (r *repositoryUser) Create(user *models.User) (*models.User, schemas.DatabaseError) {
	result := r.db.Model(user).Create(&user)
	if result.Error != nil {
		return &models.User{}, schemas.DatabaseError{
			Code: http.StatusUnprocessableEntity,
			Type: "error database",
		}
	}

	return user, schemas.DatabaseError{}
}

func (r *repositoryUser) Update(user *models.User) (*models.User, schemas.DatabaseError) {
	result := r.db.Model(user).Updates(&user)
	if result.Error != nil {
		return &models.User{}, schemas.DatabaseError{
			Code: http.StatusUnprocessableEntity,
			Type: "error database",
		}
	}

	return user, schemas.DatabaseError{}
}

func (r *repositoryUser) Delete(user *models.User) (*models.User, schemas.DatabaseError) {
	result := r.db.Model(user).Delete(&user)
	if result.Error != nil {
		return &models.User{}, schemas.DatabaseError{
			Code: http.StatusUnprocessableEntity,
			Type: "error database",
		}
	}

	return user, schemas.DatabaseError{}
}

func (r *repositoryUser) Get(userId string) (*models.User, schemas.DatabaseError) {
	var user models.User

	r.db.Where("id = ?", userId).Find(&user)
	if user.ID == "" {
		return &user, schemas.DatabaseError{
			Code: http.StatusNotFound,
			Type: "user not found",
		}
	}

	return &user, schemas.DatabaseError{}
}

func (r *repositoryUser) CheckEmailExistOnCreate(email string) (*models.User, schemas.DatabaseError) {
	var user models.User

	r.db.Where("email = ?", email).Find(&user)
	return &user, schemas.DatabaseError{}
}

func (r *repositoryUser) CheckEmailExistOnUpdate(email string, userId string) (*models.User, schemas.DatabaseError) {
	var user models.User

	r.db.Where("email = ? and id != ?", email, userId).Find(&user)
	return &user, schemas.DatabaseError{}
}

func (r *repositoryUser) GetAll() (*[]models.User, schemas.DatabaseError) {
	var user []models.User

	r.db.Find(&user)
	return &user, schemas.DatabaseError{}
}
