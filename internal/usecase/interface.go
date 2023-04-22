package usecase

import (
	"context"

	"github.com/1005281342/v2/user-manager/internal/entity"
)

type User interface {
	Index(context.Context) ([]entity.User, error)
	New(context.Context)
	Create(context.Context, string, string) (*entity.User, error)
	Show(context.Context, uint) (*entity.User, error)
	Edit(context.Context, uint) (*entity.User, error)
	Update(context.Context, uint, *string, *string) error
	Delete(context.Context, uint) error
}

type UserRepo interface {
	List(context.Context) ([]entity.User, error)
	Create(context.Context, *entity.User) (*entity.User, error)
	Get(context.Context, uint) (*entity.User, error)
	Update(context.Context, uint, map[string]interface{}) error
	Delete(context.Context, uint) error
}

type UserAPI interface {
	Nil() bool
}
