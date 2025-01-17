package user_repository

import (
	"context"

	"github.com/edwinjordan/SigmatechTest-Golang/app/repository"
	"github.com/edwinjordan/SigmatechTest-Golang/entity"
	"github.com/edwinjordan/SigmatechTest-Golang/pkg/helpers"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func New(db *gorm.DB) repository.UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

// Create implements repository.UserRepository.
func (repo *UserRepositoryImpl) Create(ctx context.Context, user entity.User) entity.UserResponse {
	userData := &User{}
	userData = userData.FromEntity(&user)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).Create(&userData).Error
	helpers.PanicIfError(err)

	return *userData.ToEntity()
}

func (repo *UserRepositoryImpl) FindSpesificData(ctx context.Context, where entity.User) []entity.UserResponse {
	user := []User{}
	userWhere := &User{}
	userWhere = userWhere.FromEntity(&where)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where(userWhere).Find(&user).Error
	helpers.PanicIfError(err)

	var tempData []entity.UserResponse
	for _, v := range user {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData
}
