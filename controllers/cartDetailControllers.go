package controllers

import (
	"net/http"
	"project/lib/database"
	"project/models"
	"strconv"

	"github.com/labstack/echo"
)

func AddToCartController(c echo.Context) error {
	var cart models.Carts

	//check id cart
	cartId, err := strconv.Atoi(c.Param("cartId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id cart",
		})
	}
	checkCartId, err := database.CheckCartId(cartId, cart)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":     "cant find cart",
			"checkCartId": checkCartId,
		})
	}

	// record user's input
	var cartDetails models.CartDetails
	c.Bind(&cartDetails) //input: productid, qty

	//check product id on table product
	productId := cartDetails.ProductsID //get product_id
	var product models.Products
	checkProductId, err := database.CheckProductId(productId, product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":        "cant find product",
			"checkProductId": checkProductId,
		})
	}

	//get price
	getProduct, err := database.GetProduct(productId)

	//set data cart details
	cartDetails = models.CartDetails{
		ProductsID: productId,
		CartsID:    cartId,
		Quantity:   cartDetails.Quantity,
		Price:      getProduct.Price,
	}

	//create cart detail
	newCartDetail, err := database.AddToCart(cartDetails)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//update total quantity and total price on table carts
	getCart, err := database.GetCart(cartId)
	newTotalPrice, err := database.GetTotalPrice(cartDetails.CartsID)
	newTotalQty, err := database.GetTotalQty(cartDetails.CartsID)

	updateTotalCart, err := database.UpdateTotalCart(getCart.ID, newTotalPrice, newTotalQty)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":      "update total quantity and total price success",
			"cartDetails": updateTotalCart,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":      "add product to cart success",
		"cartDetails": newCartDetail,
	})
}
