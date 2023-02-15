package user

import (
	"context"
	"crypto/md5"
	"encoding/hex"
)

type userService struct {
	storage Storage
}

func NewUserService(repo Storage) *userService {
	return &userService{
		storage: repo,
	}
}

func (u *userService) Create(ctx context.Context, user *UserModel) error {
	passwordHash := getMD5Hash(user.Salt + user.Password)

	user.Password = passwordHash

	if err := u.storage.CreateUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func (u *userService) FindByEmail(ctx context.Context, email string) (*UserModel, error) {
	user, err := u.storage.FindUserByEmail(ctx, email)
	if err != nil {
		return &UserModel{}, err
	}

	return user, nil
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
