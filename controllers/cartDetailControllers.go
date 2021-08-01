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
<<<<<<< Updated upstream
	getProduct, err := database.GetProduct(productId)
=======
	getProduct, _ := database.GetProduct(productId)
>>>>>>> Stashed changes

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
<<<<<<< Updated upstream
	getCart, err := database.GetCart(cartId)
	newTotalPrice, err := database.GetTotalPrice(cartDetails.CartsID)
	newTotalQty, err := database.GetTotalQty(cartDetails.CartsID)

=======
	getCart, _ := database.GetCart(cartId)
	newTotalPrice, _ := database.GetTotalPrice(cartDetails.CartsID)
	newTotalQty, _ := database.GetTotalQty(cartDetails.CartsID)
>>>>>>> Stashed changes
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
	deleteProduct, _ := database.DeleteProductFromCart(cartId, productId)

	//update total quantity and total price on table carts
	getCart, _ := database.GetCart(cartId)
	getCartDetailByCartId, _ := database.GetCartDetailByCartId(cartId)
	newTotalPrice, _ := database.GetTotalPrice(getCartDetailByCartId.CartsID)
	newTotalQty, _ := database.GetTotalQty(getCartDetailByCartId.CartsID)
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
