package ports

import (
	"github.com/upgradeskill/fp2022-crm-j-team/models"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

type UserRepositoryInterface interface {
	Create(input *models.User) (*models.User, schemas.DatabaseError)
	Get(userId string) (*models.User, schemas.DatabaseError)
	CheckEmailExistOnCreate(email string) (*models.User, schemas.DatabaseError)
	CheckEmailExistOnUpdate(email string, userId string) (*models.User, schemas.DatabaseError)
	Login(email string, password string) (*models.User, schemas.DatabaseError)
	GetAll() (*[]models.User, schemas.DatabaseError)
	Update(input *models.User) (*models.User, schemas.DatabaseError)
	Delete(input *models.User) (*models.User, schemas.DatabaseError)
}

type UserServiceInterface interface {
	Create(input *schemas.User) (*schemas.User, schemas.DatabaseError)
	Get(input *schemas.User) (*schemas.User, schemas.DatabaseError)
	GetAll() (*[]models.User, schemas.DatabaseError)
	Login(input *schemas.UserLogin) (interface{}, schemas.DatabaseError)
	Update(input *schemas.User) (*schemas.User, schemas.DatabaseError)
	Delete(input *schemas.User) (*schemas.User, schemas.DatabaseError)
}

type UserSeederInterface interface {
	ImportSeeder()
}
