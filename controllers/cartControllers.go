package controllers

import (
	"net/http"
	"project/lib/database"
	"project/models"
	"strconv"

	"github.com/labstack/echo"
)

func CreateCartController(c echo.Context) error {
	// ------------ cart -------------
	// create new cart
	cart := models.Carts{
		StatusTransactions: "ordered",
		TotalQuantity:      0,
		TotalPrice:         0,
		CustomersID:        1,
		PaymentMethodsID:   1,
	}
	c.Bind(&cart)
	newCart, _ := database.CreateCart(cart)

	//------------ cart detail -------------

	// convert product id
	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id product",
		})
	}
	// check product id on table product
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

	//convert qty
	qty, _ := strconv.Atoi(c.Param("qty"))

	//set data cart details
	cartDetails := models.CartDetails{
		ProductsID: productId,
		CartsID:    newCart.ID,
		Quantity:   qty,
		Price:      getProduct.Price,
	}

	//create cart detail
	newCartDetail, _ := database.AddToCart(cartDetails)

	//update total quantity and total price on table carts
	getCart, _ := database.GetCart(newCart.ID)
	newTotalPrice, _ := database.GetTotalPrice(cartDetails.CartsID)
	newTotalQty, _ := database.GetTotalQty(cartDetails.CartsID)
	updateTotalCart, err := database.UpdateTotalCart(getCart.ID, newTotalPrice, newTotalQty)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":      "update total quantity and total price success",
			"cartDetails": updateTotalCart,
		})
	}

	updatedCart, _ := database.GetCart(newCart.ID) //get cart updated (total qty&total price)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":      "create cart success",
		"cart":        updatedCart,
		"cartDetails": newCartDetail,
	})
}

func GetCartController(c echo.Context) error {
	//convert cart_id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id cart",
		})
	}

	//is cart id exist
	var cart models.Carts
	checkCartId, err := database.CheckCartId(id, cart)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":        "cant find cart id",
			"checkProductId": checkCartId,
		})
	}

	listCart, _ := database.GetCartById(id) //get cart by id

	products, _ := database.GetListProductCart(id) //get all products based on cart id

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success get all products by cart id",
		"cart":     listCart,
		"products": products,
	})
}
