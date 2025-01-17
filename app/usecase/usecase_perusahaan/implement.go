package usecase_perusahaan

import (
	"net/http"

	"github.com/edwinjordan/SigmatechTest-Golang/app/repository"
	"github.com/edwinjordan/SigmatechTest-Golang/config"
	"github.com/edwinjordan/SigmatechTest-Golang/entity"
	"github.com/edwinjordan/SigmatechTest-Golang/handler"
	"github.com/edwinjordan/SigmatechTest-Golang/pkg/exceptions"
	"github.com/edwinjordan/SigmatechTest-Golang/pkg/helpers"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"golang.org/x/net/html"
)

type UseCaseImpl struct {
	PerusahaanRepository repository.PerusahaanRepository
	Validate             *validator.Validate
}

func NewUseCase(perusahaanRepo repository.PerusahaanRepository, validate *validator.Validate) UseCase {
	return &UseCaseImpl{
		Validate:             validate,
		PerusahaanRepository: perusahaanRepo,
	}
}

// Create implements UseCase.
func (controller *UseCaseImpl) Create(w http.ResponseWriter, r *http.Request) {
	//panic("unimplemented")
	dataRequest := entity.Perusahaan{}
	helpers.ReadFromRequestBody(r, &dataRequest)

	err := controller.Validate.Struct(dataRequest)
	helpers.PanicIfError(err)
	perusahaan := controller.PerusahaanRepository.Create(r.Context(), dataRequest)

	webResponse := map[string]interface{}{
		"code":   200,
		"status": config.LoadMessage().SuccessCreateData,
		"data":   perusahaan,
	}

	helpers.WriteToResponseBody(w, webResponse)
}

// Delete implements UseCase.
func (controller *UseCaseImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["perusahaanId"]
	_, err := controller.PerusahaanRepository.FindById(r.Context(), id)

	if err != nil {
		webResponse := map[string]interface{}{
			"code":   404,
			"status": config.LoadMessage().GetDataByIdNotFound,
		}
		helpers.WriteToResponseBody(w, webResponse)
	} else {
		controller.PerusahaanRepository.Delete(r.Context(), id)

		webResponse := map[string]interface{}{
			"code":   200,
			"status": config.LoadMessage().SuccessDeleteData,
		}
		helpers.WriteToResponseBody(w, webResponse)
	}
}

// FindAll implements UseCase.
func (controller *UseCaseImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	dataResponse := controller.PerusahaanRepository.FindAll(r.Context())

	webResponse := map[string]interface{}{
		"code":   200,
		"status": config.LoadMessage().SuccessGetData,
		"data":   dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

// FindById implements UseCase.
func (controller *UseCaseImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["perusahaanId"]
	dataResponse, err := controller.PerusahaanRepository.FindById(r.Context(), id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

// Update implements UseCase.
func (controller *UseCaseImpl) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["perusahaanId"]
	//	dataRequest := map[string]interface{}{}
	dataRequest := entity.Perusahaan{}
	helpers.ReadFromRequestBody(r, &dataRequest)

	_, err := controller.PerusahaanRepository.FindById(r.Context(), id)
	if err != nil {
		http.Error(w, "Perusahaan not found", http.StatusNotFound)
		return
	}
	if dataRequest.PerusahaanNama != "" {
		//value, _ := dataRequest["tenor"].(float64)
		dataPerusahaan := entity.Perusahaan{
			PerusahaanNama:   html.EscapeString(dataRequest.PerusahaanNama),
			PerusahaanFee:    dataRequest.PerusahaanFee,
			PerusahaanAlamat: html.EscapeString(dataRequest.PerusahaanAlamat),
		}

		dataResponse := controller.PerusahaanRepository.Update(r.Context(), []string{"perusahaan_nama", "perusahaan_fee", "perusahaan_alamat"}, dataPerusahaan, id)
		w.WriteHeader(http.StatusOK)
		webResponse := map[string]interface{}{
			"code":   200,
			"status": config.LoadMessage().SuccessUpdateData,
			"data":   dataResponse,
		}
		helpers.WriteToResponseBody(w, webResponse)
	} else {
		w.WriteHeader(http.StatusOK)
		webResponse := map[string]interface{}{
			"code":   400,
			"status": config.LoadMessage().GetDataByIdNotFound,
			//"data":   null,
		}

		helpers.WriteToResponseBody(w, webResponse)
	}
}
