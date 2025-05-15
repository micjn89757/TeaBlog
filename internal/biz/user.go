package biz

import (
	"context"

	"go.uber.org/zap"
)

type User struct {
	Base
	Username string `gorm:"column:username;not null"`
	Password string `gorm:"column:password;not null"`
	Role     string `gorm:"column:role;not null"`
}

// define table name
func (User) TableName() string {
	return "dev.user"
}

type UserRepo interface {
	// db
	CreateUser(ctx context.Context, user *User) (string, error)
	CreateUserInBatch(ctx context.Context, users []*User) error
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, id string) error
	ListUser(ctx context.Context, count uint) ([]*User, error)

	// redis
	// GetToken(ctx context.Context)
}

type UserUsecase struct {
	repo UserRepo
}

func NewUserUsecase(repo UserRepo, logger zap.Logger) *UserUsecase {
	return &UserUsecase{repo: repo}
}

// ChangePassword
func (uc *UserUsecase) ChangePassword(id string, data *User) error {
	return nil
}

func (uc *UserUsecase) Check(ctx context.Context) error {
	return nil
}

func (uc *UserUsecase) List(ctx context.Context, count uint) error {
	return nil
}

// password need crpto
func (uc *UserUsecase) Create(ctx context.Context, user *User) error {
	return nil
}

func (uc *UserUsecase) Delete(ctx context.Context, id string) error {
	return nil
}
