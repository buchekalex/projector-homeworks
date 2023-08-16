package repository

import (
	"context"
	"errors"
	"go-mongo-crud-rest-api/internal/entity"
	"go-mongo-crud-rest-api/internal/usecase"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) usecase.Repository {
	return &repository{db: db}
}

func (r repository) GetUser(ctx context.Context, email string) (*entity.User, error) {
	var out user
	err := r.db.
		Collection("users").
		FindOne(ctx, bson.M{"email": email}).
		Decode(&out)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return toModel(out), nil
}

func (r repository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	out, err := r.db.
		Collection("users").
		InsertOne(ctx, fromModel(user))
	if err != nil {
		return nil, err
	}
	user.ID = out.InsertedID.(primitive.ObjectID).String()
	return user, nil
}

func (r repository) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	in := bson.M{}
	if user.Name != "" {
		in["name"] = user.Name
	}
	if user.Password != "" {
		in["password"] = user.Password
	}
	out, err := r.db.
		Collection("users").
		UpdateOne(ctx, bson.M{"email": user.Email}, bson.M{"$set": in})
	if err != nil {
		return nil, err
	}
	if out.MatchedCount == 0 {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (r repository) DeleteUser(ctx context.Context, email string) error {
	out, err := r.db.
		Collection("users").
		DeleteOne(ctx, bson.M{"email": email})
	if err != nil {
		return err
	}
	if out.DeletedCount == 0 {
		return ErrUserNotFound
	}
	return nil
}

type user struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
}

func fromModel(in *entity.User) user {
	return user{
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	}
}

func toModel(in user) *entity.User {
	return &entity.User{
		ID:       in.ID.String(),
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	}
}
