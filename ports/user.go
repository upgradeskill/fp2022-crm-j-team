package ports

import (
	"github.com/upgradeskill/fp2022-crm-j-team/models"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

type UserInterface interface {
	Create(input *models.User) (*models.User, schemas.DatabaseError)
	Get(userId string) (*models.User, schemas.DatabaseError)
	CheckEmailExistOnCreate(email string) (*models.User, schemas.DatabaseError)
	CheckEmailExistOnUpdate(email string, userId string) (*models.User, schemas.DatabaseError)
	GetAll() (*[]models.User, schemas.DatabaseError)
	Update(input *models.User) (*models.User, schemas.DatabaseError)
}
