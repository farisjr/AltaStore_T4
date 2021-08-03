package controllers

import (
	"net/http"
	"project/lib/database"
	"strconv"

	"github.com/labstack/echo"
)

func UpdateTransactionStatusController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	cart := database.GetTransactionStatus(id)
	cart.StatusTransactions = "paid"
	c.Bind(&cart)
	statusTransactionUpdate, err := database.UpdateTransactionStatus(cart)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//custom data cart for body response
	outputCart := map[string]interface{}{
		"ID":                  statusTransactionUpdate.ID,
		"customers_id":        statusTransactionUpdate.CustomersID,
		"payment_methods_id":  statusTransactionUpdate.PaymentMethodsID,
		"status_transactions": statusTransactionUpdate.StatusTransactions,
		"total_quantity":      statusTransactionUpdate.TotalQuantity,
		"total_price":         statusTransactionUpdate.TotalPrice,
		"CreatedAt":           statusTransactionUpdate.CreatedAt,
		"UpdatedAt":           statusTransactionUpdate.UpdatedAt,
		"DeletedAt":           statusTransactionUpdate.DeletedAt,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success updating transaction status",
		"data":    outputCart,
	})

}
