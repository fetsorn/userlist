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

func (controller *UserController) CreateUser(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, u)
	}
	if err := c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	result, err := controller.UserService.AddUserService(u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, u)
	}
	return c.JSON(http.StatusOK, result)
}

func (controller *UserController) DeleteUser(c echo.Context) error {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	deleteResult, err := controller.UserService.DeleteUserService(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, id)
	}
	return c.JSON(http.StatusOK, deleteResult)
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
