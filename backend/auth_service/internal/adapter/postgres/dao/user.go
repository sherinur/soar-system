package dao

import (
	"fmt"
	"time"

	"github.com/sherinur/soar-system/backend/auth_service/internal/model"
)

type User struct {
	ID              int        `db:"id,pk"`
	Email           string     `db:"email,unique"`
	PasswordHash    string     `db:"password_hash"`
	FirstName       string     `db:"first_name"`
	SecondName      string     `db:"second_name"`
	RoleIDs         []int      `db:"role_ids,array"`
	GroupIDs        []int      `db:"group_ids,array"`
	Status          UserStatus `db:"status"`
	OrganizationID  int        `db:"organization_id"`
	CreatedAt       time.Time  `db:"created_at"`
	UpdatedAt       time.Time  `db:"updated_at"`
	LastLoginDate   time.Time  `db:"last_login_date"`
	PasswordExpDate time.Time  `db:"password_exp_date"`
}

func FromUser(user model.User) User {
	return User{
		ID:              user.ID,
		Email:           user.Email,
		PasswordHash:    user.Password,
		FirstName:       user.FirstName,
		SecondName:      user.SecondName,
		RoleIDs:         user.RoleIDs,
		GroupIDs:        user.GroupIDs,
		Status:          UserStatus(user.Status.String()),
		OrganizationID:  user.OrganizationID,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		LastLoginDate:   user.LastLoginDate,
		PasswordExpDate: user.PasswordExpDate,
	}
}

func ToUser(user *User) model.User {
	// status := user.Status.String()
	return model.User{
		ID:         user.ID,
		Email:      user.Email,
		Password:   user.PasswordHash,
		FirstName:  user.FirstName,
		SecondName: user.SecondName,
		RoleIDs:    user.RoleIDs,
		GroupIDs:   user.GroupIDs,
		// Status:          ,
		OrganizationID:  user.OrganizationID,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		LastLoginDate:   user.LastLoginDate,
		PasswordExpDate: user.PasswordExpDate,
	}
}

type UserFilter struct {
	IDs             []int      `db:"ids,array"`
	Email           string     `db:"email"`
	EmailSubstring  string     `db:"email_substring"`
	FirstName       string     `db:"first_name"`
	SecondName      string     `db:"second_name"`
	RoleIDs         []int      `db:"role_ids,array"`
	GroupIDs        []int      `db:"group_ids,array"`
	Status          UserStatus `db:"status"`
	OrganizationID  int        `db:"organization_id"`
	CreatedAfter    time.Time  `db:"created_after"`
	CreatedBefore   time.Time  `db:"created_before"`
	ActiveOnly      bool       `db:"active_only"`
	WithExpiredPass bool       `db:"with_expired_pass"`
}

type UserUpdateData struct {
	Email           *string     `db:"email"`
	Password        *string     `db:"password"`
	FirstName       *string     `db:"first_name"`
	SecondName      *string     `db:"second_name"`
	Roles           *[]int      `db:"roles,array"`
	AddRoles        []int       `db:"add_roles,array"`
	RemoveRoles     []int       `db:"remove_roles,array"`
	Groups          *[]int      `db:"groups,array"`
	AddGroups       []int       `db:"add_groups,array"`
	RemoveGroups    []int       `db:"remove_groups,array"`
	Status          *UserStatus `db:"status"`
	OrganizationID  *int        `db:"organization_id"`
	LastLoginDate   *time.Time  `db:"last_login_date"`
	PasswordExpDate *time.Time  `db:"password_exp_date"`
}

type UserStatus string

const (
	Active      UserStatus = "Active"
	Deactivated UserStatus = "Deactivated"
)

func (s UserStatus) String() string {
	switch s {
	case Active:
		return "Active"
	case Deactivated:
		return "Deactivated"
	default:
		return "Deactivated"
	}
}

func (s *UserStatus) Scan(value interface{}) error {
	if value == nil {
		*s = ""
		return nil
	}
	switch v := value.(type) {
	case string:
		*s = UserStatus(v)
		return nil
	default:
		return fmt.Errorf("failed to scan UserStatus, unexpected type: %T", v)
	}
}

func (s UserStatus) Value() (interface{}, error) {
	return string(s), nil
}
