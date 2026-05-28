package http_handler

import (
	"net/http"
	"study/internal/service"
	"study/models"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

func (h Handler) PostUser(c echo.Context) error {
	var user models.User

	c.Bind(&user)

	h.service.CreateUser(user)

	return c.String(200, "user saved")
}

func (h Handler) GetUsers(c echo.Context) error {
	users, err := h.service.GetUsers()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(200, users)
}

func (h Handler) GetUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	user, err := h.service.GetUser(user.Name)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(200, user)
}
