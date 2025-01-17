package repository

import (
	"context"

	"github.com/edwinjordan/SigmatechTest-Golang/entity"
)

type PerusahaanAssetRepository interface {
	Create(ctx context.Context, perusahaanasset entity.PerusahaanAsset) entity.PerusahaanAssetResponse
	FindAll(ctx context.Context) []entity.PerusahaanAssetResponse
	FindSpesificData(ctx context.Context, where entity.PerusahaanAsset) []entity.PerusahaanAssetResponse
	FindById(ctx context.Context, perusahaanId string) (entity.PerusahaanAssetResponse, error)
	Update(ctx context.Context, selectField interface{}, perusahaanasset entity.PerusahaanAsset, perusahaanassetId string) entity.PerusahaanAssetResponse
}
