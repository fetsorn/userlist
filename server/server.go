package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"userlist/controllers"
	"userlist/models"
)

func Routes() (e *echo.Echo) {
	e = echo.New()
	e.Debug = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	var Controllers = controllers.NewController()
	e.Validator = &models.UserValidator{validator.New()}
	g := e.Group("/api/v1/users", UserCtx)
	g.DELETE("/id/:id", Controllers.DeleteUser)
	g.GET("/id/:id", Controllers.GetUser)
	g.GET("/first/:first", Controllers.GetUserByFirst)
	g.GET("/first/:first", Controllers.GetUserByFirst)
	g.GET("/last/:last", Controllers.GetUserByLast)
	g.GET("/city/:city", Controllers.GetUserByCity)
	g.GET("/country/:country", Controllers.GetUserByCountry)
	g.GET("", Controllers.GetUsers)
	g.POST("", Controllers.AddUser)
	return e
}
func UserCtx(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

func Start() {
	e := Routes()
	http.Handle("/", e)
	e.Logger.Fatal(e.Start(":4000"))
}
