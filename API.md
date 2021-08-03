# Application Programming Interface Documentation

## List
| No | Web Service | Method | URL |
|----|-------------|--------|-----|
| 1 | [Register](#register) | POST | /api/customers |
| 2 | [Login](#login) | POST | /api/login |
| 3 | [Get Customer](#get-customer) | GET | /api/customers/:id |
| 4 | [Get all products](#get-all-products) | GET | /api/products |
| 5 | [Get all products by product category](#get-all-products-by-product-category) | GET | /api/products/productcategories/:name |
| 6 | [Get product](#get-product) | GET | /api/products |
| 7 | [Create new shopping cart](#create-new-shopping-cart) | POST | /api/cart/:productId/:qty |
| 8 | [Add product to cart](#add-product-to-cart) | POST | api/carts/:cartId/details |
| 9 | [Get all products on shopping cart](#get-all-products-on-shopping-cart) | GET | /api/carts/:id |
| 10 | [Delete product from shopping cart](#delete-product-from-shopping-cart) | DELETE | /api/cartDetails/:products_id/:carts_id |
| 11 | [Create transaction and payment](#create-transaction-and-payment) | POST | /api/transactions |

### Register
URL : `/api/customers`
Method : `POST`

Response
```json
{
    "message": "success create new user",
    "user": {
        "ID": 1,
        "CreatedAt": "2021-08-03T20:41:27.538+07:00",
        "UpdatedAt": "2021-08-03T20:41:27.538+07:00",
        "DeletedAt": null,
        "name": "pikachu",
        "address": "bogor",
        "gender": "M",
        "email": "pikachu@gmail.com",
        "password": "2222",
        "token": "",
        "Carts": null
    }
}
```

### Login
URL : `/api/login`
Method : `POST`

Response
```json
{
    "status": "succes login",
    "users": {
        "ID": 1,
        "CreatedAt": "2021-07-31T16:01:34.253+07:00",
        "UpdatedAt": "2021-08-02T21:40:18.26+07:00",
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
```

### Get Customer
URL : `/api/customers/:id`
Method : `GET`

Response
```json
{
{
    "message": "success get all users",
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
}
```

### Get all products
URL : `/api/products`
Method : `GET`

Response
```json
{
    "message": "success get all products",
    "products": [
        {
            "ID": 1,
            "CreatedAt": "2021-07-31T13:33:25.558+07:00",
            "UpdatedAt": "2021-07-31T13:33:25.558+07:00",
            "DeletedAt": null,
            "name": "iphone 12",
            "code": "e01",
            "status": "available",
            "price": 20000000,
            "description": "",
            "ProductCategoriesID": 0
        },
        {
            "ID": 13,
            "CreatedAt": "2021-07-31T14:32:10.067+07:00",
            "UpdatedAt": "2021-07-31T14:32:10.067+07:00",
            "DeletedAt": null,
            "name": "asus max pro",
            "code": "e02",
            "status": "available",
            "price": 4000000,
            "description": "",
            "ProductCategoriesID": 0
        },
        {
            "ID": 14,
            "CreatedAt": "2021-07-31T14:36:39.687+07:00",
            "UpdatedAt": "2021-07-31T14:36:39.687+07:00",
            "DeletedAt": null,
            "name": "laptop asus",
            "code": "e03",
            "status": "available",
            "price": 3000000,
            "description": "",
            "ProductCategoriesID": 0
        },
        {
            "ID": 15,
            "CreatedAt": "2021-07-31T14:37:55.48+07:00",
            "UpdatedAt": "2021-08-01T21:10:48.075+07:00",
            "DeletedAt": null,
            "name": "mie ayam",
            "code": "m01",
            "status": "available",
            "price": 20000,
            "description": "",
            "ProductCategoriesID": 0
        },
        {
            "ID": 16,
            "CreatedAt": "2021-07-31T14:38:31.671+07:00",
            "UpdatedAt": "2021-07-31T14:38:31.671+07:00",
            "DeletedAt": null,
            "name": "bakso",
            "code": "m02",
            "status": "available",
            "price": 30000,
            "description": "",
            "ProductCategoriesID": 0
        },
        {
            "ID": 17,
            "CreatedAt": "2021-07-31T14:39:07.993+07:00",
            "UpdatedAt": "2021-07-31T14:39:07.993+07:00",
            "DeletedAt": null,
            "name": "kemeja",
            "code": "p01",
            "status": "available",
            "price": 100000,
            "description": "",
            "ProductCategoriesID": 0
        },
        {
            "ID": 18,
            "CreatedAt": "2021-07-31T14:39:40.061+07:00",
            "UpdatedAt": "2021-08-01T21:54:48.745+07:00",
            "DeletedAt": null,
            "name": "gamis",
            "code": "p02",
            "status": "available",
            "price": 120000,
            "description": "",
            "ProductCategoriesID": 0
        }
    ]
}
```

### Get all products by product category
URL : `/api/products/productcategories/:name`
Method : `GET`

Response
```json
{
    "message": "success get all product by product category",
    "product by category": [
        {
            "ID": 17,
            "CreatedAt": "2021-07-31T14:39:07.993+07:00",
            "UpdatedAt": "2021-07-31T14:39:07.993+07:00",
            "DeletedAt": null,
            "name": "kemeja",
            "code": "p01",
            "status": "available",
            "price": 100000,
            "description": "kemeja biru",
            "ProductCategoriesID": 2
        },
        {
            "ID": 18,
            "CreatedAt": "2021-07-31T14:39:40.061+07:00",
            "UpdatedAt": "2021-08-01T21:54:48.745+07:00",
            "DeletedAt": null,
            "name": "gamis",
            "code": "p02",
            "status": "available",
            "price": 120000,
            "description": "gamis putih",
            "ProductCategoriesID": 2
        }
    ]
}
```

### Get product
URL : `/api/products/:id`
Method : `GET`

Response
```json
{
    "message": "success get product by id",
    "products": {
        "ID": 18,
        "CreatedAt": "2021-07-31T14:39:40.061+07:00",
        "UpdatedAt": "2021-08-01T21:54:48.745+07:00",
        "DeletedAt": null,
        "name": "gamis",
        "code": "p02",
        "status": "available",
        "price": 120000,
        "description": "",
        "ProductCategoriesID": 0
    }
}
```

### Create new shopping cart
URL : `/api/cart/:productId/:qty`
Method : `POST`

Response
```json
{
    "cart": {
        "ID": 0,
        "CreatedAt": "2021-08-02T00:28:25.5+07:00",
        "UpdatedAt": "2021-08-02T00:28:25.858+07:00",
        "DeletedAt": null,
        "id": 18,
        "status_transactions": "ordered",
        "total_quantity": 3,
        "total_price": 120000,
        "Products": null,
        "CustomersID": 1,
        "PaymentMethodsID": 1
    },
    "cartDetails": {
        "products_id": 2,
        "carts_id": 18,
        "quantity": 3,
        "price": 40000,
        "CreatedAt": "2021-08-02T00:28:25.578+07:00",
        "UpdatedAt": "2021-08-02T00:28:25.576+07:00"
    },
    "status": "create cart success"
}
```

### Add product to cart
URL : `api/carts/:cartId/details`
Method : `POST`

Response
```json
{
    "cartDetails": {
        "products_id": 3,
        "carts_id": 2,
        "quantity": 2,
        "price": 200000,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "2021-08-02T00:26:12.285+07:00"
    },
    "status": "add product to cart success"
}
```

### Get all products on shopping cart
URL : `/api/carts/:id`
Method : `GET`

Response
```json
{
    "cart": {
        "ID": 0,
        "CreatedAt": "2021-07-31T12:09:32+07:00",
        "UpdatedAt": "2021-07-31T22:41:36.274+07:00",
        "DeletedAt": null,
        "id": 1,
        "status_transactions": "ordered",
        "total_quantity": 3,
        "total_price": 190000,
        "Products": null,
        "CustomersID": 1,
        "PaymentMethodsID": 1
    },
    "message": "success get all products by cart id",
    "products": [
        {
            "ID": 0,
            "CreatedAt": "2021-07-31T12:06:53+07:00",
            "UpdatedAt": "2021-07-31T12:06:53+07:00",
            "DeletedAt": null,
            "id": 1,
            "name": "Baju A",
            "code": "AAA",
            "status": "active",
            "price": 75000,
            "description": "baju a",
            "Carts": null,
            "ProductCategoriesID": 1
        },
        {
            "ID": 0,
            "CreatedAt": "2021-07-31T13:58:19+07:00",
            "UpdatedAt": "2021-07-31T13:58:19+07:00",
            "DeletedAt": null,
            "id": 2,
            "name": "Celana B",
            "code": "BBBBBBBBS",
            "status": "active",
            "price": 40000,
            "description": "zzzzzzz",
            "Carts": null,
            "ProductCategoriesID": 2
        }
    ]
}
```

### Delete product from shopping cart
URL : `/api/cartDetails/:products_id/:carts_id`
Method : `DELETE`

Response
```json
{
    "message": "delete product on table cart_details success",
    "product": {
        "products_id": 3,
        "carts_id": 2,
        "quantity": 2,
        "price": 200000,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "2021-08-02T00:26:12.285+07:00"
    }
}
```

### Create transaction and payment
URL : `/api/transactions`
Method : `POST`

Response
```json
{
    
}
```
