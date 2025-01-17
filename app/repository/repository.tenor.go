package repository

import (
	"context"

	"github.com/edwinjordan/SigmatechTest-Golang/entity"
)

type TenorRepository interface {
	Create(ctx context.Context, tenor entity.Tenor) entity.TenorResponse
	Update(ctx context.Context, selectField interface{}, tenor entity.Tenor, tenorId string) entity.TenorResponse
	Delete(ctx context.Context, tenorId string)
	FindById(ctx context.Context, tenorId string) (entity.TenorResponse, error)
	FindAll(ctx context.Context) []entity.TenorResponse
	FindSpesificData(ctx context.Context, where entity.Tenor) []entity.TenorResponse
}
