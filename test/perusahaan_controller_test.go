package test

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/edwinjordan/SigmatechTest-Golang/app/usecase/usecase_perusahaan"
	"github.com/edwinjordan/SigmatechTest-Golang/entity"
	"github.com/edwinjordan/SigmatechTest-Golang/pkg/mysql"
	"github.com/edwinjordan/SigmatechTest-Golang/repository/perusahaan_repository"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setupTestDBPerusahaan() *sql.DB {
	// Implement the logic to setup and return a test database connection
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang_test_sigma")
	if err != nil {
		panic(err)
	}
	return db
}

func setupRouterPerusahaan() http.Handler {
	validate := validator.New()
	db := mysql.DBConnectGorm()
	router := mux.NewRouter()
	perusahaanRepository := perusahaan_repository.New(db)
	perusahaanController := usecase_perusahaan.NewUseCase(perusahaanRepository, validate)
	router.HandleFunc("/api/perusahaan", perusahaanController.FindAll).Methods("GET")
	router.HandleFunc("/api/perusahaan", perusahaanController.Create).Methods("POST")
	router.HandleFunc("/api/perusahaan/{perusahaanId}", perusahaanController.Update).Methods("PUT")
	router.HandleFunc("/api/perusahaan/{perusahaanId}", perusahaanController.Delete).Methods("DELETE")

	return router

}

func TestGetPerusahaan(t *testing.T) {
	db := setupTestDBPerusahaan()
	defer db.Close()
	//truncateCategory(db)
	router := setupRouterPerusahaan()

	// Insert test data
	db.Exec("INSERT INTO tb_perusahaan(perusahaan_id, perusahaan_nama,perusahaan_fee,perusahaan_alamat,perusahaan_create_at) VALUES ('4778ba0e5d7b31d39706e107d668ccf7', 'PT AAA','100000','Kediri','2025-01-19 00:00:00')")

	req, _ := http.NewRequest("GET", "/api/perusahaan", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	body, _ := io.ReadAll(rr.Body)
	var perusahaan []entity.Perusahaan
	json.Unmarshal(body, &perusahaan)

	//assert.Equal(t, 2, len(categories))
	if len(perusahaan) > 0 {
		assert.Equal(t, "PT AAA", perusahaan[0].PerusahaanNama)
	}
}

func TestCreatePerusahaan(t *testing.T) {
	db := setupTestDBPerusahaan()
	defer db.Close()
	//truncateCategory(db)
	router := setupRouterPerusahaan()

	payload := `{"perusahaan_nama": "PT PSSI","perusahaan_fee" : 200000 ,"perusahaan_alamat" : "Malang"}`
	req, _ := http.NewRequest("POST", "/api/perusahaan", strings.NewReader(payload))
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
	assert.Equal(t, "PT PSSI", responseBody["data"].(map[string]interface{})["perusahaan_nama"])
}

func TestCreatePerusahaanFailed(t *testing.T) {
	db := setupTestDBPerusahaan()
	defer db.Close()
	//	truncateCategory(db)
	router := setupRouterPerusahaan()

	requestBody := strings.NewReader(`{"perusahaan_nama" : "","perusahaan_fee" : 70000 ,"perusahaan_alamat" : "Jakarta"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3131/api/perusahaan", requestBody)
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

func TestUpdatePerusahaan(t *testing.T) {
	db := setupTestDBPerusahaan()
	defer db.Close()
	//truncateCategory(db)
	router := setupRouterPerusahaan()

	// Insert test data
	db.Exec("INSERT INTO tb_perusahaan (perusahaan_id, perusahaan_nama,perusahaan_fee,perusahaan_alamat) VALUES ('4778ba0e5d7b31d39706e107d668ccfd', 'PT AFG',75000,'Kediri')")

	payload := strings.NewReader(`{"perusahaan_nama" : "PT WBI"}`)
	req := httptest.NewRequest(http.MethodPut, "http://localhost:3131/api/perusahaan/4778ba0e5d7b31d39706e107d668ccfd", payload)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	response := rr.Result()

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Berhasil mengubah data", responseBody["status"])
	assert.Equal(t, "PT WBI", responseBody["data"].(map[string]interface{})["perusahaan_nama"])
}

func TestDeletePerusahaanSuccess(t *testing.T) {
	db := setupTestDBPerusahaan()
	//truncateCategory(db)
	defer db.Close()
	router := setupRouterPerusahaan()

	db.Exec("INSERT INTO tb_perusahaan (perusahaan_id, perusahaan_nama,perusahaan_fee,perusahaan_alamat) VALUES ('4778ba0e5d7b31d39706e107d668cccd', 'PT AFG',75000,'Kediri')")

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3131/api/perusahaan/4778ba0e5d7b31d39706e107d668cccd", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Berhasil menghapus data", responseBody["status"])
}

func TestDeletePerusahaanFailed(t *testing.T) {
	db := setupTestDBPerusahaan()
	//truncateCategory(db)
	defer db.Close()
	router := setupRouterPerusahaan()

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3131/api/perusahaan/404", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	//assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "Data tidak ditemukan", responseBody["status"])
}
