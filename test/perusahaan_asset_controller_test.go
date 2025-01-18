package test

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/edwinjordan/SigmatechTest-Golang/app/usecase/usecase_perusahaanasset"
	"github.com/edwinjordan/SigmatechTest-Golang/pkg/mysql"
	"github.com/edwinjordan/SigmatechTest-Golang/repository/perusahaanasset_repository"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setupTestDBPerusahaanAsset() *sql.DB {
	// Implement the logic to setup and return a test database connection
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang_test_sigma")
	if err != nil {
		panic(err)
	}
	return db
}

func setupRouterPerusahaanAsset() http.Handler {
	validate := validator.New()
	db := mysql.DBConnectGorm()
	router := mux.NewRouter()
	perusahaanAssetRepository := perusahaanasset_repository.New(db)
	perusahaanAssetController := usecase_perusahaanasset.NewUseCase(perusahaanAssetRepository, validate)
	router.HandleFunc("/api/perusahaanasset", perusahaanAssetController.Create).Methods("POST")

	return router

}

func TestCreatePerusahaanAsset(t *testing.T) {
	db := setupTestDBPerusahaanAsset()
	defer db.Close()
	//truncateCategory(db)
	router := setupRouterPerusahaanAsset()

	payload := `{"perusahaan_id": "4778ba0e5d7b31d39706e107d668ccf7","perusahaan_asset_nama" : "Montor" ,"perusahaan_asset_otr_price" : 50000,"perusahaan_asset_stock_availability" : 5}`
	req, _ := http.NewRequest("POST", "/api/perusahaanasset", strings.NewReader(payload))
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
	assert.Equal(t, "Montor", responseBody["data"].(map[string]interface{})["perusahaan_asset_nama"])
}

func TestCreatePerusahaanAssetFailed(t *testing.T) {
	db := setupTestDBPerusahaanAsset()
	defer db.Close()
	//	truncateCategory(db)
	router := setupRouterPerusahaanAsset()

	requestBody := strings.NewReader(`{"perusahaan_id": "4778ba0e5d7b31d39706e107d668ccf7","perusahaan_asset_nama" : "" ,"perusahaan_asset_otr_price" : 50000,"perusahaan_asset_stock_availability" : 5}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3131/api/perusahaanasset", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	//assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Data tidak ditemukan", responseBody["status"])
}
