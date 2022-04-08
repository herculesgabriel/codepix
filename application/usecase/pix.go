package usecase

import "github.com/herculesgabriel/codepix/domain/model"

// type IPixKeyRepository interface {
// 	Register(pixKey *PixKey) (*PixKey, error)
// 	FindByKeyAndKind(key string, kind string) (*PixKey, error)
// 	AddBank(bank *Bank) error
// 	AddAccount(account *Account) error
// 	FindAccount(id string) (*Account, error)
// 	FindBank(id string) (*Bank, error)
// }

type PixUseCase struct {
	PixKeyRepository model.IPixKeyRepository
}

func (p *PixUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error) {
	account, err := p.PixKeyRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, key, account)
	if err != nil {
		return nil, err
	}

	_, err = p.PixKeyRepository.Register(pixKey)
	if err != nil {
		return nil, err
	}

	return pixKey, nil
}

func (p *PixUseCase) FindByKeyAndKind(key string, kind string) (*model.PixKey, error) {
	pixKey, err := p.FindByKeyAndKind(key, kind)
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}
