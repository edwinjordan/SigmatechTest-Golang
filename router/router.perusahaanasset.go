package router

import (
	"github.com/edwinjordan/SigmatechTest-Golang/app/usecase/usecase_perusahaanasset"
	"github.com/edwinjordan/SigmatechTest-Golang/repository/perusahaanasset_repository"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func PerusahaanAssetRouter(db *gorm.DB, validate *validator.Validate, router *mux.Router) {
	perusahaanAssetRepository := perusahaanasset_repository.New(db)
	perusahaanAssetController := usecase_perusahaanasset.NewUseCase(perusahaanAssetRepository, validate)
	router.HandleFunc("/api/perusahaanasset", perusahaanAssetController.FindAll).Methods("GET")
	router.HandleFunc("/api/perusahaanasset", perusahaanAssetController.Create).Methods("POST")
}
