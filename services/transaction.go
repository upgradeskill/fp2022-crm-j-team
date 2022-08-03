package services

import (
	"github.com/upgradeskill/fp2022-crm-j-team/models"
	"github.com/upgradeskill/fp2022-crm-j-team/ports"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

type Transaction struct {
	repository ports.Transaction
}

func NewTransaction(repository ports.Transaction) *Transaction {
	return &Transaction{repository: repository}
}

func (service *Transaction) Create(dto *schemas.Transaction) (*models.Transaction, schemas.DatabaseError) {
	res, err := service.repository.Create(dto)
	return res, err
}

func (service *Transaction) FindById(transactionId string) (*models.Transaction, schemas.DatabaseError) {
	res, err := service.repository.FindById(transactionId)
	return res, err
}

func (service *Transaction) FindAll(dto *schemas.TransactionPage) (*[]models.Transaction, schemas.DatabaseError) {
	res, err := service.repository.FindAll(dto)
	return res, err
}

func (service *Transaction) Update(transactionId string, dto *schemas.Transaction) (*models.Transaction, schemas.DatabaseError) {
	res, err := service.repository.Update(transactionId, dto)
	return res, err
}

func (service *Transaction) Delete(transactionId string) (*models.Transaction, schemas.DatabaseError) {
	res, err := service.repository.Delete(transactionId)
	return res, err
}
