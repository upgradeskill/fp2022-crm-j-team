package entities

import (
	"github.com/upgradeskill/fp2022-crm-j-team/models"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

type Product interface {
	Create(input *schemas.Product) (*models.Product, schemas.DatabaseError)
}
