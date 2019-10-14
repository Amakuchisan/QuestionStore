package model

import (
	"github.com/jmoiron/sqlx"
	"time"
)

// User -- This is user model
type User struct {
	ID        uint64     `db:"id"`
	Name      string     `db:"name"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

// UserModel -- have method about user table in database
type UserModel struct {
	db *sqlx.DB
}

// UserModelImpl -- Define database control methods about user
type UserModelImpl interface {
	All() ([]User, error)
	// Create(user *User) error
	// Delete(user *User) error
	// FindByID(id uint) (User, error)
}

// NewUserModel UserModelの初期化
func NewUserModel(db *sqlx.DB) *UserModel {
	return &UserModel{db: db}
}

// All -- SELECT * FROM users
func (u *UserModel) All() ([]User, error) {
	users := []User{}
	err := u.db.Select(&users, "SELECT * from user")
	if err != nil {
		return nil, err
	}
	return users, nil
}
