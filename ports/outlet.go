package ports

import (
	"github.com/upgradeskill/fp2022-crm-j-team/models"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

type Outlet interface {
	Create(input *schemas.Outlet) (*models.Outlet, schemas.DatabaseError)
	Get(input *schemas.Outlet) (*models.Outlet, schemas.DatabaseError)
	GetAll(input *schemas.Outlet) (*[]models.Outlet, schemas.DatabaseError)
	Update(input *schemas.Outlet) (*models.Outlet, schemas.DatabaseError)
	Delete(input *schemas.Outlet) (*models.Outlet, schemas.DatabaseError)
}
