package router

import (
	"github.com/edwinjordan/SigmatechTest-Golang/app/usecase/usecase_tenor"
	"github.com/edwinjordan/SigmatechTest-Golang/repository/tenor_repository"
	"github.com/edwinjordan/SigmatechTest-Golang/repository/user_repository"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func TenorRouter(db *gorm.DB, validate *validator.Validate, router *mux.Router) {
	tenorRepository := tenor_repository.New(db)
	userRepository := user_repository.New(db)
	tenorController := usecase_tenor.NewUseCase(tenorRepository, userRepository, validate)
	router.HandleFunc("/api/tenor", tenorController.FindAll).Methods("GET")
	router.HandleFunc("/api/tenor/{tenorId}", tenorController.FindById).Methods("GET")
	router.HandleFunc("/api/tenor", tenorController.Create).Methods("POST")
	router.HandleFunc("/api/tenor/{tenorId}", tenorController.Update).Methods("PUT")
	router.HandleFunc("/api/tenor/{tenorId}", tenorController.Delete).Methods("DELETE")
}
