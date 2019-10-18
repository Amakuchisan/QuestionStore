package repository

import (
	"github.com/Amakuchisan/QuestionStore/model"
	"github.com/jmoiron/sqlx"
)

// QuestionModel -- have method about user table in database
type QuestionModel struct {
	db *sqlx.DB
}

// NewQuestionModel QuestionModelの初期化
func NewQuestionModel(db *sqlx.DB) *QuestionModel {
	return &QuestionModel{db: db}
}

// All -- SELECT * FROM questions
func (q *QuestionModel) All() ([]model.Question, error) {
	questions := []model.Question{}
	err := q.db.Select(&questions, "SELECT * from question")
	if err != nil {
		return nil, err
	}
	return questions, nil
}
