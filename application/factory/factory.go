package factory

import (
	"github.com/herculesgabriel/codepix/application/usecase"
	"github.com/herculesgabriel/codepix/infrastructure/repository"
	"github.com/jinzhu/gorm"
)

func TransactionUseCaseFactory(database *gorm.DB) usecase.TransactionUseCase {
	transactionRepository := repository.TransactionRepositoryDB{DB: database}
	pixRepository := repository.PixKeyRepositoryDB{DB: database}

	transactionUseCase := usecase.TransactionUseCase{
		TransactionRepository: &transactionRepository,
		PixRepository:         &pixRepository,
	}

	return transactionUseCase
}
