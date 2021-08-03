package routes

import (
	"project/constants"
	"project/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {

	e.PUT("/api/transactions/:id", controllers.UpdateTransactionStatusController)

	//------------------Non Authorized----------------------//
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
	e.DELETE("/customers/:id", controllers.DeleteCustomersByIdController)
	e.PUT("/customers/:id", controllers.UpdateCustomersController)

	e.POST("/products", controllers.CreateProductsController)
	e.GET("/products", controllers.GetProductsController)
	e.GET("/products/:id", controllers.GetProductidController)
	e.DELETE("/products/:id", controllers.DeleteProductByIdController)
	e.PUT("/products/:id", controllers.UpdateProductController)

	//carts
	e.POST("/api/cart/:productId/:qty", controllers.CreateCartController)                            // create new shopping cart
	e.POST("api/carts/:cartId/details", controllers.AddToCartController)                             //add product to cart
	e.GET("/api/carts/:id", controllers.GetCartController)                                           //get all product on a cart
	e.DELETE("/api/cartDetails/:carts_id/:products_id", controllers.DeleteProductFromCartController) //add product to existing shopping cart

	//----------------Authorized Only----------------------//
	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	r.GET("/customers/:id", controllers.GetCustomersidController)

	//carts
	r.POST("/carts/:productId/:qty", controllers.CreateCartController)                           // create new shopping cart
	r.POST("/carts/:cartId/details", controllers.AddToCartController)                            //add product to cart
	r.GET("/carts/:id", controllers.GetCartController)                                           //get all product on a cart
	r.DELETE("/cartDetails/:carts_id/:products_id", controllers.DeleteProductFromCartController) //delete product from cart

	r.GET("/products/productcategories/:name", controllers.GetProductByProductCategoryController)

	r.POST("/productcategories", controllers.CreateProductCategoriesController)
	r.GET("/productcategories", controllers.GetProductCategoriesController)
	r.GET("/productcategories/:id", controllers.GetProductCategoriesIdController)
	r.DELETE("/productcategories/:id", controllers.DeleteProductCategoriesByIdController)
	r.PUT("/productcategories/:id", controllers.UpdateProductCategoriesController)
}
