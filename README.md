# AltaStore Project

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)
[![Go.Dev reference](https://img.shields.io/badge/gorm-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/gorm.io/gorm?tab=doc)
[![Go.Dev reference](https://img.shields.io/badge/echo-reference-blue?logo=go&logoColor=white)](https://github.com/labstack/echo)

## Table of Content

  - [Team Member](#team-member)
  - [About](#about)
    - [Features](#features)
    - [API Documentation](#api-documentation)
  - [Getting started](#getting-started)
    - [Installing](#installing)
    - [Layout](#layout)

## Team Member

- [Faris Jourdy Ramadhan](https://github.com/farisjr)
- [Mohammad Ryan Priatama](https://github.com/ryanpriatama)
- [Riska Kurnia Dewi](https://github.com/riskakrndw)

## About

This project is example of E-Commerce application providing REST API

### Features

- [x] Register
- [x] Login
- [x] Get customer
- [x] Edit customer
- [x] Delete customer
- [x] Create category
- [x] Get all categories
- [x] Get category
- [x] Edit category
- [x] Delete category
- [x] Create product
- [x] Get all products
- [x] Get product
- [x] Get all products by product category
- [x] Edit product
- [x] Delete product
- [x] Create new cart
- [x] Add product to cart
- [x] Get all products on cart
- [x] Delete product from cart
- [x] Delete cart
- [x] Create payment method
- [x] Get all payment methods
- [x] Get payment method
- [x] Edit payment method
- [x] Delete payment method
- [x] Transaction

### API Documentation

Our Application Programming Interface is available at [API page.](API.md)

## Getting Started

Below we describe how to start this project

### Installing

You must download and install `golang`, follow [this instruction](https://golang.org/doc/install) to install.

After Golang installed, Follow this instructions
```bash
$ git clone https://github.com/farisjr/AltaStore_T4.git
$ go run main.go
```

Go to `http://localhost:8000/` to [start this application.](http://localhost:8000/)

### Layout

```tree
├── config
│   ├── config.go
├── constants
│   ├── constants.go
├── controllers
│   ├── cartControllers.go
│   ├── cartDetailControllers.go
│   ├── customerControllers.go
│   ├── paymentMethodsControllers.go
│   ├── productControllers.go
│   ├── productTypeControllers.go
├── lib
│   ├── database
│   │   └── cartDetails.go
│   │   └── carts.go
│   │   └── customers.go
│   │   └── paymentMethods.go
│   │   └── productDescriptions.go
│   │   └── products.go
│   │   └── productTypes.go
├── middlewares
│   ├── jwtMiddlewares.go
├── models
│   │── cartDetails.go
│   │── carts.go
│   │── customers.go
│   │── paymentMethods.go
│   │── productCategories.go
│   │── products.go
│   │── productTypes.go
├── routes
│   ├── routes.go
├── .gitignore
├── API.md
├── go.mod
├── go.sum
├── main.go
├── README.md
```