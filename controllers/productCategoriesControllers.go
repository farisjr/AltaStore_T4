package controllers

import (
	"net/http"
	"project/lib/database"
	"project/models"
	"strconv"

	"github.com/labstack/echo"
)

func CreateProductCategoriesController(c echo.Context) error {
	productCategory := models.ProductCategories{}
	c.Bind(&productCategory)

	productCategoryAdd, err := database.CreateProductCategories(productCategory)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//custom data for body response
	output := map[string]interface{}{
		"CreatedAt": productCategoryAdd.CreatedAt,
		"UpdatedAt": productCategoryAdd.UpdatedAt,
		"DeletedAt": productCategoryAdd.DeletedAt,
		"id":        productCategoryAdd.ID,
		"name":      productCategoryAdd.Name,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create product category",
		"user":    output,
	})
}

func GetProductCategoriesController(c echo.Context) error {
	productcategories, err := database.GetProductCategories()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":            "success get all product categories",
		"product categories": productcategories,
	})
}

func GetProductCategoriesIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	productcategories, err := database.GetProductCategoriesId(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//custom data for body response
	output := map[string]interface{}{
		"CreatedAt": productcategories.CreatedAt,
		"UpdatedAt": productcategories.UpdatedAt,
		"DeletedAt": productcategories.DeletedAt,
		"id":        productcategories.ID,
		"name":      productcategories.Name,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":            "success get product categories by id",
		"product categories": output,
	})
}

func DeleteProductCategoriesByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	productcategories, err := database.GetProductCategoriesId(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	productcategoriesdeleted, err := database.DeleteProductCategoriesById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":               "success delete product selected",
		"product before delete": productcategories,
		"product after delete":  productcategoriesdeleted,
	})
}

func UpdateProductCategoriesController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	productcategories := database.GetUpdateProductCategories(id)
	c.Bind(&productcategories)
	productUpdateCategories, err1 := database.UpdateProductCategories(productcategories)
	if err1 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot post data",
		})
	}

	//custom data for body response
	output := map[string]interface{}{
		"CreatedAt": productUpdateCategories.CreatedAt,
		"UpdatedAt": productUpdateCategories.UpdatedAt,
		"DeletedAt": productUpdateCategories.DeletedAt,
		"id":        productUpdateCategories.ID,
		"name":      productUpdateCategories.Name,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":        "success update product categories",
		"update product": output,
	})
}
