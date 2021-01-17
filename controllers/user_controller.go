package controllers

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	Response := responses.ListAll(&users)
	Response.Httpstatus = true
	return c.JSON(http.StatusOK, Response)
}

func (controller *UserController) GetUser(c echo.Context) error {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	user, err := controller.UserService.GetUserService(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, id)
	}
	Response := responses.ListOne(&user)
	Response.Httpstatus = true
	return c.JSON(http.StatusOK, Response)
}
