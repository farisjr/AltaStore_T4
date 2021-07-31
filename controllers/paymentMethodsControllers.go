package controllers

import (
	"net/http"
	"project/lib/database"
	"project/models"
	"strconv"

	"github.com/labstack/echo"
)


func GetAllController(c echo.Context) error {
	paymentMethods, err := database.GetPaymentMethods()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success get all payment method",
		"users":  paymentMethods,
	})
}

func GetOneController(c echo.Context) error {
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
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get payment method data",
		"user":    paymentMethods,
	})
}

func CreateController(c echo.Context) error {
	// binding data
	paymentMethod := models.PaymentMethods{}
	c.Bind(&paymentMethod)
	paymentMethods, err := database.CreatePaymentMethods(paymentMethod)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create payment methods",
		"user":     paymentMethods,
	})
}

func DeleteController(c echo.Context) error {
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
		"user":    paymentMethods,
	})

}

func UpdateController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	paymentMethod := models.PaymentMethods{}
	c.Bind(&paymentMethod)
	paymentMethods, err := database.UpdatePaymentMethodsById(paymentMethod, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success updating payment method",
		"data":    paymentMethods,
	})
}
