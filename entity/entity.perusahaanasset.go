package entity

import "time"

type PerusahaanAsset struct {
	PerusahaanAssetId                string    `json:"perusahaan_asset_id"`
	PerusahaanId                     string    `json:"perusahaan_id"`
	PerusahaanAssetNama              string    `json:"perusahaan_asset_nama"`
	PerusahaanAssetOtrPrice          int       `json:"perusahaan_asset_otr_price"`
	PerusahaanAssetStockAvailability int       `json:"perusahaan_asset_stock_availability"`
	PerusahaanAssetCreateAt          time.Time `json:"-"`
	PerusahaanAssetUpdateAt          time.Time `json:"-"`
}

type PerusahaanAssetResponse struct {
	PerusahaanAssetId                string    `json:"perusahaan_asset_id"`
	PerusahaanId                     string    `json:"perusahaan_id"`
	PerusahaanAssetNama              string    `json:"perusahaan_asset_nama"`
	PerusahaanAssetOtrPrice          int       `json:"perusahaan_asset_otr_price"`
	PerusahaanAssetStockAvailability int       `json:"perusahaan_asset_stock_availability"`
	PerusahaanAssetCreateAt          time.Time `json:"-"`
	PerusahaanAssetUpdateAt          time.Time `json:"-"`
}
