package handler

import (
	"github.com/Amakuchisan/QuestionBox/model"
	"github.com/labstack/echo"
	"net/http"
)

type (
	userHandler struct {
		userModel model.UserModelImpl
	}
	// UserHandleImplement -- Define handler about users
	UserHandleImplement interface {
		UserAll(c echo.Context) error
	}
)

// NewUserHandler -- Initialize handler about user
func NewUserHandler(userModel model.UserModelImpl) UserHandleImplement {
	return &userHandler{userModel}
}

// UserAll -- user page
func (u *userHandler) UserAll(c echo.Context) error {
	// users := []model.User{}
	users, err := u.userModel.All()
	// err := database.DB.Select(&users, "SELECT * from user")
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}
