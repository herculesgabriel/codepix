package model_test

import (
	"testing"

	"github.com/herculesgabriel/codepix/domain/model"
	uuid "github.com/satori/go.uuid"

	"github.com/stretchr/testify/require"
)

func TestModel_NewPixKey(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, err := model.NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Wesley"
	account, err := model.NewAccount(ownerName, accountNumber, bank)

	kind := "email"
	key := "j@j.com"
	pixKey, err := model.NewPixKey(kind, key, account)

	require.NotEmpty(t, uuid.FromStringOrNil(pixKey.ID))
	require.Equal(t, pixKey.Kind, kind)
	require.Equal(t, pixKey.Status, "active")

	kind = "cpf"
	_, err = model.NewPixKey(kind, key, account)
	require.Nil(t, err)

	_, err = model.NewPixKey("nome", key, account)
	require.NotNil(t, err)
}
