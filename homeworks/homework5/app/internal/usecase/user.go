package usecase

import (
	"context"
	"go-mongo-crud-rest-api/internal/entity"
)

type UserUseCase struct {
	repository  Repository
	searchIndex SearchIndex
}

func (u UserUseCase) FindUser(ctx context.Context, email string) (*entity.User, error) {
	user, err := u.searchIndex.FindUser(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserUseCase) GetUser(ctx context.Context, email string) (*entity.User, error) {
	user, err := u.repository.GetUser(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserUseCase) CreateUser(ctx context.Context, in *entity.User) (*entity.User, error) {
	user, err := u.repository.CreateUser(ctx, in)
	if err != nil {
		return nil, err
	}

	if _, err := u.searchIndex.CreateUser(ctx, in); err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserUseCase) UpdateUser(ctx context.Context, in *entity.User) (*entity.User, error) {
	user, err := u.repository.UpdateUser(ctx, in)
	if err != nil {
		return nil, err
	}

	if _, err := u.searchIndex.UpdateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserUseCase) DeleteUser(ctx context.Context, email string) error {
	if err := u.repository.DeleteUser(ctx, email); err != nil {
		return err
	}

	if err := u.searchIndex.DeleteUser(ctx, email); err != nil {
		return err
	}

	return nil
}

func NewUserUseCase(repository Repository,
	elasticRepository SearchIndex) *UserUseCase {
	return &UserUseCase{
		repository:  repository,
		searchIndex: elasticRepository,
	}
}
