package controllers

import (
	"net/http"
	"project/lib/database"
	"project/middlewares"
	"project/models"
	"strconv"

	"github.com/labstack/echo"
)

func CreateCartController(c echo.Context) error {

	// ------------ cart -------------//
	//rec user input
	cart := models.Carts{}
	c.Bind(&cart) //input: payment method id

	// get id user login
	customerId := middlewares.ExtractTokenCustomerId(c)

	//check product id on table product
	paymentId := cart.PaymentMethodsID
	var payment models.PaymentMethods
	checkPayment, err := database.CheckPayment(paymentId, payment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":        "cant find payment method",
			"checkProductId": checkPayment,
		})
	}

	//set data cart and create new cart
	cart = models.Carts{
		StatusTransactions: "ordered",
		TotalQuantity:      0,
		TotalPrice:         0,
		CustomersID:        customerId,
		PaymentMethodsID:   paymentId,
	}
	newCart, _ := database.CreateCart(cart)

	//------------ cart detail -------------//
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
	UpdateTotalCart(newCart.ID)

	//get cart updated (total qty&total price)
	updatedCart, _ := database.GetCart(newCart.ID)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":      "create cart success",
		"cart":        updatedCart,
		"cartDetails": newCartDetail,
	})
}

//func for update total quantity and total price on table carts
func UpdateTotalCart(cartId int) (int, int) {
	newTotalPrice, _ := database.GetTotalPrice(cartId)
	newTotalQty, _ := database.GetTotalQty(cartId)
	newCart, _ := database.UpdateTotalCart(cartId, newTotalPrice, newTotalQty)

	return newCart.TotalQuantity, newCart.TotalPrice
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

func DeleteCartController(c echo.Context) error {
	//convert cart id
	cartId, err := strconv.Atoi(c.Param("id"))
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

	//delete cart and products included on it
	deletedCart, _ := database.DeleteCart(cartId)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "delete cart success",
		"Deleted Cart": deletedCart,
	})
}
