package main

import (
	"github.com/edwinjordan/SigmatechTest-Golang/config"
	"github.com/edwinjordan/SigmatechTest-Golang/pkg/helpers"
	"github.com/edwinjordan/SigmatechTest-Golang/pkg/mysql"
	"github.com/edwinjordan/SigmatechTest-Golang/router"
	"github.com/go-playground/validator/v10"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	validate := validator.New()
	db := mysql.DBConnectGorm()
	route := mux.NewRouter()

	/** remove comment on this line if you want to run db migration **/
	// migrate.Migrate(db)

	/* setting cors */
	corsOpt := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodDelete,
			http.MethodPut,
			http.MethodPatch,
		},
		AllowedHeaders: []string{
			"*",
		},
	})

	router.UserRouter(db, validate, route)
	router.TenorRouter(db, validate, route)
	router.PerusahaanRouter(db, validate, route)
	router.PerusahaanAssetRouter(db, validate, route)
	router.TransactionRouter(db, validate, route)

	server := http.Server{
		Addr:    config.GetEnv("HOST_ADDR"),
		Handler: corsOpt.Handler(route),
	}

	err := server.ListenAndServe()
	helpers.PanicIfError(err)
}
