package model

import (
	"time"
)

// Question -- This is question model
type Question struct {
	ID        uint64     `db:"id"`
	Title     string     `db:"title"`
	Body      string     `db:"body"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
	UID       uint64     `db:"user_id"`
}

// QuestionModelImpl -- Define database control methods about question
type QuestionModelImpl interface {
	All() ([]Question, error)
}
