package ports

import (
	"github.com/upgradeskill/fp2022-crm-j-team/models"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

type Category interface {
	Create(input *schemas.Category) (*models.Category, schemas.DatabaseError)
	Get(input *schemas.Category) (*models.Category, schemas.DatabaseError)
	GetAll(input *schemas.Category) (*[]models.Category, schemas.DatabaseError)
	Update(input *schemas.Category) (*models.Category, schemas.DatabaseError)
	Delete(input *schemas.Category) (*models.Category, schemas.DatabaseError)
}
