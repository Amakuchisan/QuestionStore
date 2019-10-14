package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/guni1192/miniP/model"
)

// handler -- user api handler
type handler struct {
	userModel model.UserModelImpl
}

// HandleImplement -- user's handler for Interface
type HandleImplement interface {
	GetUsers(c echo.Context) error
	CreateUser(c echo.Context) error
	DetailUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}

// NewUserHandler is new handler for user
func NewUserHandler(userModel model.UserModelImpl) HandleImplement {
	return &handler{userModel}
}

// GetUser for GET /api/v1/users
func (h *handler) GetUsers(c echo.Context) error {
	users, err := h.userModel.All()
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, users)
}

// CreateUser for POST /api/v1/users
func (h *handler) CreateUser(c echo.Context) error {
	user := model.User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := h.userModel.Create(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusCreated, user)
}

// DetailUser for GET /api/v1/users/:id
func (h *handler) DetailUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err := h.userModel.FindByID(uint(id))

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, user)
}

// DeleteUser for DELETE /api/v1/users
func (h *handler) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user := model.User{}
	user.ID = uint(id)
	if err = h.userModel.Delete(&user); err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, user)
}
