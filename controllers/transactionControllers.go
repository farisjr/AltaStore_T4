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
	c.Bind(&cart)
	statusTransactionUpdate, err := database.UpdateTransactionStatus(cart)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success updating transaction status",
		"data":    statusTransactionUpdate,
	})

}
