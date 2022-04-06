package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

const (
	TransactionPending   string = "pending"
	TransactionConfirmed string = "confirmed"
	TransactionCompleted string = "completed"
	TransactionError     string = "error"
)

type ITransactionRepository interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	Find(id string) (*Transaction, error)
}

type Transactions struct {
	Transaction []Transaction
}

type Transaction struct {
	Base              `valid:"required"`
	AccountFrom       *Account `valid:"-"`
	PixKeyTo          *PixKey  `valid:"-"`
	Amount            float64  `json:"amount" valid:"notnull"`
	Status            string   `json:"status" valid:"notnull"`
	Description       string   `json:"description" valid:"notnull"`
	CancelDescription string   `json:"cancel_description" valid:"notnull"`
}

func (t *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(t)
	if err != nil {
		return err
	}

	if t.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}
	if t.Status != TransactionPending && t.Status != TransactionConfirmed && t.Status != TransactionCompleted && t.Status != TransactionError {
		return errors.New("invalid status")
	}
	if t.PixKeyTo.AccountID == t.AccountFrom.ID {
		return errors.New("source and destination account cannot be the same")
	}

	return nil
}

func NewTransaction(accountFrom *Account, pixKeyTo *PixKey, amount float64, description string) (*Transaction, error) {
	transaction := Transaction{
		AccountFrom: accountFrom,
		PixKeyTo:    pixKeyTo,
		Amount:      amount,
		Description: description,
		Status:      TransactionPending,
	}

	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()

	err := transaction.isValid()
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (t *Transaction) Confirm() error {
	t.Status = TransactionConfirmed
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

func (t *Transaction) Complete() error {
	t.Status = TransactionCompleted
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

func (t *Transaction) Cancel(description string) error {
	t.Status = TransactionError
	t.Description = description
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}
