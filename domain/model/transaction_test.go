package model_test

import (
	"testing"

	"github.com/herculesgabriel/codepix/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestNewTransaction(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, _ := model.NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Wesley"
	account, _ := model.NewAccount(ownerName, accountNumber, bank)

	accountNumberDestination := "abcdestination"
	ownerName = "Mariana"
	accountDestination, _ := model.NewAccount(ownerName, accountNumberDestination, bank)

	kind := "email"
	key := "j@j.com"
	pixKey, _ := model.NewPixKey(kind, key, accountDestination)

	require.NotEqual(t, account.ID, accountDestination.ID)

	amount := 3.10
	statusTransaction := "pending"
	transaction, err := model.NewTransaction(account, pixKey, amount, "My description")
	//
	require.Nil(t, err)
	require.NotNil(t, uuid.FromStringOrNil(transaction.ID))
	require.Equal(t, transaction.Amount, amount)
	require.Equal(t, transaction.Status, statusTransaction)
	require.Equal(t, transaction.Description, "My description")
	require.Empty(t, transaction.CancelDescription)

	pixKeySameAccount, err := model.NewPixKey(kind, key, account)

	_, err = model.NewTransaction(account, pixKeySameAccount, amount, "My description")
	require.NotNil(t, err)

	_, err = model.NewTransaction(account, pixKey, 0, "My description")
	require.NotNil(t, err)

}

func TestModel_ChangeStatusOfATransaction(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, _ := model.NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Wesley"
	account, _ := model.NewAccount(ownerName, accountNumber, bank)

	accountNumberDestination := "abcdestination"
	ownerName = "Mariana"
	accountDestination, _ := model.NewAccount(ownerName, accountNumberDestination, bank)

	kind := "email"
	key := "j@j.com"
	pixKey, _ := model.NewPixKey(kind, key, accountDestination)

	amount := 3.10
	transaction, _ := model.NewTransaction(account, pixKey, amount, "My description")

	transaction.Complete()
	require.Equal(t, transaction.Status, model.TransactionCompleted)

	transaction.Cancel("Error")
	require.Equal(t, transaction.Status, model.TransactionError)
	require.Equal(t, transaction.CancelDescription, "Error")

}
