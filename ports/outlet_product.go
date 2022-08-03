package ports

import (
	"github.com/upgradeskill/fp2022-crm-j-team/models"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

type OutletProduct interface {
	Create(input *schemas.OutletProduct) (*models.OutletProduct, schemas.DatabaseError)
}
