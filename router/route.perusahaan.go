package router

import (
	"github.com/edwinjordan/SigmatechTest-Golang/app/usecase/usecase_perusahaan"
	"github.com/edwinjordan/SigmatechTest-Golang/repository/perusahaan_repository"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func PerusahaanRouter(db *gorm.DB, validate *validator.Validate, router *mux.Router) {
	perusahaanRepository := perusahaan_repository.New(db)
	perusahaanController := usecase_perusahaan.NewUseCase(perusahaanRepository, validate)
	router.HandleFunc("/api/perusahaan", perusahaanController.FindAll).Methods("GET")
	router.HandleFunc("/api/perusahaan/{perusahaanId}", perusahaanController.FindById).Methods("GET")
	router.HandleFunc("/api/perusahaan", perusahaanController.Create).Methods("POST")
	router.HandleFunc("/api/perusahaan/{perusahaanId}", perusahaanController.Update).Methods("PUT")
	router.HandleFunc("/api/perusahaan/{perusahaanId}", perusahaanController.Delete).Methods("DELETE")
}
