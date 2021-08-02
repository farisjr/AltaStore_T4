package controllers

import (
	"fmt"
	"net/http"
	"project/lib/database"
	"project/models"

	"strconv"

	"github.com/labstack/echo"
)

//get all users info
func GetProductsController(c echo.Context) error {
	products, err := database.GetProducts()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success get all products",
		"products": products,
	})
}

func GetProductidController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	products, err := database.GetProductid(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success get product by id",
		"products": products,
	})
}

func CreateProductsController(c echo.Context) error {
	var product models.Products
	//product := models.Products{}
	c.Bind(&product)
	fmt.Println(product)
	productAdd, err := database.CreateProduct(product)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "success create new product",
		"product add": productAdd,
	})
}

func GetProductByProductCategoryController(c echo.Context) error {
	name := c.Param("name")
	productbycategories, err := database.GetProductByProductCategory(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	fmt.Println(productbycategories)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":             "success get all product by product category",
		"product by category": productbycategories,
	})
}

func DeleteProductByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	products, err := database.GetProductid(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	productdeleted, err := database.DeleteProductById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":               "success delete product selected",
		"product before delete": products,
		"product after delete":  productdeleted,
	})
}

func UpdateProductController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	product := database.GetUpdateProduct(id)
	c.Bind(&product)
	productUpdate, err1 := database.UpdateProduct(product)
	if err1 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot post data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":        "success update product",
		"update product": productUpdate,
	})
}
