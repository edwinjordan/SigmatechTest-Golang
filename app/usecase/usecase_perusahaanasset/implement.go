package usecase_perusahaanasset

import (
	"net/http"

	"github.com/edwinjordan/SigmatechTest-Golang/app/repository"
	"github.com/edwinjordan/SigmatechTest-Golang/config"
	"github.com/edwinjordan/SigmatechTest-Golang/entity"
	"github.com/edwinjordan/SigmatechTest-Golang/pkg/helpers"
	"github.com/go-playground/validator/v10"
)

type UseCaseImpl struct {
	PerusahaanAssetRepository repository.PerusahaanAssetRepository
	Validate                  *validator.Validate
}

func NewUseCase(perusahaanAssetRepo repository.PerusahaanAssetRepository, validate *validator.Validate) UseCase {
	return &UseCaseImpl{
		Validate:                  validate,
		PerusahaanAssetRepository: perusahaanAssetRepo,
	}
}

// Create implements UseCase.
func (controller *UseCaseImpl) Create(w http.ResponseWriter, r *http.Request) {
	dataRequest := entity.PerusahaanAsset{}
	helpers.ReadFromRequestBody(r, &dataRequest)

	err := controller.Validate.Struct(dataRequest)
	helpers.PanicIfError(err)
	if dataRequest.PerusahaanAssetNama != "" {
		perusahaanasset := controller.PerusahaanAssetRepository.Create(r.Context(), dataRequest)

		webResponse := map[string]interface{}{
			"code":   200,
			"status": config.LoadMessage().SuccessCreateData,
			"data":   perusahaanasset,
		}

		helpers.WriteToResponseBody(w, webResponse)
	} else {
		webResponse := map[string]interface{}{
			"code":   400,
			"status": config.LoadMessage().GetDataByIdNotFound,
		}

		helpers.WriteToResponseBody(w, webResponse)
	}
}

// FindAll implements UseCase.
func (controller *UseCaseImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	dataResponse := controller.PerusahaanAssetRepository.FindAll(r.Context())

	webResponse := map[string]interface{}{
		"code":   200,
		"status": config.LoadMessage().SuccessGetData,
		"data":   dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}
