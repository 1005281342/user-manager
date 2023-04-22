package user

import (
	"context"

	"github.com/1005281342/v2/user-manager/internal/config"
	"github.com/1005281342/v2/user-manager/internal/entity"
	"github.com/1005281342/v2/user-manager/internal/usecase"
	"github.com/1005281342/v2/user-manager/internal/usecase/api"
	"github.com/1005281342/v2/user-manager/internal/usecase/repo"
	"github.com/1005281342/v2/user-manager/pkg/db"
)

// Controller for users
type Controller struct {
	config.Config
	uc usecase.User
}

func Load(cfg config.Config) (*Controller, error) {
	var t, err = db.New(cfg.Gorm.Driver, cfg.Gorm.Dsn)
	if err != nil {
		return nil, err
	}
	if err = t.AutoMigrate(&entity.User{}); err != nil {
		return nil, err
	}

	var uc = usecase.NewUserUC(repo.NewUserRepo(t), api.NewUserAPI())

	return &Controller{uc: uc}, nil
}

// Index of users
// GET /user
func (c *Controller) Index(ctx context.Context) (users []entity.User, err error) {
	return c.uc.Index(ctx)
}

// New returns a view for creating a new user
// GET /user/new
func (c *Controller) New(ctx context.Context) {
	c.uc.New(ctx)
}

// Create user
// POST /user
func (c *Controller) Create(ctx context.Context, email string, password string) (*entity.User, error) {
	return c.uc.Create(ctx, email, password)
}

// Show user
// GET /user/:id
func (c *Controller) Show(ctx context.Context, id uint) (user *entity.User, err error) {
	return c.uc.Show(ctx, id)
}

// Edit returns a view for editing a user
// GET /user/:id/edit
func (c *Controller) Edit(ctx context.Context, id uint) (user *entity.User, err error) {
	return c.uc.Edit(ctx, id)
}

// Update user
// PATCH /user/:id
func (c *Controller) Update(ctx context.Context, id uint, email *string, password *string) error {
	return c.uc.Update(ctx, id, email, password)
}

// Delete user
// DELETE /user/:id
func (c *Controller) Delete(ctx context.Context, id uint) error {
	return c.uc.Delete(ctx, id)
}
