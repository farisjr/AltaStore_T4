package controllers

import (
	"net/http"
	"project/lib/database"
	"project/models"

	"strconv"

	"github.com/labstack/echo"
)

func GetCustomersidController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	customers, err := database.GetCustomers(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   customers,
	})
}

func CreateCustomersController(c echo.Context) error {
	customer := models.Customers{}
	c.Bind(&customer)

	customerAdd, err := database.CreateCustomer(customer)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    customerAdd,
	})
}

func LoginCustomersControllers(c echo.Context) error {
	customer := models.Customers{}
	c.Bind(&customer)
	customers, err := database.LoginCustomers(customer.Email, customer.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "succes login",
		"users":  customers,
	})
}
