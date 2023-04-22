package usecase

import (
	"context"

	"github.com/1005281342/v2/user-manager/internal/entity"
	"github.com/1005281342/v2/user-manager/pkg/hash"
)

type UserUC struct {
	repo UserRepo
	api  UserAPI
}

func NewUserUC(repo UserRepo, api UserAPI) *UserUC {
	return &UserUC{repo: repo, api: api}
}

func (u *UserUC) Index(ctx context.Context) ([]entity.User, error) {
	return u.repo.List(ctx)
}

func (u *UserUC) New(ctx context.Context) {
}

func (u *UserUC) Create(ctx context.Context, email string, password string) (*entity.User, error) {
	var err error
	password, err = hash.Password(password)
	if err != nil {
		return nil, err
	}

	var user *entity.User
	if user, err = u.repo.Create(ctx, &entity.User{
		Email:    email,
		Password: password,
	}); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUC) Show(ctx context.Context, id uint) (*entity.User, error) {
	return u.repo.Get(ctx, id)
}

func (u *UserUC) Edit(ctx context.Context, id uint) (user *entity.User, err error) {
	return &entity.User{}, nil
}

func (u *UserUC) Update(ctx context.Context, id uint, email *string, password *string) error {
	var user = make(map[string]interface{})
	if email != nil {
		user["email"] = *email
	}

	if password != nil {
		var err error
		user["password"], err = hash.Password(*password)
		if err != nil {
			return err
		}
	}
	return u.repo.Update(ctx, id, user)
}

func (u *UserUC) Delete(ctx context.Context, id uint) error {
	return u.repo.Delete(ctx, id)
}
