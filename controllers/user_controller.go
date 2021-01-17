package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"userlist/models"
	"userlist/responses"
	"userlist/services"
)

type UserController struct {
	UserService services.UserService
}

var controller *UserController

func NewController() *UserController {
	controller = &UserController{services.NewService()}
	return controller
}

func (controller *UserController) GetUsers(c echo.Context) error {
	var users models.Users
	users, err := controller.UserService.GetUsersService()
	if err != nil {
		return c.JSON(http.StatusBadRequest, users)
	}
	Response := responses.Listalluser(&users)
	Response.Httpstatus = true
	return c.JSON(http.StatusOK, Response)
}
