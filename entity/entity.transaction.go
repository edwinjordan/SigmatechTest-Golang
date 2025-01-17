package entity

import "time"

type Transaction struct {
	TransactionId                string    `json:"transaction_id"`
	TransactionUserId            string    `json:"transaction_user_id"`
	TransactionTenorId           string    `json:"transaction_tenor_id"`
	TransactionPerusahaanAssetId string    `json:"transaction_perusahaan_asset_id"`
	TransactionCreateAt          time.Time `json:"-"`
	TransactionUpdateAt          time.Time `json:"-"`
}

type TransactionResponse struct {
	TransactionId                string    `json:"transaction_id"`
	TransactionUserId            string    `json:"transaction_user_id"`
	TransactionTenorId           string    `json:"transaction_tenor_id"`
	TransactionPerusahaanAssetId string    `json:"transaction_perusahaan_asset_id"`
	TransactionCreateAt          time.Time `json:"-"`
	TransactionUpdateAt          time.Time `json:"-"`
}

type TransactionListResponse struct {
	TransactionID           string `json:"transaction_id"`
	TransactionUserId       string `json:"transaction_user_id"`
	PerusahaanAssetOtrPrice int64  `json:"perusahaan_asset_otr_price"`
	PerusahaanFee           int64  `json:"perusahaan_fee"`
	Tenor                   int64  `json:"tenor"`
	TenorInterest           int64  `json:"tenor_interest"`
	PerusahaanAssetNama     string `json:"perusahaan_asset_nama"`
	TotalPrice              int64  `json:"total_price"`
}
