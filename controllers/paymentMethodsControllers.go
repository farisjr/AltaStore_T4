package controllers

import (
	"net/http"
	"project/lib/database"
	"project/models"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllPaymentMethodsController(c echo.Context) error {
	paymentMethods, err := database.GetPaymentMethods()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success get all payment method",
		"data":   paymentMethods,
	})
}

func GetOnePaymentMethodsController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	paymentMethods, err := database.GetPaymentMethodsById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//custom data for body response
	output := map[string]interface{}{
		"CreatedAt": paymentMethods.CreatedAt,
		"UpdatedAt": paymentMethods.UpdatedAt,
		"DeletedAt": paymentMethods.DeletedAt,
		"id":        paymentMethods.ID,
		"name":      paymentMethods.Name,
		"status":    paymentMethods.Status,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get payment method data",
		"data":    output,
	})
}

func CreatePaymentMethodsController(c echo.Context) error {
	// binding data
	paymentMethod := models.PaymentMethods{}
	c.Bind(&paymentMethod)
	paymentMethods, err := database.CreatePaymentMethods(paymentMethod)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//custom data for body response
	output := map[string]interface{}{
		"CreatedAt": paymentMethods.CreatedAt,
		"UpdatedAt": paymentMethods.UpdatedAt,
		"DeletedAt": paymentMethods.DeletedAt,
		"id":        paymentMethods.ID,
		"name":      paymentMethods.Name,
		"status":    paymentMethods.Status,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create payment methods",
		"data":     output,
	})
}

func DeletePaymentMethodsController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	paymentMethods, err := database.DeletePaymentMethodsById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete payment method",
		"data":    paymentMethods,
	})

}

func UpdatePaymentMethodsController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	paymentMethod := database.GetUpdatePaymentMethod(id)
	c.Bind(&paymentMethod)
	paymentMethodUpdate, err := database.UpdatePaymentMethod(paymentMethod)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//custom data for body response
	output := map[string]interface{}{
		"CreatedAt": paymentMethodUpdate.CreatedAt,
		"UpdatedAt": paymentMethodUpdate.UpdatedAt,
		"DeletedAt": paymentMethodUpdate.DeletedAt,
		"id":        paymentMethodUpdate.ID,
		"name":      paymentMethodUpdate.Name,
		"status":    paymentMethodUpdate.Status,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success updating payment method",
		"data":    output,
	})
}
