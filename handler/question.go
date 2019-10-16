package handler

import (
	"github.com/Amakuchisan/QuestionBox/model"
	"github.com/labstack/echo"
	"net/http"
)

type (
	questionHandler struct {
		questionModel model.QuestionModelImpl
	}
	// QuestionHandleImplement -- Define handler about questions
	QuestionHandleImplement interface {
		// QuestionAll(c echo.Context) error
		QuestionsPage(c echo.Context) error
	}
)

// NewQuestionHandler -- Initialize handler about question
func NewQuestionHandler(questionModel model.QuestionModelImpl) QuestionHandleImplement {
	return &questionHandler{questionModel}
}

// QuestionsPage -- list questions
func (q *questionHandler) QuestionsPage(c echo.Context) error {
	questions, err := q.questionModel.All()
	if err != nil {
		return err
	}
	return c.Render(http.StatusOK, "question.html", map[string]interface{}{
		"question": questions,
	})
}
