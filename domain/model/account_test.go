package model_test

import (
	"testing"

	"github.com/herculesgabriel/codepix/domain/model"
	uuid "github.com/satori/go.uuid"

	"github.com/stretchr/testify/require"
)

func TestModel_NewAccount(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, err := model.NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Wesley"
	account, err := model.NewAccount(ownerName, accountNumber, bank)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(account.ID))
	require.Equal(t, account.Number, accountNumber)
	require.Equal(t, account.BankID, bank.ID)

	_, err = model.NewAccount(ownerName, "", bank)
	require.NotNil(t, err)
}
