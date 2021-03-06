package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Amakuchisan/QuestionStore/model"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/objx"
)

type (
	questionHandler struct {
		questionModel model.QuestionModelImpl
	}
	// QuestionHandleImplement -- Define handler about questions
	QuestionHandleImplement interface {
		QuestionDetail(c echo.Context) error
		QuestionsTitleList(c echo.Context) error
		PostQuestion(c echo.Context) error
	}
)

// NewQuestionHandler -- Initialize handler about question
func NewQuestionHandler(questionModel model.QuestionModelImpl) QuestionHandleImplement {
	return &questionHandler{questionModel}
}

// QuestionFormHandler -- post question
func QuestionFormHandler(c echo.Context) error {
	auth, err := c.Cookie("auth")
	if err != nil {
		return err
	}
	userData := objx.MustFromBase64(auth.Value)

	return c.Render(http.StatusOK, "form", map[string]interface{}{
		"title": "CreateQuestion",
		"name":  userData["name"],
	})
}

// QuestionDetail -- show question
func (q *questionHandler) QuestionDetail(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)

	if err != nil {
		return err
	}

	question, err := q.questionModel.FindByID(id)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "quest", map[string]interface{}{
		"title":    question.Title,
		"question": question,
	})
}

// QuestionsTitleList -- list questions
func (q *questionHandler) QuestionsTitleList(c echo.Context) error {
	questions, err := q.questionModel.All()
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "question", map[string]interface{}{
		"title":    "Question",
		"question": questions,
	})
}

func (q *questionHandler) PostQuestion(c echo.Context) error {
	auth, err := c.Cookie("auth")
	if err != nil {
		return err
	}

	convertedUserData := make(map[string]string)
	pureUserData := objx.MustFromBase64(auth.Value)

	for key, value := range pureUserData {
		switch value := value.(type) {
		case string:
			convertedUserData[key] = value
		default:
			convertedUserData[key] = fmt.Sprintf("%v", value)
		}
	}

	subject := c.FormValue("subject")
	body := c.FormValue("question")
	uid, err := strconv.ParseUint(convertedUserData["id"], 10, 64)
	if err != nil {
		return err
	}

	question := model.Question{Title: subject, Body: body, UID: uid}

	err = q.questionModel.Create(&question)
	if err != nil {
		return err
	}
	return c.Redirect(http.StatusSeeOther, "/questions")
}
