package perusahaanasset_repository

import (
	"time"

	"github.com/edwinjordan/SigmatechTest-Golang/entity"
	"github.com/edwinjordan/SigmatechTest-Golang/pkg/helpers"
	"gorm.io/gorm"
)

type PerusahaanAsset struct {
	PerusahaanAssetId                string    `gorm:"column:perusahaan_asset_id"`
	PerusahaanId                     string    `gorm:"column:perusahaan_id"`
	PerusahaanAssetNama              string    `gorm:"column:perusahaan_asset_nama"`
	PerusahaanAssetOtrPrice          int       `gorm:"column:perusahaan_asset_otr_price"`
	PerusahaanAssetStockAvailability int       `gorm:"column:perusahaan_asset_stock_availability"`
	PerusahaanAssetCreateAt          time.Time `gorm:"column:perusahaan_asset_create_at"`
	PerusahaanAssetUpdateAt          time.Time `gorm:"column:perusahaan_asset_update_at"`
}

func (PerusahaanAsset) TableName() string {
	return "tb_perusahaan_asset"
}

func (model *PerusahaanAsset) BeforeCreate(tx *gorm.DB) (err error) {
	model.PerusahaanAssetId = helpers.GenUUID()

	model.PerusahaanAssetCreateAt = helpers.CreateDateTime()
	model.PerusahaanAssetUpdateAt = helpers.CreateDateTime()
	return
}

func (model *PerusahaanAsset) BeforeUpdate(tx *gorm.DB) (err error) {

	model.PerusahaanAssetUpdateAt = helpers.CreateDateTime()
	return
}

func (PerusahaanAsset) FromEntity(e *entity.PerusahaanAsset) *PerusahaanAsset {
	return &PerusahaanAsset{
		PerusahaanAssetId:                e.PerusahaanAssetId,
		PerusahaanId:                     e.PerusahaanId,
		PerusahaanAssetNama:              e.PerusahaanAssetNama,
		PerusahaanAssetOtrPrice:          e.PerusahaanAssetOtrPrice,
		PerusahaanAssetStockAvailability: e.PerusahaanAssetStockAvailability,
		PerusahaanAssetCreateAt:          e.PerusahaanAssetCreateAt,
		PerusahaanAssetUpdateAt:          e.PerusahaanAssetUpdateAt,
	}
}

func (model *PerusahaanAsset) ToEntity() *entity.PerusahaanAssetResponse {
	modelData := &entity.PerusahaanAssetResponse{
		PerusahaanAssetId:                model.PerusahaanAssetId,
		PerusahaanId:                     model.PerusahaanId,
		PerusahaanAssetNama:              model.PerusahaanAssetNama,
		PerusahaanAssetOtrPrice:          model.PerusahaanAssetOtrPrice,
		PerusahaanAssetStockAvailability: model.PerusahaanAssetStockAvailability,
		PerusahaanAssetCreateAt:          model.PerusahaanAssetCreateAt,
		PerusahaanAssetUpdateAt:          model.PerusahaanAssetUpdateAt,
	}
	return modelData
}
