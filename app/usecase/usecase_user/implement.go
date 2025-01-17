package usecase_user

import (
	"net/http"
	"path/filepath"

	"github.com/edwinjordan/SigmatechTest-Golang/app/repository"
	"github.com/edwinjordan/SigmatechTest-Golang/entity"
	"github.com/edwinjordan/SigmatechTest-Golang/handler"
	"github.com/edwinjordan/SigmatechTest-Golang/pkg/exceptions"
	"github.com/edwinjordan/SigmatechTest-Golang/pkg/helpers"
	"github.com/go-playground/validator/v10"
)

type UseCaseImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUseCase(userRepo repository.UserRepository, validate *validator.Validate) UseCase {
	return &UseCaseImpl{
		Validate:       validate,
		UserRepository: userRepo,
	}
}

// Create implements UseCase.
func (controller *UseCaseImpl) Create(w http.ResponseWriter, r *http.Request) {
	dataRequest := entity.User{}

	helpers.ReadFromRequestBody(r, &dataRequest)

	/* validasi */
	err := controller.Validate.Struct(dataRequest)
	helpers.PanicIfError(err)
	dataRequest.FotoKtpPath = "uploaded_files/"
	dataRequest.FotoSelfiePath = "uploaded_files/"

	/* ambil data file jika ada saat mengirim file menggunakan multipart/form-data */
	filePhoto := dataRequest.FotoKtpData
	if filePhoto != nil {
		/* cek format file */
		ext := filepath.Ext(filePhoto["FileName"].(string))
		if ok, err := helpers.FileUploadFormat(ext[1:], "png|jpeg|jpg"); !ok {
			panic(exceptions.NewBadRequestError(err.Error()))
		}

		dataRequest.FotoKtp = helpers.GenUUID() + "." + ext[1:]
		helpers.SaveFileFromBase64(dataRequest.FotoKtp, filePhoto["Base64"].(string), "./"+dataRequest.FotoKtpPath)
	}

	filePhotoSelfie := dataRequest.FotoSelfieData
	if filePhotoSelfie != nil {
		/* cek format file */
		ext := filepath.Ext(filePhotoSelfie["FileName"].(string))
		if ok, err := helpers.FileUploadFormat(ext[1:], "png|jpeg|jpg"); !ok {
			panic(exceptions.NewBadRequestError(err.Error()))
		}

		dataRequest.FotoSelfie = helpers.GenUUID() + "." + ext[1:]
		helpers.SaveFileFromBase64(dataRequest.FotoSelfie, filePhotoSelfie["Base64"].(string), "./"+dataRequest.FotoSelfiePath)
	}
	/* tambah ke table */
	response := controller.UserRepository.Create(r.Context(), dataRequest)

	dataResponse := &entity.UserResponse{
		UserId:   response.UserId,
		NIK:      response.NIK,
		FullName: response.FullName,
	}

	webResponse := handler.WebResponse{
		Error:   false,
		Message: "Berhasil Menambah Data",
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}
