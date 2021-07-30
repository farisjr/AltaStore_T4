package routes

import (
	"project/constants"
	"project/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {
	e.POST("/customers", controllers.CreateCustomersController)
	e.GET("/customers/:id", controllers.GetCustomersidController)
	e.POST("/login", controllers.LoginCustomersControllers)
	//e.DELETE("/users/:id", controllers.DeleteUserController)
	//e.PUT("/users/:id", controllers.UpdateUserController)

	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	r.GET("/customers/:id", controllers.GetCustomersidController)

}
