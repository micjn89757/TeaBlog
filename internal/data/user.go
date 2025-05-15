package data

import (
	"context"

	"github.com/micjn89757/TeaBlog/internal/biz"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type userRepo struct {
	data *Data
	log  *zap.Logger
}

func NewUserRepo(data *Data, logger *zap.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  logger,
	}
}

func (ur *userRepo) CreateUser(ctx context.Context, user *biz.User) (string, error) {
	res := ur.data.db.Create(user)

	if err := res.Error; err != nil {
		return "", err
	}

	return user.ID, nil
}

func (ur *userRepo) CreateUserInBatch(ctx context.Context, users []*biz.User) error {
	if err := ur.data.db.Create(&users).Error; err != nil {
		return err
	}

	return nil
}

func (ur *userRepo) UpdateUser(ctx context.Context, user *biz.User) error {
	if err := ur.data.db.Model(user).Omit("id", "created_at", "deleted_at").Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepo) DeleteUser(ctx context.Context, id string) error {
	if err := ur.data.db.Delete(&biz.User{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (ur *userRepo) ListUser(ctx context.Context, count uint) ([]*biz.User, error) {
	var users []*biz.User
	if err := ur.data.db.Find(&users).Limit(int(count)).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			ur.log.Error("get users of", zap.Uint("count", count), zap.Error(err))
		}
		return nil, err
	}

	return nil, nil
}

func (ur *userRepo) GetUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	var u biz.User
	if err := ur.data.db.Select("id", "name", "password").Where(user).First(&u).Error; err != nil {
		if err != gorm.ErrRecordNotFound { // if user not find, don't need error log
			ur.log.Error("get User info failed", zap.Error(err)) // system panic do error log
		}
		return nil, err
	}

	return &u, nil
}
