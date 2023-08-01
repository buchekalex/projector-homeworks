package repository

import (
	"context"
	"go-mongo-crud-rest-api/internal/model"
)

type Repository interface {
	GetUser(ctx context.Context, email string) (model.User, error)
	CreateUser(ctx context.Context, in model.User) (model.User, error)
	UpdateUser(ctx context.Context, in model.User) (model.User, error)
	DeleteUser(ctx context.Context, email string) error
}
