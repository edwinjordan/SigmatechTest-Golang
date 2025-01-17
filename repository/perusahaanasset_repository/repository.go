package perusahaanasset_repository

import (
	"context"
	"errors"

	"github.com/edwinjordan/SigmatechTest-Golang/app/repository"
	"github.com/edwinjordan/SigmatechTest-Golang/entity"
	"github.com/edwinjordan/SigmatechTest-Golang/pkg/helpers"
	"gorm.io/gorm"
)

type PerusahaanAssetRepositoryImpl struct {
	DB *gorm.DB
}

func New(db *gorm.DB) repository.PerusahaanAssetRepository {
	return &PerusahaanAssetRepositoryImpl{
		DB: db,
	}
}

// Update implements repository.PerusahaanAssetRepository.
func (repo *PerusahaanAssetRepositoryImpl) Update(ctx context.Context, selectField interface{}, perusahaanasset entity.PerusahaanAsset, perusahaanassetId string) entity.PerusahaanAssetResponse {
	perusahaanAssetData := &PerusahaanAsset{}
	perusahaanAssetData = perusahaanAssetData.FromEntity(&perusahaanasset)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("perusahaan_asset_id = ?", perusahaanassetId).Select(selectField).Updates(&perusahaanAssetData).Error
	helpers.PanicIfError(err)

	return *perusahaanAssetData.ToEntity()
}

// FindById implements repository.PerusahaanAssetRepository.
func (repo *PerusahaanAssetRepositoryImpl) FindById(ctx context.Context, perusahaanId string) (entity.PerusahaanAssetResponse, error) {
	transData := &PerusahaanAsset{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).
		Where("perusahaan_asset_id = ?", perusahaanId).
		First(&transData).Error
	if err != nil {
		return *transData.ToEntity(), errors.New("data tidak ditemukan")
	}
	return *transData.ToEntity(), nil
}

// Create implements repository.PerusahaanAssetRepository.
func (repo *PerusahaanAssetRepositoryImpl) Create(ctx context.Context, perusahaanasset entity.PerusahaanAsset) entity.PerusahaanAssetResponse {
	perusahaanAssetData := &PerusahaanAsset{}
	perusahaanAssetData = perusahaanAssetData.FromEntity(&perusahaanasset)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).Create(&perusahaanAssetData).Error
	helpers.PanicIfError(err)

	return *perusahaanAssetData.ToEntity()
}

// FindAll implements repository.PerusahaanAssetRepository.
func (repo *PerusahaanAssetRepositoryImpl) FindAll(ctx context.Context) []entity.PerusahaanAssetResponse {
	var tempData []entity.PerusahaanAssetResponse

	perusahaanAsset := []PerusahaanAsset{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Find(&perusahaanAsset).Error
	helpers.PanicIfError(err)

	for _, v := range perusahaanAsset {
		tempData = append(tempData, *v.ToEntity())
	}

	return tempData
}

func (repo *PerusahaanAssetRepositoryImpl) FindSpesificData(ctx context.Context, where entity.PerusahaanAsset) []entity.PerusahaanAssetResponse {
	perusahaanasset := []PerusahaanAsset{}
	perusahaanassetWhere := &PerusahaanAsset{}
	perusahaanassetWhere = perusahaanassetWhere.FromEntity(&where)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where(perusahaanassetWhere).Find(&perusahaanasset).Error
	helpers.PanicIfError(err)

	var tempData []entity.PerusahaanAssetResponse
	for _, v := range perusahaanasset {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData
}
