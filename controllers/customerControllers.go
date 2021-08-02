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
	customers, err := database.GetCustomersid(id)
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

func DeleteCustomersByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	customers, err := database.GetCustomersid(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	customersdeleted, err := database.DeleteCustomersById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":               "success delete product selected",
		"product before delete": customers,
		"product after delete":  customersdeleted,
	})
}

func UpdateCustomersController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	customers := database.GetUpdateCustomers(id)
	c.Bind(&customers)
	customersUpdate, err1 := database.UpdateCustomers(customers)
	if err1 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot post data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":         "success update product",
		"update customer": customersUpdate,
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
