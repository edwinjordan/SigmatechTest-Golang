package transaction_repository

import (
	"context"

	"github.com/edwinjordan/SigmatechTest-Golang/app/repository"
	"github.com/edwinjordan/SigmatechTest-Golang/entity"
	"github.com/edwinjordan/SigmatechTest-Golang/pkg/helpers"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	DB *gorm.DB
}

func New(db *gorm.DB) repository.TransactionRepository {
	return &TransactionRepositoryImpl{
		DB: db,
	}
}

// Create implements repository.TransactionRepository.
func (repo *TransactionRepositoryImpl) Create(ctx context.Context, transaction entity.Transaction) entity.TransactionResponse {
	transactionData := &Transaction{}
	transactionData = transactionData.FromEntity(&transaction)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).Create(&transactionData).Error
	helpers.PanicIfError(err)

	return *transactionData.ToEntity()
}

// FindAll implements repository.TransactionRepository.
