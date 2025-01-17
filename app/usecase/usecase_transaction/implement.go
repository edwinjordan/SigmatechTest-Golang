package usecase_transaction

import (
	"net/http"
	"sync"

	"github.com/edwinjordan/SigmatechTest-Golang/app/repository"
	"github.com/edwinjordan/SigmatechTest-Golang/config"
	"github.com/edwinjordan/SigmatechTest-Golang/entity"
	"github.com/edwinjordan/SigmatechTest-Golang/pkg/exceptions"
	"github.com/edwinjordan/SigmatechTest-Golang/pkg/helpers"
	"github.com/go-playground/validator/v10"
)

type UseCaseImpl struct {
	TransactionRepository     repository.TransactionRepository
	UserRepository            repository.UserRepository
	TenorRepository           repository.TenorRepository
	PerusahaanAssetRepository repository.PerusahaanAssetRepository
	Validate                  *validator.Validate
}

func NewUseCase(transactionRepo repository.TransactionRepository, userRepo repository.UserRepository, tenorRepo repository.TenorRepository, PerusahaanAssetRepo repository.PerusahaanAssetRepository, validate *validator.Validate) UseCase {
	return &UseCaseImpl{
		TransactionRepository:     transactionRepo,
		UserRepository:            userRepo,
		TenorRepository:           tenorRepo,
		PerusahaanAssetRepository: PerusahaanAssetRepo,
		Validate:                  validate,
	}
}

// Create implements UseCase.
func (controller *UseCaseImpl) Create(w http.ResponseWriter, r *http.Request) {
	dataRequest := entity.Transaction{}
	dataRequestAsset := entity.PerusahaanAsset{}
	helpers.ReadFromRequestBody(r, &dataRequest)

	err := controller.Validate.Struct(dataRequest)
	helpers.PanicIfError(err)

	var wg sync.WaitGroup

	/* check if user exist */
	wg.Add(1)
	go func() {
		defer wg.Done()
		dataUser := controller.UserRepository.FindSpesificData(r.Context(), entity.User{
			UserId: dataRequest.TransactionUserId,
		})

		if dataUser == nil {
			panic(exceptions.NewConflictError("User Tidak Ditemukan"))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		/* check if tenor exist */
		dataTenor := controller.TenorRepository.FindSpesificData(r.Context(), entity.Tenor{
			TenorId: dataRequest.TransactionTenorId,
			UserId:  dataRequest.TransactionUserId,
		})

		if dataTenor == nil {
			panic(exceptions.NewConflictError("User Tidak Dapat Mengakses Tenor Tersebut"))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		dataPerusahaanAsset, err := controller.PerusahaanAssetRepository.FindById(r.Context(), dataRequest.TransactionPerusahaanAssetId)

		if err != nil {
			panic(exceptions.NewConflictError("Perusahaan Asset Tidak Ditemukan"))
		} else if dataPerusahaanAsset.PerusahaanAssetStockAvailability == 0 {
			panic(exceptions.NewConflictError("Perusahaan Asset Tidak Tersedia Untuk Saat Ini"))
		}
		dataRequestAsset.PerusahaanAssetStockAvailability = dataPerusahaanAsset.PerusahaanAssetStockAvailability - 1

		controller.PerusahaanAssetRepository.Update(r.Context(), []string{"perusahaan_asset_stock_availability"}, dataRequestAsset, dataRequest.TransactionPerusahaanAssetId)

		transaction := controller.TransactionRepository.Create(r.Context(), dataRequest)

		webResponse := map[string]interface{}{
			"code":   200,
			"status": config.LoadMessage().SuccessCreateData,
			"data":   transaction,
		}

		helpers.WriteToResponseBody(w, webResponse)

	}()

	// Wait for all goroutines to finish
	wg.Wait()

	/* reduce asset availability */

}
