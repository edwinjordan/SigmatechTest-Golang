package repository

import (
	"context"

	"github.com/edwinjordan/SigmatechTest-Golang/entity"
)

type TransactionRepository interface {
	Create(ctx context.Context, transaction entity.Transaction) entity.TransactionResponse
	FindAll(ctx context.Context, where map[string]interface{}) []map[string]interface{}
}
