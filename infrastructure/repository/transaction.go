package repository

import (
	"fmt"

	"github.com/herculesgabriel/codepix/domain/model"
	"github.com/jinzhu/gorm"
)

type TransactionRepositoryDB struct {
	DB *gorm.DB
}

func (t *TransactionRepositoryDB) Register(transaction *model.Transaction) error {
	err := t.DB.Create(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *TransactionRepositoryDB) Save(transaction *model.Transaction) error {
	err := t.DB.Save(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *TransactionRepositoryDB) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction
	t.DB.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("transaction was not found")
	}
	return &transaction, nil
}
