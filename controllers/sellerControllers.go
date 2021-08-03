package controllers

import (
	"net/http"
	"project/lib/database"
	"project/models"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllSellersController(c echo.Context) error {
	sellers, err := database.GetSellers()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success get all sellers",
		"data":   sellers,
	})
}

func GetOneSellersController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	sellers, err := database.GetSellersById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get sellers data",
		"data":    sellers,
	})
}

func CreateSellersController(c echo.Context) error {
	seller := models.Sellers{}
	c.Bind(&seller)

	sellerAdd, err := database.CreateSeller(seller)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new seller",
		"data":    sellerAdd,
	})
}

func DeleteSellerByIdController(c echo.Context) error {
	// binding data
	seller := models.Sellers{}
	c.Bind(&seller)
	sellers, err := database.CreateSellers(seller)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create new seller",
		"data":     sellers,
	})
}

func DeleteSellersController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	sellerDeleted, err := database.DeleteSellers(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete selected seller",
		"data":    sellerDeleted,
	})
}

func UpdateSellerController(c echo.Context) error {
	sellers, err := database.DeleteSellersById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete seller",
		"data":    sellers,
	})

}

func UpdateSellersController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	sellers := database.GetUpdateSellers(id)
	c.Bind(&sellers)
	sellerUpdate, err1 := database.UpdateSellers(sellers)
	if err1 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot post data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update seller",
		"data":    sellerUpdate,
	})
}

func LoginSellerControllers(c echo.Context) error {
	seller := models.Sellers{}
	c.Bind(&seller)
	sellers, err := database.LoginSellers(seller.Email, seller.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "succes login",
		"data":   sellers,
	})
}
