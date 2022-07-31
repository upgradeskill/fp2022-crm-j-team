package entities

import (
	"github.com/upgradeskill/fp2022-crm-j-team/models"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

type UserInterface interface {
	Create(input *schemas.User) (*models.User, schemas.DatabaseError)
	Get(userId string) (*models.User, schemas.DatabaseError)
	GetAll() (*[]models.User, schemas.DatabaseError)
	Delete(userId string) (bool, schemas.DatabaseError)
	Update(input *schemas.User) (*models.User, schemas.DatabaseError)
}
