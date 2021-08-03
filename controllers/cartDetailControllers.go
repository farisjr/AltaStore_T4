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

	//check id cart is exist
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
	getProduct, _ := database.GetProduct(productId)

	//set data cart details
	cartDetails = models.CartDetails{
		ProductsID: productId,
		CartsID:    cartId,
		Quantity:   cartDetails.Quantity,
		Price:      getProduct.Price,
	}

	//create cart detail
	newCartDetail, _ := database.AddToCart(cartDetails)

	//update total quantity and total price on table carts
	newTotalQty, newTotalPrice := UpdateTotalCart(cartId)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":         "add product to cart success",
		"cartDetails":    newCartDetail,
		"Total Quantity": newTotalQty,
		"Total Price":    newTotalPrice,
	})
}

func DeleteProductFromCartController(c echo.Context) error {
	//convert cart id
	cartId, err := strconv.Atoi(c.Param("carts_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid cart id",
		})
	}

	//check is cart id exist on table cart
	var cart models.Carts
	checkCartId, err := database.CheckCartId(cartId, cart)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":     "cant find cart",
			"checkCartId": checkCartId,
		})
	}

	//convert product id
	productId, err := strconv.Atoi(c.Param("products_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid product id",
		})
	}

	//check is product id exist on table product
	var product models.Products
	checkProductId, err := database.CheckProductId(productId, product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":     "cant find product",
			"checkCartId": checkProductId,
		})
	}

	//check is product id and cart id exist on table cart_detail
	var cartDetails models.CartDetails
	checkProductAndCartId, err := database.CheckProductAndCartId(productId, cartId, cartDetails)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":     "cant find product id and cart id",
			"checkCartId": checkProductAndCartId,
		})
	}

	//---------delete product------//
	countProduct, _ := database.CountProductOnCart(cartId) //count product
	var deleteProduct interface{}
	var newTotalQty, newTotalPrice int

	if countProduct > 1 { //if product on cart > 1, delete product on cart detail + update total on cart
		deleteProduct, _ = database.DeleteProductFromCart(cartId, productId)
		newTotalQty, newTotalPrice = UpdateTotalCart(cartId)
	} else if countProduct == 1 { //if product only 1, delete product on cart detail + delete cart + output total = 0
		deleteProduct, _ = database.DeleteProductFromCart(cartId, productId)
		database.DeleteCart(cartId)
		newTotalPrice = 0
		newTotalQty = 0
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":         "delete product on table cart_details success",
		"Deleted Product": deleteProduct,
		"Total Quantity":  newTotalQty,
		"Total Price":     newTotalPrice,
	})
}
