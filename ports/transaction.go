package ports

import (
	"github.com/upgradeskill/fp2022-crm-j-team/models"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

type Transaction interface {
	Create(dto *schemas.Transaction) (*models.Transaction, schemas.DatabaseError)
	FindById(transactionId string) (*models.Transaction, schemas.DatabaseError)
	FindAll(dto *schemas.TransactionPage) (*[]models.Transaction, schemas.DatabaseError)
	Update(transactionId string, dto *schemas.Transaction) (*models.Transaction, schemas.DatabaseError)
	Delete(transactionId string) (*models.Transaction, schemas.DatabaseError)
}
