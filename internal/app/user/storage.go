package user

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var ErrDuplicateEmail = errors.New("duplicate email")

type Storage interface {
	CreateUser(ctx context.Context, user *UserModel) error
	FindUserByEmail(ctx context.Context, email string) (*UserModel, error)
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database, collectionName string) *userRepository {
	return &userRepository{
		collection: db.Collection(collectionName),
	}
}

func (m *userRepository) CreateUser(ctx context.Context, user *UserModel) error {
	exist, err := m.FindUserByEmail(ctx, user.Email)
	switch {
	case exist != nil:
		return ErrDuplicateEmail
	case errors.Is(err, mongo.ErrNoDocuments):
		break
	default:
		return err
	}

	if _, err := m.collection.InsertOne(ctx, user); err != nil {
		return fmt.Errorf("failed to create user due error: %v", err)
	}

	return nil
}

func (m *userRepository) FindUserByEmail(ctx context.Context, email string) (u *UserModel, err error) {
	query := bson.M{"email": email}

	if err := m.collection.FindOne(ctx, query).Decode(&u); err != nil {
		return nil, err
	}

	return u, nil
}
