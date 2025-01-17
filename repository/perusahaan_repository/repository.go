package perusahaan_repository

import (
	"context"
	"errors"

	"github.com/edwinjordan/SigmatechTest-Golang/app/repository"
	"github.com/edwinjordan/SigmatechTest-Golang/entity"
	"github.com/edwinjordan/SigmatechTest-Golang/pkg/helpers"
	"gorm.io/gorm"
)

type PerusahaanRepositoryImpl struct {
	DB *gorm.DB
}

func New(db *gorm.DB) repository.PerusahaanRepository {
	return &PerusahaanRepositoryImpl{
		DB: db,
	}
}

// Create implements repository.PerusahaanRepository.
func (repo *PerusahaanRepositoryImpl) Create(ctx context.Context, perusahaan entity.Perusahaan) entity.PerusahaanResponse {
	perusahaanData := &Perusahaan{}
	perusahaanData = perusahaanData.FromEntity(&perusahaan)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).Create(&perusahaanData).Error
	helpers.PanicIfError(err)

	return *perusahaanData.ToEntity()
}

// Delete implements repository.PerusahaanRepository.
func (repo *PerusahaanRepositoryImpl) Delete(ctx context.Context, perusahaanId string) {
	perusahaan := &Perusahaan{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("perusahaan_id = ?", perusahaanId).Delete(&perusahaan).Error
	helpers.PanicIfError(err)
}

// FindAll implements repository.PerusahaanRepository.
func (repo *PerusahaanRepositoryImpl) FindAll(ctx context.Context) []entity.PerusahaanResponse {
	var tempData []entity.PerusahaanResponse

	perusahaan := []Perusahaan{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Find(&perusahaan).Error
	helpers.PanicIfError(err)

	for _, v := range perusahaan {
		tempData = append(tempData, *v.ToEntity())
	}

	return tempData
}

// FindById implements repository.PerusahaanRepository.
func (repo *PerusahaanRepositoryImpl) FindById(ctx context.Context, perusahaanId string) (entity.PerusahaanResponse, error) {
	perusahaanData := &Perusahaan{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).
		Where("perusahaan_id = ?", perusahaanId).
		First(&perusahaanData).Error
	//panic(err)
	if err != nil {
		return *perusahaanData.ToEntity(), errors.New("data perusahaan tidak ditemukan")
	}
	return *perusahaanData.ToEntity(), nil
}

// Update implements repository.PerusahaanRepository.
func (repo *PerusahaanRepositoryImpl) Update(ctx context.Context, selectField interface{}, perusahaan entity.Perusahaan, perusahaanId string) entity.PerusahaanResponse {
	perusahaanData := &Perusahaan{}
	perusahaanData = perusahaanData.FromEntity(&perusahaan)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("perusahaan_id = ?", perusahaanId).Select(selectField).Updates(&perusahaanData).Error
	helpers.PanicIfError(err)

	return *perusahaanData.ToEntity()
}
