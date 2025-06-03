package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sherinur/soar-system/backend/auth_service/internal/adapter/postgres/dao"
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
	query := fmt.Sprintf(`INSERT INTO %s 
	(email,  password_hash, first_name,
	 second_name, status, 
	 organization_id, created_at, updated_at, 
	 last_login_date, password_exp_date) 
	 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`, u.table)

	userDao := dao.FromUser(user)
	_, err := u.db.ExecContext(ctx, query,
		userDao.Email,
		userDao.PasswordHash,
		userDao.FirstName,
		userDao.SecondName,
		// userDao.RoleIDs,
		// userDao.GroupIDs,
		userDao.Status,
		userDao.OrganizationID,
		userDao.CreatedAt,
		userDao.UpdatedAt,
		userDao.LastLoginDate,
		userDao.PasswordExpDate,
	)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) Get(ctx context.Context, id int) (*model.User, error) {
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, u.table)
	userDao := dao.User{}

	row := u.db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&userDao.ID,
		&userDao.Email,
		&userDao.PasswordHash,
		&userDao.FirstName,
		&userDao.SecondName,
		// &userDao.RoleIDs,
		// &userDao.GroupIDs,
		&userDao.Status,
		&userDao.OrganizationID,
		&userDao.CreatedAt,
		&userDao.UpdatedAt,
		&userDao.LastLoginDate,
		&userDao.PasswordExpDate,
	)
	if err != nil {
		return nil, err
	}

	user := dao.ToUser(&userDao)

	return &user, nil
}

func (u *User) GetAll(ctx context.Context) ([]model.User, error) {
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, u.table)

	var users []model.User
	rows, err := u.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		userDao := dao.User{}
		err := rows.Scan(
			&userDao.ID,
			&userDao.Email,
			&userDao.PasswordHash,
			&userDao.FirstName,
			&userDao.SecondName,
			// &userDao.RoleIDs,
			// &userDao.GroupIDs,
			&userDao.Status,
			&userDao.OrganizationID,
			&userDao.CreatedAt,
			&userDao.UpdatedAt,
			&userDao.LastLoginDate,
			&userDao.PasswordExpDate,
		)
		if err != nil {
			return nil, err
		}

		user := dao.ToUser(&userDao)
		users = append(users, user)
	}

	return users, nil
}

func (u *User) Update(ctx context.Context, filter model.UserFilter, update model.UserUpdateData) error {
	return nil
}

func (u *User) Delete(ctx context.Context, id int) error {
	return nil
}
