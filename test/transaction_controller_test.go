package test

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/edwinjordan/SigmatechTest-Golang/app/usecase/usecase_transaction"
	"github.com/edwinjordan/SigmatechTest-Golang/pkg/mysql"
	"github.com/edwinjordan/SigmatechTest-Golang/repository/perusahaanasset_repository"
	"github.com/edwinjordan/SigmatechTest-Golang/repository/tenor_repository"
	"github.com/edwinjordan/SigmatechTest-Golang/repository/transaction_repository"
	"github.com/edwinjordan/SigmatechTest-Golang/repository/user_repository"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setupTestDBTransaction() *sql.DB {
	// Implement the logic to setup and return a test database connection
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang_test_sigma")
	if err != nil {
		panic(err)
	}
	return db
}

func setupRouterTransaction() http.Handler {
	validate := validator.New()
	db := mysql.DBConnectGorm()
	router := mux.NewRouter()
	transactionRepository := transaction_repository.New(db)
	userRepository := user_repository.New(db)
	tenorRepository := tenor_repository.New(db)
	perusahaanAssetRepository := perusahaanasset_repository.New(db)
	transactionController := usecase_transaction.NewUseCase(transactionRepository, userRepository, tenorRepository, perusahaanAssetRepository, validate)
	router.HandleFunc("/api/transaction", transactionController.Create).Methods("POST")

	return router

}

func TestCreateTransaction(t *testing.T) {
	db := setupTestDBTransaction()
	defer db.Close()
	//truncateCategory(db)
	router := setupRouterTransaction()

	payload := `{"transaction_user_id": "e9f77b213b7708fc51b095eedaae467c","transaction_tenor_id" : "ddaaeb7b6bae734173bb942b69836e7f" ,"transaction_perusahaan_asset_id" : "2e9c3e91ae9703703ab4956db442c7cf"}`
	req, _ := http.NewRequest("POST", "/api/transaction", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	response := rr.Result()

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Berhasil menambah data", responseBody["status"])
	assert.Equal(t, "e9f77b213b7708fc51b095eedaae467c", responseBody["data"].(map[string]interface{})["transaction_user_id"])
}
