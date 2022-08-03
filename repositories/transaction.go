package repositories

import (
	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/upgradeskill/fp2022-crm-j-team/helpers"
	"github.com/upgradeskill/fp2022-crm-j-team/models"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

type Transaction struct {
	db *gorm.DB
}

func NewTransaction(db *gorm.DB) *Transaction {
	return &Transaction{db: db}
}

func (repository *Transaction) Create(dto *schemas.Transaction) (*models.Transaction, schemas.DatabaseError) {
	transaction := models.Transaction{
		ID:              uuid.NewString(),
		Qty:             dto.Qty,
		OutletProductId: dto.OutletProductId,
		StaffId:         dto.StaffId,
		CreatedBy:       "system",
	}

	db := repository.db.Model(&transaction)
	db.Create(&transaction)

	return &transaction, schemas.DatabaseError{}
}

func (repository *Transaction) FindById(transactionId string) (*models.Transaction, schemas.DatabaseError) {
	var transaction models.Transaction
	repository.db.Where("id = ?", transactionId).First(&transaction)

	return &transaction, schemas.DatabaseError{}
}

func (repository *Transaction) FindAll(dto *schemas.TransactionPage) (*[]models.Transaction, schemas.DatabaseError) {
	var transactions []models.Transaction

	query := repository.db.Scopes(helpers.Paginate(dto.Page, dto.PageSize))
	query.Find(&transactions)

	return &transactions, schemas.DatabaseError{}
}

func (repository *Transaction) Update(transactionId string, dto *schemas.Transaction) (*models.Transaction, schemas.DatabaseError) {
	transaction := models.Transaction{
		ID:              transactionId,
		Qty:             dto.Qty,
		OutletProductId: dto.OutletProductId,
		StaffId:         dto.StaffId,
		CreatedBy:       "system",
	}

	db := repository.db.Model(&transaction)
	db.Updates(&transaction)

	return &transaction, schemas.DatabaseError{}
}

func (repository *Transaction) Delete(transactionId string) (*models.Transaction, schemas.DatabaseError) {
	transaction := models.Transaction{
		ID: transactionId,
	}

	db := repository.db.Model(&transaction)
	db.Delete(&transaction)

	return &transaction, schemas.DatabaseError{}
}
