package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Account struct {
	Base      `valid:"required"`
	OwnerName string    `json:"owner_name" valid:"notnull"`
	Bank      *Bank     `valid:"-"`
	Number    string    `json:"number" valid:"notnull"`
	PixKeys   []*PixKey `valid:"-"`
}

func (a *Account) isValid() error {
	_, err := govalidator.ValidateStruct(a)
	if err != nil {
		return err
	}
	return nil
}

func NewAccount(ownerName string, number string, bank *Bank) (*Account, error) {
	account := Account{
		OwnerName: ownerName,
		Number:    number,
		Bank:      bank,
	}

	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()
	account.UpdatedAt = time.Now()

	err := account.isValid()
	if err != nil {
		return nil, err
	}
	return &account, nil
}
