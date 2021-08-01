package controllers

import (
	"net/http"
	"project/lib/database"
	"project/models"
	"strconv"

	"github.com/labstack/echo"
)

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

	//get cart by id
	listCart, err := database.GetCartById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//get all products based on cart id
	products, err := database.GetListProductCart(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success get all products by cart id",
		"cart":     listCart,
		"products": products,
	})
}
