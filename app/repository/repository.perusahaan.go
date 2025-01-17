package repository

import (
	"context"

	"github.com/edwinjordan/SigmatechTest-Golang/entity"
)

type PerusahaanRepository interface {
	Create(ctx context.Context, perusahaan entity.Perusahaan) entity.PerusahaanResponse
	Update(ctx context.Context, selectField interface{}, perusahaan entity.Perusahaan, perusahaanId string) entity.PerusahaanResponse
	Delete(ctx context.Context, perusahaanId string)
	FindById(ctx context.Context, perusahaanId string) (entity.PerusahaanResponse, error)
	FindAll(ctx context.Context) []entity.PerusahaanResponse
}
