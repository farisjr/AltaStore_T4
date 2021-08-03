package routes

import (
	"project/constants"
	"project/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {

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

	//----------------Authorized Only----------------------//
	r := e.Group("")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	r.GET("/customers/:id", controllers.GetCustomersidController)

	//carts
	r.POST("/carts/:productId/:qty", controllers.CreateCartController)                           // create new shopping cart
	r.POST("/carts/:cartId/details", controllers.AddToCartController)                            //add product to cart
	r.GET("/carts/:id", controllers.GetCartController)                                           //get all product on a cart
	r.DELETE("/carts/:id", controllers.DeleteCartController)                                     //delete cart and all products included on it
	r.DELETE("/cartDetails/:carts_id/:products_id", controllers.DeleteProductFromCartController) //delete product from cart

	r.PUT("/transactions/:id", controllers.UpdateTransactionStatusController)

	r.GET("/products/productcategories/:name", controllers.GetProductByProductCategoryController)

	r.POST("/productcategories", controllers.CreateProductCategoriesController)
	r.GET("/productcategories", controllers.GetProductCategoriesController)
	r.GET("/productcategories/:id", controllers.GetProductCategoriesIdController)
	r.DELETE("/productcategories/:id", controllers.DeleteProductCategoriesByIdController)
	r.PUT("/productcategories/:id", controllers.UpdateProductCategoriesController)
}
