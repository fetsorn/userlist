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

func (controller *UserController) AddUser(c echo.Context) error {
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

func (controller *UserController) GetUserByFirst(c echo.Context) error {
	p := c.Param("first")

	user, err := controller.UserService.GetUserByFirstService(p)
	if err != nil {
		return c.JSON(http.StatusBadRequest, p)
	}
	Response := responses.ListOne(&user)
	Response.Httpstatus = true
	return c.JSON(http.StatusOK, Response)
}

func (controller *UserController) GetUserByLast(c echo.Context) error {
	p := c.Param("last")

	user, err := controller.UserService.GetUserByLastService(p)
	if err != nil {
		return c.JSON(http.StatusBadRequest, p)
	}
	Response := responses.ListOne(&user)
	Response.Httpstatus = true
	return c.JSON(http.StatusOK, Response)
}

func (controller *UserController) GetUserByCity(c echo.Context) error {
	p := c.Param("city")

	user, err := controller.UserService.GetUserByCityService(p)
	if err != nil {
		return c.JSON(http.StatusBadRequest, p)
	}
	Response := responses.ListOne(&user)
	Response.Httpstatus = true
	return c.JSON(http.StatusOK, Response)
}

func (controller *UserController) GetUserByCountry(c echo.Context) error {
	p := c.Param("country")

	user, err := controller.UserService.GetUserByCountryService(p)
	if err != nil {
		return c.JSON(http.StatusBadRequest, p)
	}
	Response := responses.ListOne(&user)
	Response.Httpstatus = true
	return c.JSON(http.StatusOK, Response)
}
