# Application Programming Interface Documentation

## List
| No | Web Service | Method | URL |
|----|-------------|--------|-----|
| 1 | [Register](#register) | POST | /customers |
| 2 | [Login](#login) | POST | /login |
| 3 | [Get customer](#get-customer) | GET | /customers/:id |
| 4 | [Edit customer](#edit-customer) | PUT | /customers/:id |
| 5 | [Delete customer](#delete-customer) | DELETE | /customers/:id |
| 6 | [Create category](#create-category) | POST | /productcategories |
| 7 | [Get all categories](#get-all-categories) | GET | /productcategories |
| 8 | [Get category](#get-category) | GET | /productcategories/:id |
| 9 | [Edit category](#edit-category) | PUT | /productcategories/:id |
| 10 | [Delete category](#delete-category) | DELETE | /productcategories/:id |
| 11 | [Create product](#create-product) | POST | /products |
| 12 | [Get all products](#get-all-products) | GET | /products |
| 13 | [Get product](#get-product) | GET | /products/:id |
| 14 | [Get all products by product category](#get-all-products-by-product-category) | GET | /products/productcategories/:name |
| 15 | [Edit product](#edit-product) | PUT | /products/:id |
| 16 | [Delete product](#delete-product) | DELETE | /products/:id |
| 17 | [Create new cart](#create-new-shopping-cart) | POST | /cart/:productId/:qty |
| 18 | [Add product to cart](#add-product-to-cart) | POST | /carts/:cartId/details |
| 19 | [Get all products on cart](#get-all-products-on-shopping-cart) | GET | /carts/:id |
| 20 | [Delete product from cart](#delete-product-from-shopping-cart) | DELETE | /cartDetails/:carts_id/:products_id |
| 21 | [Delete cart](#delete-shopping-cart) | DELETE | /carts/:id |
| 22 | [Create payment method](#create-payment-method) | POST | /paymentMethods |
| 23 | [Get all payment methods](#get-all-payment-methods) | GET | /paymentMethods |
| 24 | [Get payment method](#get-payment-method) | GET | /paymentMethods/:id |
| 25 | [Edit payment method](#edit-payment-method) | PUT | /paymentMethods/:id |
| 26 | [Delete payment method](#delete-payment-method) | DELETE | /paymentMethods/:id |
| 27 | [Transaction](#transaction) | PUT | /api/transactions/:id |

## Register
### URL : `/customers`
### Method : `POST`

### Body Request
```json
{
	"name" : "nia",
	"gender" : "P",
	"address" : "Jakarta",
	"email" : "nia@gmail.com",
	"password" : "123"
}
```

### Body Response
```json
{
    "message": "success create new user",
    "user": {
<<<<<<< HEAD
<<<<<<< HEAD
        "ID": 0,
=======
>>>>>>> 2958cf9b541a1a4b111722eae10aa7ea148a81fc
        "CreatedAt": "2021-08-03T12:57:09.167+07:00",
        "UpdatedAt": "2021-08-03T12:57:09.167+07:00",
        "DeletedAt": null,
        "id": 3,
        "name": "dewi",
        "address": "Jakarta",
        "gender": "P",
        "email": "dewi@gmail.com",
        "password": "",
=======
        "ID": 5,
        "CreatedAt": "2021-08-04T00:33:33.589+07:00",
        "UpdatedAt": "2021-08-04T00:33:33.589+07:00",
        "DeletedAt": null,
        "name": "nia",
        "address": "Jakarta",
        "gender": "P",
        "email": "nia@gmail.com",
        "password": "123",
>>>>>>> origin/fixing_cart_v2
        "token": "",
        "Carts": null
    }
}
```

## Login
### URL : `/login`
### Method : `POST`

### Body Request
```json
{
	"email" : "dewi@gmail.com",
	"password" : "123"
<<<<<<< HEAD
}
```

### Body Response
```json
{
    "status": "succes login",
    "users": {
        "CreatedAt": "2021-08-03T12:57:09.167+07:00",
        "UpdatedAt": "2021-08-03T13:03:14.935+07:00",
        "DeletedAt": null,
        "id": 3,
        "name": "Dewi",
        "address": "Jakarta",
        "gender": "P",
        "email": "dewi@gmail.com",
        "password": "123",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2Mjc5NzQ5ODcsInVzZXJJZCI6M30.6wbSybF31bNXu3BRlqgN1tWlKZbwS99f1p3atw5fpgw",
        "Carts": null
    }
}
```

## Get Customer
### URL : `/customers/:id`
### Method : `GET`

=======
}
```

### Body Response
```json
<<<<<<< HEAD
{
    "status": "succes login",
    "users": {
        "ID": 3,
        "CreatedAt": "2021-08-03T12:57:09.167+07:00",
        "UpdatedAt": "2021-08-03T22:58:20.691+07:00",
        "DeletedAt": null,
        "name": "Dewi",
        "address": "Jakarta",
        "gender": "P",
        "email": "dewi@gmail.com",
        "password": "123",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MjgwMTU2NTcsInVzZXJJZCI6M30.F-KyjJniT0-VsFcooGmaX9y6VP0PZZ90GEzUioVfsWQ",
        "Carts": null
    }
}
=======
{   
    "message": "success get user",
    "users": {
        "ID": 1,
        "CreatedAt": "2021-07-31T16:01:34.253+07:00",
        "UpdatedAt": "2021-08-03T23:05:05.169+07:00",
        "DeletedAt": null,
        "name": "pikachu",
        "address": "bogor",
        "gender": "M",
        "email": "pikachu@gmail.com",
        "password": "2222",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MjgwMTAzMDUsInVzZXJJZCI6MX0.u6d7zO1VF8eEtx7RmAgeDDSM6Nx5zxTyc20Un2WrObk",
        "Carts": null
    }
} 

>>>>>>> 2958cf9b541a1a4b111722eae10aa7ea148a81fc
```

## Get Customer
### URL : `/customers/:id`
### Method : `GET`

>>>>>>> origin/fixing_cart_v2
### Body Response
```json
{
    "message": "success get all users",
    "users": {
        "ID": 5,
        "CreatedAt": "2021-08-04T00:33:33.589+07:00",
        "UpdatedAt": "2021-08-04T00:33:33.589+07:00",
        "DeletedAt": null,
        "name": "nia",
        "address": "Jakarta",
        "gender": "P",
        "email": "nia@gmail.com",
        "password": "123",
        "token": "",
        "Carts": null
    }
}
```

<<<<<<< HEAD
## Get all products
### URL : `/products`
### Method : `GET`
=======
## Edit customer
### URL : `/customers/:id`
### Method : `PUT`

### Body Request
```json
{
	"name" : "Ka",
	"gender" : "P",
	"address" : "Bogor",
	"email" : "nia@gmail.com",
	"password" : "123"
}
```
>>>>>>> origin/fixing_cart_v2

### Body Response
```json
{
<<<<<<< HEAD
    "message": "success get all products",
    "products": [
        {
            "ID": 1,
            "CreatedAt": "2021-07-31T12:06:53+07:00",
            "UpdatedAt": "2021-07-31T12:06:53+07:00",
            "DeletedAt": null,
            "name": "Maybelline Fit Me 220",
            "code": "MYBLLN_FTM_220",
            "status": "active",
            "price": 75000,
            "description": "Foundation Fit Me shade 220 from Maybelline",
            "Carts": null,
            "ProductCategoriesID": 1
        },
        {
            "ID": 2,
            "CreatedAt": "2021-07-31T13:58:19+07:00",
            "UpdatedAt": "2021-07-31T13:58:19+07:00",
            "DeletedAt": null,
            "name": "Tarte Shape Tape Contour Concealer Fair",
            "code": "TRT_STCC_Fair",
            "status": "active",
            "price": 40000,
            "description": "Shape Tape Contour Concealer shade fair from Tarte",
            "Carts": null,
            "ProductCategoriesID": 2
        },
        {
            "ID": 3,
            "CreatedAt": "2021-08-01T14:30:10+07:00",
            "UpdatedAt": "2021-08-01T14:30:10+07:00",
            "DeletedAt": null,
            "name": "Make Over Ultra Cover Liquid Matt Foundation",
            "code": "MKVR_UCLM",
            "status": "active",
            "price": 200000,
            "description": "Foundation Ultra Cover Liquid Matt from Make Over",
            "Carts": null,
            "ProductCategoriesID": 1
        },
        {
            "ID": 4,
            "CreatedAt": "2021-08-03T12:49:06+07:00",
            "UpdatedAt": "2021-08-03T12:49:06+07:00",
            "DeletedAt": null,
            "name": "Maybelline SuperStay Matte Ink Lipstick",
            "code": "MYBLLN_SMI",
            "status": "active",
            "price": 80000,
            "description": "Lipstick SuperStay Matte Ink from Maybelline",
            "Carts": null,
            "ProductCategoriesID": 5
        },
        {
            "ID": 5,
            "CreatedAt": "2021-08-03T12:50:01+07:00",
            "UpdatedAt": "2021-08-03T12:50:01+07:00",
            "DeletedAt": null,
            "name": "Loreal Lash Paradise",
            "code": "LRL_LP",
            "status": "active",
            "price": 149000,
            "description": "Mascara Lash Paradise from Loreal",
            "Carts": null,
            "ProductCategoriesID": 3
        }
    ]
}
```

## Get all products by product category
### URL : `/products/productcategories/:name`
### Method : `GET`
=======
    "message": "success update customer",
    "update customer": {
        "ID": 5,
        "CreatedAt": "2021-08-04T00:33:33.589+07:00",
        "UpdatedAt": "2021-08-04T00:40:46.572+07:00",
        "DeletedAt": null,
        "name": "Ka",
        "address": "Bogor",
        "gender": "P",
        "email": "nia@gmail.com",
        "password": "123",
        "token": "",
        "Carts": null
    }
}
```

## Delete Customer
### URL : `/customers/:id`
### Method : `DELETE`
>>>>>>> origin/fixing_cart_v2

### Body Response
```json
{
<<<<<<< HEAD
    "message": "success delete customer selected",
    "customer after delete": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": "2021-08-04T00:41:52.333+07:00",
        "name": "",
        "address": "",
        "gender": "",
        "email": "",
        "password": "",
        "token": "",
        "Carts": null
    },
    "customer before delete": {
        "ID": 5,
        "CreatedAt": "2021-08-04T00:33:33.589+07:00",
        "UpdatedAt": "2021-08-04T00:40:46.572+07:00",
        "DeletedAt": null,
        "name": "Ka",
        "address": "Bogor",
        "gender": "P",
        "email": "nia@gmail.com",
        "password": "123",
        "token": "",
        "Carts": null
    }
=======
    "message": "success get all product by product category",
    "products": [
        {
            "ID": 1,
            "CreatedAt": "2021-07-31T12:06:53+07:00",
            "UpdatedAt": "2021-07-31T12:06:53+07:00",
            "DeletedAt": null,
            "name": "Maybelline Fit Me 220",
            "code": "MYBLLN_FTM_220",
            "status": "active",
            "price": 75000,
            "description": "Foundation Fit Me shade 220 from Maybelline",
            "Carts": null,
            "ProductCategoriesID": 1
        }, 
        {
            "ID": 3,
            "CreatedAt": "2021-08-01T14:30:10+07:00",
            "UpdatedAt": "2021-08-01T14:30:10+07:00",
            "DeletedAt": null,
            "name": "Make Over Ultra Cover Liquid Matt Foundation",
            "code": "MKVR_UCLM",
            "status": "active",
            "price": 200000,
            "description": "Foundation Ultra Cover Liquid Matt from Make Over",
            "Carts": null,
            "ProductCategoriesID": 1
        }
    ]
>>>>>>> 2958cf9b541a1a4b111722eae10aa7ea148a81fc
}
```

<<<<<<< HEAD
## Get product
### URL : `/products/:id`
### Method : `GET`

### Body Response
=======
## Create category
### URL : `/productcategories`
### Method : `POST`

### Body Request
>>>>>>> origin/fixing_cart_v2
```json
{
<<<<<<< HEAD
	"name" : "Powder"
=======
    "message": "success get product by id",
    "products": {
            "ID": 1,
            "CreatedAt": "2021-07-31T12:06:53+07:00",
            "UpdatedAt": "2021-07-31T12:06:53+07:00",
            "DeletedAt": null,
            "name": "Maybelline Fit Me 220",
            "code": "MYBLLN_FTM_220",
            "status": "active",
            "price": 75000,
            "description": "Foundation Fit Me shade 220 from Maybelline",
            "Carts": null,
            "ProductCategoriesID": 1
    }    
>>>>>>> 2958cf9b541a1a4b111722eae10aa7ea148a81fc
}
```

<<<<<<< HEAD
## Create new shopping cart
URL : `/cart/:productId/:qty`
Method : `POST`

### Body Request
```json
{
	"payment_methods_id" : 2
}
```

### Body Response
```json
{
    "cart": {
        "ID": 0,
        "CreatedAt": "2021-08-03T13:10:35.03+07:00",
        "UpdatedAt": "2021-08-03T13:10:35.309+07:00",
        "DeletedAt": null,
        "id": 66,
        "status_transactions": "ordered",
        "total_quantity": 2,
        "total_price": 298000,
        "products": null,
        "customers_id": 3,
        "payment_methods_id": 2
    },
    "cartDetails": {
        "products_id": 5,
        "carts_id": 66,
        "quantity": 2,
        "price": 149000,
        "CreatedAt": "2021-08-03T13:10:35.132+07:00",
        "UpdatedAt": "2021-08-03T13:10:35.13+07:00"
    },
    "status": "Create cart success"
}
```

## Add product to cart
### URL : `/carts/:cartId/details`
### Method : `POST`

### Body Request
```json
{
	"products_id" : 1,
	"quantity": 1
}
```

### Body Response
```json
{
    "Total Price": 373000,
    "Total Quantity": 3,
    "cartDetails": {
        "products_id": 1,
        "carts_id": 66,
        "quantity": 1,
        "price": 75000,
        "CreatedAt": "2021-08-03T13:19:43.187+07:00",
        "UpdatedAt": "2021-08-03T13:19:43.183+07:00"
    },
    "status": "Add product to cart success"
}
```

## Get all products on shopping cart
### URL : `/carts/:id`
### Method : `GET`

### Body Response
```json
{
    "cart": {
        "ID": 0,
        "CreatedAt": "2021-08-03T13:10:35.03+07:00",
        "UpdatedAt": "2021-08-03T13:10:35.309+07:00",
        "DeletedAt": null,
        "id": 66,
        "status_transactions": "ordered",
        "total_quantity": 2,
        "total_price": 298000,
        "products": null,
        "customers_id": 3,
        "payment_methods_id": 2
    },
    "products": [
        {
            "ID": 5,
            "CreatedAt": "2021-08-03T12:50:01+07:00",
            "UpdatedAt": "2021-08-03T12:50:01+07:00",
            "DeletedAt": null,
            "name": "Loreal Lash Paradise",
            "code": "LRL_LP",
            "status": "active",
            "price": 149000,
            "description": "Mascara Lash Paradise from Loreal",
            "Carts": null,
=======
### Body Response
```json
{
    "message": "success create product category",
    "user": {
        "CreatedAt": "2021-08-04T00:49:06.687+07:00",
        "DeletedAt": null,
        "UpdatedAt": "2021-08-04T00:49:06.687+07:00",
        "id": 7,
        "name": "Powder"
    }
}
```

## Get All Categories
### URL : `/productcategories`
### Method : `GET`

### Body Response
```json
{
    "message": "success get all product categories",
    "product categories": [
        {
            "ID": 0,
            "CreatedAt": "2021-07-31T12:06:17+07:00",
            "UpdatedAt": "2021-07-31T12:06:17+07:00",
            "DeletedAt": null,
            "id": 1,
            "name": "Foundation",
            "Products": null
        },
        {
            "ID": 0,
            "CreatedAt": "2021-07-31T12:06:27+07:00",
            "UpdatedAt": "2021-07-31T12:06:27+07:00",
            "DeletedAt": null,
            "id": 2,
            "name": "Concealer",
            "Products": null
        },
        {
            "ID": 0,
            "CreatedAt": "2021-08-03T12:40:00+07:00",
            "UpdatedAt": "2021-08-03T12:40:00+07:00",
            "DeletedAt": null,
            "id": 3,
            "name": "Mascara",
            "Products": null
        },
        {
            "ID": 0,
            "CreatedAt": "2021-08-03T12:40:17+07:00",
            "UpdatedAt": "2021-08-03T12:40:17+07:00",
            "DeletedAt": null,
            "id": 4,
            "name": "Blush On",
            "Products": null
        },
        {
            "ID": 0,
            "CreatedAt": "2021-08-03T12:40:39+07:00",
            "UpdatedAt": "2021-08-03T12:40:39+07:00",
            "DeletedAt": null,
            "id": 5,
            "name": "Lipstick",
            "Products": null
        },
        {
            "ID": 0,
            "CreatedAt": "2021-08-04T00:45:46.616+07:00",
            "UpdatedAt": "2021-08-04T00:45:46.616+07:00",
            "DeletedAt": null,
            "id": 6,
            "name": "Eyeliner",
            "Products": null
        },
        {
            "ID": 0,
            "CreatedAt": "2021-08-04T00:49:06.687+07:00",
            "UpdatedAt": "2021-08-04T00:49:06.687+07:00",
            "DeletedAt": null,
            "id": 7,
            "name": "Powder",
            "Products": null
        }
    ]
}
```

## Get Category
### URL : `/productcategories`
### Method : `GET`

### Body Response
```json
{
    "message": "success get product categories by id",
    "product categories": {
        "CreatedAt": "2021-07-31T12:06:27+07:00",
        "DeletedAt": null,
        "UpdatedAt": "2021-07-31T12:06:27+07:00",
        "id": 2,
        "name": "Concealer"
    }
}
```

## Edit category
### URL : `/productcategories/:id`
### Method : `PUT`

### Body Request
```json
{
	"name" : "Compact Powder"
}
```

### Body Response
```json
{
    "message": "success update product categories",
    "update product": {
        "CreatedAt": "2021-08-04T00:49:06.687+07:00",
        "DeletedAt": null,
        "UpdatedAt": "2021-08-04T01:03:16.519+07:00",
        "id": 7,
        "name": "Compact Powder"
    }
}
```

## Delete category
### URL : `/productcategories/:id`
### Method : `DELETE`

### Body Response
```json
{
    "message": "success delete product selected",
    "product after delete": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": "2021-08-04T01:04:29.292+07:00",
        "id": 0,
        "name": "",
        "Products": null
    },
    "product before delete": {
        "ID": 0,
        "CreatedAt": "2021-08-04T00:49:06.687+07:00",
        "UpdatedAt": "2021-08-04T01:03:16.519+07:00",
        "DeletedAt": null,
        "id": 7,
        "name": "Compact Powder",
        "Products": null
    }
}
```

## Create product
### URL : `/products`
### Method : `POST`

### Body Request
```json
{
	"product_categories_id" : 2,
	"name" : "tes",
	"code" : "AAA",
	"status" : "active",
	"price" : 10000,
	"description" : "pppp"
}
```

### Body Response
```json
{
    "message": "success create new product",
    "product add": {
        "ID": 8,
        "CreatedAt": "2021-08-04T01:09:35.903+07:00",
        "UpdatedAt": "2021-08-04T01:09:35.903+07:00",
        "DeletedAt": null,
        "name": "tes",
        "code": "AAA",
        "status": "active",
        "price": 10000,
        "description": "pppp",
        "product_categories_id": 2
    }
}
```

## Get all products
### URL : `/products`
### Method : `GET`

### Body Response
```json
{
    "message": "success get all products",
    "products": [
        {
            "ID": 1,
            "CreatedAt": "2021-07-31T12:06:53+07:00",
            "UpdatedAt": "2021-07-31T12:06:53+07:00",
            "DeletedAt": null,
            "name": "Maybelline Fit Me 220",
            "code": "MYBLLN_FTM_220",
            "status": "active",
            "price": 75000,
            "description": "Foundation Fit Me shade 220 from Maybelline",
            "Carts": null,
            "ProductCategoriesID": 1
        },
        {
            "ID": 2,
            "CreatedAt": "2021-07-31T13:58:19+07:00",
            "UpdatedAt": "2021-07-31T13:58:19+07:00",
            "DeletedAt": null,
            "name": "Tarte Shape Tape Contour Concealer Fair",
            "code": "TRT_STCC_Fair",
            "status": "active",
            "price": 40000,
            "description": "Shape Tape Contour Concealer shade fair from Tarte",
            "Carts": null,
            "ProductCategoriesID": 2
        },
        {
            "ID": 3,
            "CreatedAt": "2021-08-01T14:30:10+07:00",
            "UpdatedAt": "2021-08-01T14:30:10+07:00",
            "DeletedAt": null,
            "name": "Make Over Ultra Cover Liquid Matt Foundation",
            "code": "MKVR_UCLM",
            "status": "active",
            "price": 200000,
            "description": "Foundation Ultra Cover Liquid Matt from Make Over",
            "Carts": null,
            "ProductCategoriesID": 1
        },
        {
            "ID": 4,
            "CreatedAt": "2021-08-03T12:49:06+07:00",
            "UpdatedAt": "2021-08-03T12:49:06+07:00",
            "DeletedAt": null,
            "name": "Maybelline SuperStay Matte Ink Lipstick",
            "code": "MYBLLN_SMI",
            "status": "active",
            "price": 80000,
            "description": "Lipstick SuperStay Matte Ink from Maybelline",
            "Carts": null,
            "ProductCategoriesID": 5
        },
        {
            "ID": 5,
            "CreatedAt": "2021-08-03T12:50:01+07:00",
            "UpdatedAt": "2021-08-03T12:50:01+07:00",
            "DeletedAt": null,
            "name": "Loreal Lash Paradise",
            "code": "LRL_LP",
            "status": "active",
            "price": 149000,
            "description": "Mascara Lash Paradise from Loreal",
            "Carts": null,
>>>>>>> origin/fixing_cart_v2
            "ProductCategoriesID": 3
        }
    ],
    "status": "Success get all products by cart id"
}
```

<<<<<<< HEAD
## Delete product from shopping cart
### URL : `/cartDetails/:products_id/:carts_id`
### Method : `DELETE`
=======
## Get product
### URL : `/products/:id`
### Method : `GET`
>>>>>>> origin/fixing_cart_v2

### Body Response
```json
{
<<<<<<< HEAD
    "Deleted Product": {
        "products_id": 5,
        "carts_id": 66,
        "quantity": 2,
        "price": 149000,
        "CreatedAt": "2021-08-03T13:10:35.132+07:00",
        "UpdatedAt": "2021-08-03T13:10:35.13+07:00"
    },
    "Total Price": 75000,
    "Total Quantity": 1,
    "status": "Delete product on table cart_details success"
}
```

## Delete shopping cart
### URL : `/cartDetails/:products_id/:carts_id`
### Method : `DELETE`

### Body Response
```json
{
    "Deleted Cart": {
        "ID": 0,
        "CreatedAt": "2021-08-03T13:10:35.03+07:00",
        "UpdatedAt": "2021-08-03T13:23:39.206+07:00",
        "DeletedAt": null,
        "id": 66,
        "status_transactions": "ordered",
        "total_quantity": 1,
        "total_price": 75000,
        "products": null,
        "customers_id": 3,
        "payment_methods_id": 2
    },
    "status": "Delete cart success"
}
```

## Create transaction and payment
### URL : `/transactions`
### Method : `POST`

### Body Request
```json
{

}
```

=======
    "message": "success get product by id",
    "products": {
        "ID": 2,
        "CreatedAt": "2021-07-31T13:58:19+07:00",
        "UpdatedAt": "2021-07-31T13:58:19+07:00",
        "DeletedAt": null,
        "name": "Tarte Shape Tape Contour Concealer Fair",
        "code": "TRT_STCC_Fair",
        "status": "active",
        "price": 40000,
        "description": "Shape Tape Contour Concealer shade fair from Tarte",
        "product_categories_id": 2
    }
}
```

## Get all products by product category
### URL : `/products/productcategories/:name`
### Method : `GET`

### Body Response
```json
{
    "message": "success get all product by product category",
    "product by category": [
        {
            "ID": 4,
            "CreatedAt": "2021-08-03T12:49:06+07:00",
            "UpdatedAt": "2021-08-03T12:49:06+07:00",
            "DeletedAt": null,
            "name": "Maybelline SuperStay Matte Ink Lipstick",
            "code": "MYBLLN_SMI",
            "status": "active",
            "price": 80000,
            "description": "Lipstick SuperStay Matte Ink from Maybelline",
            "product_categories_id": 5
        }
    ]
}
```

## Edit product
### URL : `/products/:id`
### Method : `PUT`

### Body Request
```json
{
	"product_categories_id" : 2,
	"name" : "Mascara Hypersharp",
	"code" : "AAA",
	"status" : "active",
	"price" : 10000,
	"description" : "pppp"
}
```

### Body Response
```json
{
    "message": "success update product",
    "update product": {
        "ID": 8,
        "CreatedAt": "2021-08-04T01:09:35.903+07:00",
        "UpdatedAt": "2021-08-04T01:22:25.288+07:00",
        "DeletedAt": null,
        "name": "Mascara Hypersharp",
        "code": "AAA",
        "status": "active",
        "price": 10000,
        "description": "pppp",
        "product_categories_id": 2
    }
}
```

## Delete product
### URL : `/products/:id`
### Method : `DELETE`

### Body Response
```json
{
    "message": "success delete product selected",
    "product after delete": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": "2021-08-04T01:24:51.509+07:00",
        "name": "",
        "code": "",
        "status": "",
        "price": 0,
        "description": "",
        "product_categories_id": 0
    },
    "product before delete": {
        "ID": 8,
        "CreatedAt": "2021-08-04T01:09:35.903+07:00",
        "UpdatedAt": "2021-08-04T01:22:25.288+07:00",
        "DeletedAt": null,
        "name": "Mascara Hypersharp",
        "code": "AAA",
        "status": "active",
        "price": 10000,
        "description": "pppp",
        "product_categories_id": 2
    }
}
```

## Create new cart
URL : `/cart/:productId/:qty`
Method : `POST`

### Body Request
```json
{
	"payment_methods_id" : 2
}
```

### Body Response
```json
{
    "cart": {
        "CreatedAt": "2021-08-03T23:34:03.685+07:00",
        "DeletedAt": null,
        "ID": 94,
        "UpdatedAt": "2021-08-03T23:34:03.839+07:00",
        "customers_id": 3,
        "payment_methods_id": 2,
        "status_transactions": "ordered",
        "total_price": 298000,
        "total_quantity": 2
    },
    "cartDetails": {
        "products_id": 5,
        "carts_id": 94,
        "quantity": 2,
        "price": 149000,
        "CreatedAt": "2021-08-03T23:34:03.736+07:00",
        "UpdatedAt": "2021-08-03T23:34:03.733+07:00"
    },
    "status": "Create cart success"
}
```

## Add product to cart
### URL : `/carts/:cartId/details`
### Method : `POST`

### Body Request
```json
{
	"products_id" : 1,
	"quantity": 1
}
```

### Body Response
```json
{
    "Total Price": 373000,
    "Total Quantity": 3,
    "cartDetails": {
        "products_id": 1,
        "carts_id": 94,
        "quantity": 1,
        "price": 75000,
        "CreatedAt": "2021-08-03T23:35:15.102+07:00",
        "UpdatedAt": "2021-08-03T23:35:15.1+07:00"
    },
    "status": "Add product to cart success"
}
```

## Get all products on cart
### URL : `/carts/:id`
### Method : `GET`

### Body Response
```json
{
    "cart": {
        "CreatedAt": "2021-08-03T23:34:03.685+07:00",
        "DeletedAt": null,
        "ID": 94,
        "UpdatedAt": "2021-08-03T23:34:03.839+07:00",
        "customers_id": 3,
        "payment_methods_id": 2,
        "status_transactions": "ordered",
        "total_price": 298000,
        "total_quantity": 2
    },
    "products": [
        {
            "ID": 5,
            "CreatedAt": "2021-08-03T12:50:01+07:00",
            "UpdatedAt": "2021-08-03T12:50:01+07:00",
            "DeletedAt": null,
            "name": "Loreal Lash Paradise",
            "code": "LRL_LP",
            "status": "active",
            "price": 149000,
            "description": "Mascara Lash Paradise from Loreal",
            "ProductCategoriesID": 3
        }
    ],
    "status": "Success get all products by cart id"
}
```

## Delete product from cart
### URL : `/cartDetails/:products_id/:carts_id`
### Method : `DELETE`

### Body Response
```json
{
    "Deleted Product": {
        "products_id": 1,
        "carts_id": 94,
        "quantity": 1,
        "price": 75000,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "2021-08-03T23:37:33.589+07:00"
    },
    "Total Price": 298000,
    "Total Quantity": 2,
    "status": "Delete product on table cart_details success"
}
```

## Delete cart
### URL : `/cartDetails/:products_id/:carts_id`
### Method : `DELETE`

### Body Response
```json
{
    "Deleted Cart": {
        "CreatedAt": "2021-08-03T23:34:03.685+07:00",
        "DeletedAt": null,
        "ID": 94,
        "UpdatedAt": "2021-08-03T23:37:49.829+07:00",
        "customers_id": 3,
        "payment_methods_id": 2,
        "status_transactions": "ordered",
        "total_price": 298000,
        "total_quantity": 2
    },
    "status": "Delete cart success"
}
```

## Create payment method
### URL : `/paymentMethods`
### Method : `POST`

### Body Request
```json
{
	"name" : "Gopay",
	"status" : "active"
}
```

### Body Response
```json
{
    "data": {
        "CreatedAt": "2021-08-04T01:31:49.871+07:00",
        "DeletedAt": null,
        "UpdatedAt": "2021-08-04T01:31:49.871+07:00",
        "id": 5,
        "name": "Gopay",
        "status": "active"
    },
    "messages": "success create payment methods"
}
```

## Get all payment methods
### URL : `/paymentMethods`
### Method : `GET`

### Body Response
```json
{
    "data": [
        {
            "ID": 0,
            "CreatedAt": "2021-07-31T12:07:08+07:00",
            "UpdatedAt": "2021-07-31T12:07:08+07:00",
            "DeletedAt": null,
            "id": 1,
            "name": "ovo",
            "status": "active",
            "Carts": null
        },
        {
            "ID": 0,
            "CreatedAt": "2021-08-03T08:29:44+07:00",
            "UpdatedAt": "2021-08-03T08:29:44+07:00",
            "DeletedAt": null,
            "id": 2,
            "name": "DANA",
            "status": "active",
            "Carts": null
        },
        {
            "ID": 0,
            "CreatedAt": "2021-08-04T01:31:49.871+07:00",
            "UpdatedAt": "2021-08-04T01:31:49.871+07:00",
            "DeletedAt": null,
            "id": 5,
            "name": "Gopay",
            "status": "active",
            "Carts": null
        }
    ],
    "status": "success get all payment method"
}
```

## Get payment method
### URL : `/paymentMethods/:id`
### Method : `GET`

### Body Response
```json
{
    "data": {
        "CreatedAt": "2021-08-03T08:29:44+07:00",
        "DeletedAt": null,
        "UpdatedAt": "2021-08-03T08:29:44+07:00",
        "id": 2,
        "name": "DANA",
        "status": "active"
    },
    "message": "success get payment method data"
}
```

## Edit payment method
### URL : `/paymentMethods/:id`
### Method : `PUT`

### Body Request
```json
{
	"name" : "Credit",
	"status" : "active"
}
```

>>>>>>> origin/fixing_cart_v2
### Body Response
```json
{
    "data": {
        "CreatedAt": "2021-08-04T01:31:49.871+07:00",
        "DeletedAt": null,
        "UpdatedAt": "2021-08-04T01:37:37.794+07:00",
        "id": 5,
        "name": "Credit",
        "status": "active"
    },
    "message": "success updating payment method"
}
```

## Delete payment method
### URL : `/paymentMethods/:id`
### Method : `DELETE`

### Body Response
```json
{
    "data": [],
    "message": "success delete payment method"
}
```

## Transaction
### URL : `/api/transactions/:id`
### Method : `PUT`

### Body Response
```json
{
    "data": {
        "CreatedAt": "2021-07-31T17:54:44+07:00",
        "DeletedAt": null,
        "ID": 2,
        "UpdatedAt": "2021-08-04T02:11:15.939+07:00",
        "customers_id": 1,
        "payment_methods_id": 1,
        "status_transactions": "paid",
        "total_price": 375000,
        "total_quantity": 5
    },
    "message": "success updating transaction status"
}
```