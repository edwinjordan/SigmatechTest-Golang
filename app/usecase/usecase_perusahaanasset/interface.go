package usecase_perusahaanasset

import "net/http"

type UseCase interface {
	Create(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
}
