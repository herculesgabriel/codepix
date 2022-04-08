package usecase

import "github.com/herculesgabriel/codepix/domain/model"

type TransactionUseCase struct {
	TransactionRepository model.ITransactionRepository
	PixRepository         model.IPixKeyRepository
}

func (t *TransactionUseCase) Register(accountID string, pixKeyToID string, pixKeyToKind string, amount float64, description string) (*model.Transaction, error) {
	account, err := t.PixRepository.FindAccount(accountID)
	if err != nil {
		return nil, err
	}

	pixKey, err := t.PixRepository.FindByKeyAndKind(pixKeyToID, pixKeyToKind)
	if err != nil {
		return nil, err
	}

	transaction, err := model.NewTransaction(account, pixKey, amount, description)
	if err != nil {
		return nil, err
	}

	err = t.TransactionRepository.Register(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *TransactionUseCase) Confirm(transactionId string) (*model.Transaction, error) {
	transaction, err := t.TransactionRepository.Find(transactionId)
	if err != nil {
		return nil, err
	}

	transaction.Status = model.TransactionConfirmed

	err = t.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *TransactionUseCase) Complete(transactionId string) (*model.Transaction, error) {
	transaction, err := t.TransactionRepository.Find(transactionId)
	if err != nil {
		return nil, err
	}

	transaction.Status = model.TransactionCompleted

	err = t.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *TransactionUseCase) Error(transactionId string, reason string) (*model.Transaction, error) {
	transaction, err := t.TransactionRepository.Find(transactionId)
	if err != nil {
		return nil, err
	}

	transaction.Status = model.TransactionError
	transaction.CancelDescription = reason

	err = t.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
