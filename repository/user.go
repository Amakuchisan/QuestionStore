package repository

import (
	"github.com/Amakuchisan/QuestionBox/model"
	"github.com/jmoiron/sqlx"
)

// UserModel -- have method about user table in database
type UserModel struct {
	db *sqlx.DB
}

// NewUserModel UserModelの初期化
func NewUserModel(db *sqlx.DB) *UserModel {
	return &UserModel{db: db}
}

// All -- SELECT * FROM users
func (u *UserModel) All() ([]model.User, error) {
	users := []model.User{}
	err := u.db.Select(&users, "SELECT * from user")
	if err != nil {
		return nil, err
	}
	return users, nil
}
