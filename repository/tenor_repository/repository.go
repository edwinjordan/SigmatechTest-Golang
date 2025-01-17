package tenor_repository

import (
	"context"
	"errors"

	"github.com/edwinjordan/SigmatechTest-Golang/app/repository"
	"github.com/edwinjordan/SigmatechTest-Golang/entity"
	"github.com/edwinjordan/SigmatechTest-Golang/pkg/helpers"
	"gorm.io/gorm"
)

type TenorRepositoryImpl struct {
	DB *gorm.DB
}

func New(db *gorm.DB) repository.TenorRepository {
	return &TenorRepositoryImpl{
		DB: db,
	}
}

// Create implements repository.TenorRepository.
func (repo *TenorRepositoryImpl) Create(ctx context.Context, tenor entity.Tenor) entity.TenorResponse {
	tenorData := &Tenor{}
	tenorData = tenorData.FromEntity(&tenor)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).Create(&tenorData).Error
	helpers.PanicIfError(err)

	return *tenorData.ToEntity()

}

// Delete implements repository.TenorRepository.
func (repo *TenorRepositoryImpl) Delete(ctx context.Context, tenorId string) {
	tenor := &Tenor{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("tenor_id = ?", tenorId).Delete(&tenor).Error
	helpers.PanicIfError(err)
}

// FindAll implements repository.TenorRepository.
func (repo *TenorRepositoryImpl) FindAll(ctx context.Context) []entity.TenorResponse {
	var tempData []entity.TenorResponse

	tenor := []Tenor{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Find(&tenor).Error
	helpers.PanicIfError(err)

	for _, v := range tenor {
		tempData = append(tempData, *v.ToEntity())
	}

	return tempData
}

// FindById implements repository.TenorRepository.
func (repo *TenorRepositoryImpl) FindById(ctx context.Context, tenorId string) (entity.TenorResponse, error) {
	tenorData := &Tenor{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).
		Where("tenor_id = ?", tenorId).
		First(&tenorData).Error
	//panic(err)
	if err != nil {
		return *tenorData.ToEntity(), errors.New("data tenor tidak ditemukan")
	}
	return *tenorData.ToEntity(), nil
}

// Update implements repository.TenorRepository.
func (repo *TenorRepositoryImpl) Update(ctx context.Context, selectField interface{}, tenor entity.Tenor, tenorId string) entity.TenorResponse {
	tenorData := &Tenor{}
	tenorData = tenorData.FromEntity(&tenor)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("tenor_id = ?", tenorId).Select(selectField).Updates(&tenorData).Error
	helpers.PanicIfError(err)

	return *tenorData.ToEntity()

}

func (repo *TenorRepositoryImpl) FindSpesificData(ctx context.Context, where entity.Tenor) []entity.TenorResponse {
	tenor := []Tenor{}
	tenorWhere := &Tenor{}
	tenorWhere = tenorWhere.FromEntity(&where)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where(tenorWhere).Find(&tenor).Error
	helpers.PanicIfError(err)

	var tempData []entity.TenorResponse
	for _, v := range tenor {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData
}
