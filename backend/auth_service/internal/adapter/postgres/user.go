package postgres

import (
	"context"
	"database/sql"

	"github.com/sherinur/soar-system/backend/auth_service/internal/model"
)

type User struct {
	table string
	db    *sql.DB
}

const (
	tableUsers = "users"
)

func NewUser(tableName string) *User {
	return &User{
		table: tableUsers,
	}
}

func (u *User) Create(ctx context.Context, user model.User) error {
	return nil
}

func (u *User) Get(ctx context.Context, id int) (*model.User, error) {
	return nil, nil
}

func (u *User) GetAll(ctx context.Context, id int) ([]model.User, error) {
	return nil, nil
}

func (u *User) Update(ctx context.Context, filter model.UserFilter, update model.UserUpdateData) error {
	return nil
}

func (u *User) Delete(ctx context.Context, id int) error {
	return nil
}
