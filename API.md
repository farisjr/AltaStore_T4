# Application Programming Interface Documentation

## List
| No | Web Service | Method | URL |
|----|-------------|--------|-----|
| 1 | [Register](#register) | POST | /customers |
| 2 | [Login](#login) | POST | /login |
| 3 | [Get Customer](#get-customer) | GET | /customers/:id |
| 4 | [Get all products](#get-all-products) | GET | /products |
| 5 | [Get all products by product category](#get-all-products-by-product-category) | GET | /products/productcategories/:name |
| 6 | [Get product](#get-product) | GET | /products |
| 7 | [Create new shopping cart](#create-new-shopping-cart) | POST | /cart/:productId/:qty |
| 8 | [Add product to cart](#add-product-to-cart) | POST | /carts/:cartId/details |
| 9 | [Get all products on shopping cart](#get-all-products-on-shopping-cart) | GET | /carts/:id |
| 10 | [Delete product from shopping cart](#delete-product-from-shopping-cart) | DELETE | /cartDetails/:products_id/:carts_id |
| 11 | [Delete shopping cart](#delete-shopping-cart) | DELETE | /carts/:id |
| 12 | [Create transaction and payment](#create-transaction-and-payment) | POST | /transactions |

## Register
### URL : `/customers`
### Method : `POST`

### Body Request
```json
{
	"name" : "dewi",
	"gender" : "P",
	"address" : "Jakarta",
	"email" : "dewi@gmail.com",
	"password" : 123
}
```

### Body Response
```json
{
    "message": "success create new user",
    "user": {
        "ID": 0,
        "CreatedAt": "2021-08-03T12:57:09.167+07:00",
        "UpdatedAt": "2021-08-03T12:57:09.167+07:00",
        "DeletedAt": null,
        "id": 3,
        "name": "dewi",
        "address": "Jakarta",
        "gender": "P",
        "email": "dewi@gmail.com",
        "password": "",
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
}
```

### Body Response
```json
{
    "status": "succes login",
    "users": {
        "ID": 0,
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

### Body Response
```json
{
    
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
            "ProductCategoriesID": 3
        }
    ]
}
```

## Get all products by product category
### URL : `/products/productcategories/:name`
### Method : `GET`

### Body Response
```json
{
    
}
```

## Get product
### URL : `/products`
### Method : `GET`

### Body Response
```json
{
    
}
```

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
            "ProductCategoriesID": 3
        }
    ],
    "status": "Success get all products by cart id"
}
```

## Delete product from shopping cart
### URL : `/cartDetails/:products_id/:carts_id`
### Method : `DELETE`

### Body Response
```json
{
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

### Body Response
```json
{
    
}
```