package transaction_repository

import (
	"time"

	"github.com/edwinjordan/SigmatechTest-Golang/entity"
	"github.com/edwinjordan/SigmatechTest-Golang/pkg/helpers"
	"gorm.io/gorm"
)

type Transaction struct {
	TransactionId                string    `gorm:"column:transaction_id"`
	TransactionUserId            string    `gorm:"column:transaction_user_id"`
	TransactionTenorId           string    `gorm:"column:transaction_tenor_id"`
	TransactionPerusahaanAssetId string    `gorm:"column:transaction_perusahaan_asset_id"`
	TransactionCreateAt          time.Time `gorm:"column:transaction_create_at"`
	TransactionUpdateAt          time.Time `gorm:"column:transaction_update_at"`
}

func (Transaction) TableName() string {
	return "tb_transaction"
}

func (model *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	model.TransactionId = helpers.GenUUID()

	model.TransactionCreateAt = helpers.CreateDateTime()
	model.TransactionUpdateAt = helpers.CreateDateTime()
	return
}

func (model *Transaction) BeforeUpdate(tx *gorm.DB) (err error) {

	model.TransactionUpdateAt = helpers.CreateDateTime()
	return
}

func (Transaction) FromEntity(e *entity.Transaction) *Transaction {
	return &Transaction{
		TransactionId:                e.TransactionId,
		TransactionUserId:            e.TransactionUserId,
		TransactionTenorId:           e.TransactionTenorId,
		TransactionPerusahaanAssetId: e.TransactionPerusahaanAssetId,
		TransactionCreateAt:          e.TransactionCreateAt,
		TransactionUpdateAt:          e.TransactionUpdateAt,
	}
}

func (model *Transaction) ToEntity() *entity.TransactionResponse {
	modelData := &entity.TransactionResponse{
		TransactionId:                model.TransactionId,
		TransactionUserId:            model.TransactionUserId,
		TransactionTenorId:           model.TransactionTenorId,
		TransactionPerusahaanAssetId: model.TransactionPerusahaanAssetId,
		TransactionCreateAt:          model.TransactionCreateAt,
		TransactionUpdateAt:          model.TransactionUpdateAt,
	}
	return modelData
}

func (model *TransactionListResponse) ToEntityList() *entity.TransactionListResponse {
	modelData := &entity.TransactionListResponse{
		TransactionID:           model.TransactionID,
		PerusahaanAssetOtrPrice: model.PerusahaanAssetOtrPrice,
		PerusahaanFee:           model.PerusahaanFee,
		Tenor:                   model.Tenor,
		TenorInterest:           model.TenorInterest,
		PerusahaanAssetNama:     model.PerusahaanAssetNama,
		TotalPrice:              model.TotalPrice,
	}
	return modelData
}

func (TransactionListResponse) FromEntityList(e *entity.TransactionListResponse) *TransactionListResponse {
	return &TransactionListResponse{
		TransactionID:           e.TransactionID,
		TransactionUserId:       e.TransactionUserId,
		PerusahaanAssetOtrPrice: e.PerusahaanAssetOtrPrice,
		PerusahaanFee:           e.PerusahaanFee,
		Tenor:                   e.Tenor,
		TenorInterest:           e.TenorInterest,
		PerusahaanAssetNama:     e.PerusahaanAssetNama,
		TotalPrice:              e.TotalPrice,
	}
}

type TransactionListResponse struct {
	TransactionID     string `json:"transaction_id"`
	TransactionUserId string `json:"transaction_user_id"`

	PerusahaanAssetOtrPrice int64  `json:"perusahaan_asset_otr_price"`
	PerusahaanFee           int64  `json:"perusahaan_fee"`
	Tenor                   int64  `json:"tenor"`
	TenorInterest           int64  `json:"tenor_interest"`
	PerusahaanAssetNama     string `json:"perusahaan_asset_nama"`
	TotalPrice              int64  `json:"total_price"`
}
