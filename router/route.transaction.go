package router

import (
	"github.com/edwinjordan/SigmatechTest-Golang/app/usecase/usecase_transaction"
	"github.com/edwinjordan/SigmatechTest-Golang/repository/perusahaanasset_repository"
	"github.com/edwinjordan/SigmatechTest-Golang/repository/tenor_repository"
	"github.com/edwinjordan/SigmatechTest-Golang/repository/transaction_repository"
	"github.com/edwinjordan/SigmatechTest-Golang/repository/user_repository"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func TransactionRouter(db *gorm.DB, validate *validator.Validate, router *mux.Router) {
	transactionRepository := transaction_repository.New(db)
	userRepository := user_repository.New(db)
	tenorRepository := tenor_repository.New(db)
	perusahaanAssetRepository := perusahaanasset_repository.New(db)
	transactionController := usecase_transaction.NewUseCase(transactionRepository, userRepository, tenorRepository, perusahaanAssetRepository, validate)
	router.HandleFunc("/api/transaction", transactionController.Create).Methods("POST")
}
