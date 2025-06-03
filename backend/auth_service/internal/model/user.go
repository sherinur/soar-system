package model

import "time"

type User struct {
	ID              int
	Email           string
	Password        string
	FirstName       string
	SecondName      string
	RoleIDs         []int
	GroupIDs        []int
	Status          UserStatus
	OrganizationID  int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	LastLoginDate   time.Time
	PasswordExpDate time.Time
}

func (u *User) Validate() error {
	// TODO: Implement validation rules
	return nil
}

type UserFilter struct {
	IDs             []int
	Email           string
	EmailSubstring  string
	FirstName       string
	SecondName      string
	RoleIDs         []int
	GroupIDs        []int
	Status          UserStatus
	OrganizationID  int
	CreatedAfter    time.Time
	CreatedBefore   time.Time
	ActiveOnly      bool
	WithExpiredPass bool
}

type UserUpdateData struct {
	Email           *string
	Password        *string
	FirstName       *string
	SecondName      *string
	Roles           *[]int
	AddRoles        []int
	RemoveRoles     []int
	Groups          *[]int
	AddGroups       []int
	RemoveGroups    []int
	Status          *UserStatus
	OrganizationID  *int
	LastLoginDate   *time.Time
	PasswordExpDate *time.Time
}

// User Status enum
type UserStatus int

const (
	Active UserStatus = iota
	Deactivated
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
