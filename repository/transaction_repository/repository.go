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

// FindAll implements repository.TransactionRepository.
func (repo *TransactionRepositoryImpl) FindAll(ctx context.Context, where map[string]interface{}) []map[string]interface{} {
	data := []map[string]interface{}{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).
		Table("tb_transaction").
		Select("tb_transaction.transaction_id as TransactionID,tb_perusahaan_asset.perusahaan_asset_otr_price as OTRPrice, tb_perusahaan.perusahaan_fee as AdminFee, tb_tenor.tenor as Tenor, tb_tenor.tenor_interest as Interest,tb_perusahaan_asset.perusahaan_asset_nama as AssetName,(tb_perusahaan_asset.perusahaan_asset_otr_price * tb_tenor.tenor_interest / 100) as Interes, (tb_perusahaan_asset.perusahaan_asset_otr_price + tb_perusahaan.perusahaan_fee) as TotalPrice").
		Joins("LEFT JOIN tb_tenor ON tb_transaction.transaction_tenor_id = tb_tenor.tenor_id").
		Joins("LEFT JOIN tb_perusahaan_asset ON tb_transaction.transaction_perusahaan_asset_id = tb_perusahaan_asset.perusahaan_asset_id").
		Joins("LEFT JOIN tb_perusahaan ON tb_perusahaan_asset.perusahaan_id = tb_perusahaan.perusahaan_id").
		Where(where).Find(&data).Error

	helpers.PanicIfError(err)

	return data
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
