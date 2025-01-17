package test

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/edwinjordan/SigmatechTest-Golang/app/usecase/usecase_tenor"
	"github.com/edwinjordan/SigmatechTest-Golang/entity"
	"github.com/edwinjordan/SigmatechTest-Golang/pkg/mysql"
	"github.com/edwinjordan/SigmatechTest-Golang/repository/tenor_repository"
	"github.com/edwinjordan/SigmatechTest-Golang/repository/user_repository"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	// Implement the logic to setup and return a test database connection
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang_test_sigma")
	if err != nil {
		panic(err)
	}
	return db
}

func setupRouter() http.Handler {
	validate := validator.New()
	db := mysql.DBConnectGorm()
	router := mux.NewRouter()
	tenorRepository := tenor_repository.New(db)
	userRepository := user_repository.New(db)
	tenorController := usecase_tenor.NewUseCase(tenorRepository, userRepository, validate)
	router.HandleFunc("/api/tenor", tenorController.FindAll).Methods("GET")
	router.HandleFunc("/api/tenor", tenorController.Create).Methods("POST")
	router.HandleFunc("/api/tenor/{tenorId}", tenorController.Update).Methods("PUT")
	router.HandleFunc("/api/tenor/{tenorId}", tenorController.Delete).Methods("DELETE")

	return router

}

func TestGetTenor(t *testing.T) {
	db := setupTestDB()
	defer db.Close()
	//truncateCategory(db)
	router := setupRouter()

	// Insert test data
	db.Exec("INSERT INTO tb_tenor (tenor_id, user_id,tenor,tenor_max_limit,tenor_interest) VALUES ('4778ba0e5d7b31d39706e107d668ccf7', 'affe0a6a8e58f36bab0dcc1cb8bbc8c9','7','2000000','25')")

	req, _ := http.NewRequest("GET", "/api/tenor", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	body, _ := io.ReadAll(rr.Body)
	var tenor []entity.Tenor
	json.Unmarshal(body, &tenor)

	//assert.Equal(t, 2, len(categories))
	if len(tenor) > 0 {
		assert.Equal(t, "affe0a6a8e58f36bab0dcc1cb8bbc8c9", tenor[0].UserId)
	}
}

func TestCreateTenor(t *testing.T) {
	db := setupTestDB()
	defer db.Close()
	//truncateCategory(db)
	router := setupRouter()

	payload := `{"user_id": "affe0a6a8e58f36bab0dcc1cb8bbc8c9","tenor" : 7 ,"tenor_max_limit" : 3000000,"tenor_interest" : 25}`
	req, _ := http.NewRequest("POST", "/api/tenor", strings.NewReader(payload))
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
	assert.Equal(t, "affe0a6a8e58f36bab0dcc1cb8bbc8c9", responseBody["data"].(map[string]interface{})["user_id"])
}

func TestCreateTenorFailed(t *testing.T) {
	db := setupTestDB()
	defer db.Close()
	//	truncateCategory(db)
	router := setupRouter()

	requestBody := strings.NewReader(`{"user_id" : "","tenor" : 7 ,"tenor_max_limit" : 3000000,"tenor_interest" : 25}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3131/api/tenor", requestBody)
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

func TestUpdateTenor(t *testing.T) {
	db := setupTestDB()
	defer db.Close()
	//truncateCategory(db)
	router := setupRouter()

	// Insert test data
	db.Exec("INSERT INTO tb_tenor (tenor_id, user_id,tenor,tenor_max_limit,tenor_interest) VALUES ('4778ba0e5d7b31d39706e107d668ccf7', 'affe0a6a8e58f36bab0dcc1cb8bbc8c9','7','2000000','25')")

	payload := strings.NewReader(`{"tenor" : 8,"user_id":"affe0a6a8e58f36bab0dcc1cb8bbc8c9"}`)
	req := httptest.NewRequest(http.MethodPut, "http://localhost:3131/api/tenor/4778ba0e5d7b31d39706e107d668ccf7", payload)
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
	assert.Equal(t, "affe0a6a8e58f36bab0dcc1cb8bbc8c9", responseBody["data"].(map[string]interface{})["user_id"])
}

func TestUpdateTenorFailed(t *testing.T) {
	db := setupTestDB()
	defer db.Close()
	//truncateCategory(db)
	router := setupRouter()

	requestBody := strings.NewReader(`{"user_id" : ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3131/api/tenor/fc4f7247e20ee1d179b219286b558329", requestBody)
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

func TestDeleteTenorSuccess(t *testing.T) {
	db := setupTestDB()
	//truncateCategory(db)
	defer db.Close()
	router := setupRouter()

	db.Exec("INSERT INTO tb_tenor (tenor_id, user_id,tenor,tenor_max_limit,tenor_interest) VALUES ('4778ba0e5d7b31d39706e107d668ccf7', 'affe0a6a8e58f36bab0dcc1cb8bbc8c9','7','2000000','25')")

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3131/api/tenor/4778ba0e5d7b31d39706e107d668ccf7", nil)
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

func TestDeleteTenorFailed(t *testing.T) {
	db := setupTestDB()
	//truncateCategory(db)
	defer db.Close()
	router := setupRouter()

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3131/api/tenor/404", nil)
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
