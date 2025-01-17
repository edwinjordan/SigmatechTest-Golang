package repository

import (
	"context"

	"github.com/edwinjordan/SigmatechTest-Golang/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user entity.User) entity.UserResponse
	FindSpesificData(ctx context.Context, where entity.User) []entity.UserResponse
}
