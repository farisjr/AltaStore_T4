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
<<<<<<< HEAD
	getProduct, err := database.GetProduct(productId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
=======
	getProduct, _ := database.GetProduct(productId)
>>>>>>> origin/feature_add_new_cart

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
<<<<<<< HEAD
	getCart, err := database.GetCart(cartId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	newTotalPrice, err := database.GetTotalPrice(cartDetails.CartsID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	newTotalQty, err := database.GetTotalQty(cartDetails.CartsID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

=======
	getCart, _ := database.GetCart(cartId)
	newTotalPrice, _ := database.GetTotalPrice(cartDetails.CartsID)
	newTotalQty, _ := database.GetTotalQty(cartDetails.CartsID)
>>>>>>> origin/feature_add_new_cart
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

	//delete product
<<<<<<< HEAD
	deleteProduct, err := database.DeleteProductFromCart(cartId, productId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//update total quantity and total price on table carts
	getCart, err := database.GetCart(cartId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//
	getCartDetailByCartId, err := database.GetCartDetailByCartId(cartId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	newTotalPrice, err := database.GetTotalPrice(getCartDetailByCartId.CartsID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	newTotalQty, err := database.GetTotalQty(getCartDetailByCartId.CartsID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

=======
	deleteProduct, _ := database.DeleteProductFromCart(cartId, productId)

	//update total quantity and total price on table carts
	getCart, _ := database.GetCart(cartId)
	getCartDetailByCartId, _ := database.GetCartDetailByCartId(cartId)
	newTotalPrice, _ := database.GetTotalPrice(getCartDetailByCartId.CartsID)
	newTotalQty, _ := database.GetTotalQty(getCartDetailByCartId.CartsID)
>>>>>>> origin/feature_add_new_cart
	updateTotalCart, err := database.UpdateTotalCart(getCart.ID, newTotalPrice, newTotalQty)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":      "update total quantity and total price success",
			"cartDetails": updateTotalCart,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "delete product on table cart_details success",
		"product": deleteProduct,
	})
}
