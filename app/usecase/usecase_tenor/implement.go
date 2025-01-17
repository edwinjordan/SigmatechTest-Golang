package usecase_tenor

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
	TenorRepository repository.TenorRepository
	UserRepository  repository.UserRepository
	Validate        *validator.Validate
}

func NewUseCase(tenorRepo repository.TenorRepository, userRepo repository.UserRepository, validate *validator.Validate) UseCase {
	return &UseCaseImpl{
		Validate:        validate,
		TenorRepository: tenorRepo,
		UserRepository:  userRepo,
	}
}

// Create implements UseCase.
func (controller *UseCaseImpl) Create(w http.ResponseWriter, r *http.Request) {
	dataRequest := entity.Tenor{}
	helpers.ReadFromRequestBody(r, &dataRequest)

	err := controller.Validate.Struct(dataRequest)
	helpers.PanicIfError(err)

	/* check if user exist */
	dataUser := controller.UserRepository.FindSpesificData(r.Context(), entity.User{
		UserId: dataRequest.UserId,
	})

	if dataUser == nil {
		panic(exceptions.NewConflictError("User Tidak Ditemukan"))
	}

	if dataRequest.UserId != "" {
		dataRequest.UserId = html.EscapeString(dataRequest.UserId)
		tenor := controller.TenorRepository.Create(r.Context(), dataRequest)

		webResponse := map[string]interface{}{
			"code":   200,
			"status": config.LoadMessage().SuccessCreateData,
			"data":   tenor,
		}

		helpers.WriteToResponseBody(w, webResponse)
	} else {
		//w.WriteHeader(http.StatusBadRequest)
		webResponse := map[string]interface{}{
			"code":   400,
			"status": config.LoadMessage().GetDataByIdNotFound,
		}

		helpers.WriteToResponseBody(w, webResponse)
	}
}

// Delete implements UseCase.
func (controller *UseCaseImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["tenorId"]
	_, err := controller.TenorRepository.FindById(r.Context(), id)

	if err != nil {
		webResponse := map[string]interface{}{
			"code":   404,
			"status": config.LoadMessage().GetDataByIdNotFound,
		}
		helpers.WriteToResponseBody(w, webResponse)
	} else {
		controller.TenorRepository.Delete(r.Context(), id)

		webResponse := map[string]interface{}{
			"code":   200,
			"status": config.LoadMessage().SuccessDeleteData,
		}
		helpers.WriteToResponseBody(w, webResponse)
	}
}

// FindAll implements UseCase.
func (controller *UseCaseImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	dataResponse := controller.TenorRepository.FindAll(r.Context())

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
	id := vars["tenorId"]
	dataResponse, err := controller.TenorRepository.FindById(r.Context(), id)
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
	id := vars["tenorId"]
	//	dataRequest := map[string]interface{}{}
	dataRequest := entity.Tenor{}
	helpers.ReadFromRequestBody(r, &dataRequest)

	_, err := controller.TenorRepository.FindById(r.Context(), id)
	if err != nil {
		http.Error(w, "Tenor not found", http.StatusNotFound)
		return
	}
	if dataRequest.UserId != "" {
		//value, _ := dataRequest["tenor"].(float64)
		dataTenor := entity.Tenor{
			UserId:        html.EscapeString(dataRequest.UserId),
			Tenor:         dataRequest.Tenor,
			TenorMaxLimit: dataRequest.TenorMaxLimit,
			TenorInterest: dataRequest.TenorInterest,
		}

		dataResponse := controller.TenorRepository.Update(r.Context(), []string{"user_id", "tenor", "tenor_max_limit", "tenor_interest"}, dataTenor, id)
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
