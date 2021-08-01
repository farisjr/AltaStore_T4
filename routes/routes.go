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

	e.POST("/customers", controllers.CreateCustomersController)
	e.GET("/customers/:id", controllers.GetCustomersidController)
	e.POST("/login", controllers.LoginCustomersControllers)
	//e.DELETE("/users/:id", controllers.DeleteUserController)
	//e.PUT("/users/:id", controllers.UpdateUserController)

	//carts
	e.POST("api/carts/:cartId/details", controllers.AddToCartController) //add product to cart
	e.GET("/api/carts/:id", controllers.GetCartController)
	e.DELETE("/api/cartDetails/:carts_id/:products_id", controllers.DeleteProductFromCartController)
	// e.POST("api/carts", controllers.CreateCartController)

	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	r.GET("/customers/:id", controllers.GetCustomersidController)

}
