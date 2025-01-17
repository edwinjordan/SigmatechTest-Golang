package router

import (
	"github.com/edwinjordan/SigmatechTest-Golang/app/usecase/usecase_user"
	"github.com/edwinjordan/SigmatechTest-Golang/repository/user_repository"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func UserRouter(db *gorm.DB, validate *validator.Validate, router *mux.Router) {
	userRepository := user_repository.New(db)
	userController := usecase_user.NewUseCase(userRepository, validate)
	router.HandleFunc("/api/user", userController.Create).Methods("POST")

}
