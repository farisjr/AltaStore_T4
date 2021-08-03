package database

import (
	"project/config"
	"project/models"
)

//check is product exist on table product
func CheckProductId(productId int, product models.Products) (interface{}, error) {
	if err := config.DB.Where("id=?", productId).First(&product).Error; err != nil {
		return nil, err
	}
	return product.ID, nil
}

// get product by id
func GetProduct(productId int) (models.Products, error) {
	var product models.Products
	if err := config.DB.Find(&product, "id=?", productId).Error; err != nil {
		return product, err
	}
	return product, nil
}

func GetProducts() (interface{}, error) {
	var products []models.Products
	if err := config.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductid(id int) (interface{}, error) {
	var product models.Products
	var count int64
	if err1 := config.DB.Model(&product).Where("id=?", id).Count(&count).Error; count == 0 {
		return nil, err1
	}
	if err := config.DB.Find(&product, "id=?", id).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func CreateProduct(products models.Products) (interface{}, error) {
	if err := config.DB.Save(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductByProductCategory(name string) (interface{}, error) {
	var productcategories models.ProductCategories
	if err := config.DB.Where("name=?", name).First(&productcategories).Error; err != nil {
		return nil, err
	}
	var products []models.Products
	if err := config.DB.Find(&products, "product_categories_id=?", productcategories.ID).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func DeleteProductById(id int) (interface{}, error) {
	var products models.Products
	if err := config.DB.Where("id=?", id).Delete(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

//update user info from database
func UpdateProduct(product models.Products) (interface{}, error) {
	if tx := config.DB.Save(&product).Error; tx != nil {
		return nil, tx
	}
	return product, nil
}

//get 1 specified user with User struct output
func GetUpdateProduct(id int) models.Products {
	var product models.Products
	config.DB.Find(&product, "id=?", id)
	return product
}
