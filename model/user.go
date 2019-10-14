package model

import (
	"time"
)

// User -- This is user model
type User struct {
	ID        uint64     `db:"id"`
	Email     string     `db:"email"`
	Name      string     `db:"name"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

// UserModelImpl -- Define database control methods about user
type UserModelImpl interface {
	All() ([]User, error)
	Create(user *User) error
	// Delete(user *User) error
	// FindByID(id uint) (User, error)
}
