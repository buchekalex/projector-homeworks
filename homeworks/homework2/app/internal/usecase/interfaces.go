package usecase

import (
	"context"
	"go-mongo-crud-rest-api/internal/entity"
)

type SearchIndex interface {
	GetUser(ctx context.Context, email string) (*entity.User, error)
	CreateUser(ctx context.Context, in *entity.User) (*entity.User, error)
	UpdateUser(ctx context.Context, in *entity.User) (*entity.User, error)
	DeleteUser(ctx context.Context, email string) error
	FindUser(ctx context.Context, email string) (*entity.User, error)
}

type Repository interface {
	GetUser(ctx context.Context, email string) (*entity.User, error)
	CreateUser(ctx context.Context, in *entity.User) (*entity.User, error)
	UpdateUser(ctx context.Context, in *entity.User) (*entity.User, error)
	DeleteUser(ctx context.Context, email string) error
}

type User interface {
	GetUser(ctx context.Context, email string) (*entity.User, error)
	CreateUser(ctx context.Context, in *entity.User) (*entity.User, error)
	UpdateUser(ctx context.Context, in *entity.User) (*entity.User, error)
	DeleteUser(ctx context.Context, email string) error
	FindUser(ctx context.Context, email string) (*entity.User, error)
}
