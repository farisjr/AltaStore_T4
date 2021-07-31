package routes

import (
	"project/constants"
	"project/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {
	e.GET("/paymentMethods", controllers.GetAllPaymentMethodsController)
	e.GET("/paymentMethods/:id", controllers.GetOnePaymentMethodsController)
	e.POST("/paymentMethods", controllers.CreatePaymentMethodsController)
	e.PUT("/paymentMethods/:id", controllers.UpdatePaymentMethodsController)
	e.DELETE("/paymentMethods/:id", controllers.DeletePaymentMethodsController)

	e.GET("/sellers", controllers.GetAllSellersController)
	e.GET("/sellers/:id", controllers.GetOneSellersController)
	e.POST("/sellers", controllers.CreateSellersController)
	e.PUT("/sellers/:id", controllers.UpdateSellersController)
	e.DELETE("/sellers/:id", controllers.DeleteSellersController)

	e.POST("/customers", controllers.CreateCustomersController)
	e.GET("/customers/:id", controllers.GetCustomersidController)
	e.POST("/login", controllers.LoginCustomersControllers)
	//e.DELETE("/users/:id", controllers.DeleteUserController)
	//e.PUT("/users/:id", controllers.UpdateUserController)

	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	r.GET("/customers/:id", controllers.GetCustomersidController)

}
