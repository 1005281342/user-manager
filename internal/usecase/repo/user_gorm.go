package repo

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/1005281342/v2/user-manager/internal/entity"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) List(ctx context.Context) ([]entity.User, error) {
	var (
		users []entity.User
		err   error
	)
	if err = u.db.Find(&users).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return users, nil
}

func (u *UserRepo) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	if user == nil {
		return nil, fmt.Errorf("user is nil")
	}

	if err := u.db.Model(&entity.User{}).Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepo) Get(ctx context.Context, id uint) (*entity.User, error) {
	if id == 0 {
		return nil, fmt.Errorf("id is 0")
	}

	var user = &entity.User{}
	if err := u.db.Model(&entity.User{}).Where("id = ?", id).Find(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepo) Update(ctx context.Context, id uint, user map[string]interface{}) error {
	return u.db.Model(&entity.User{}).Where("id = ?", id).Updates(user).Error
}

func (u *UserRepo) Delete(ctx context.Context, id uint) error {
	return u.db.Model(&entity.User{}).Where("id = ?", id).Delete(&entity.User{Model: gorm.Model{ID: id}}).Error
}
