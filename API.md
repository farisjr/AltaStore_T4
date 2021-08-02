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

}
```

### Login
URL : `/api/login`
Method : `POST`

Response
```json
{
    
}
```

### Get Customer
URL : `/api/customers/:id`
Method : `GET`

Response
```json
{
    
}
```

### Get all products
URL : `/api/products`
Method : `GET`

Response
```json
{
    
}
```

### Get all products by product category
URL : `/api/products/productcategories/:name`
Method : `GET`

Response
```json
{
    
}
```

### Get product
URL : `/api/products`
Method : `GET`

Response
```json
{
    
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