package usecase

import (
	"context"

	"github.com/sherinur/soar-system/backend/auth_service/internal/model"
)

type UserUsecase interface {
	Create(ctx context.Context, user model.User) error
	Get(ctx context.Context, id int) (*model.User, error)
	GetAll(ctx context.Context, id int) ([]model.User, error)
	Update(ctx context.Context, filter model.UserFilter, update model.UserUpdateData) error
	Delete(ctx context.Context, id int) error
}
