// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "Login",
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                }
            }
        },
        "/auth/refresh/token": {
            "post": {
                "description": "RefreshToken",
                "tags": [
                    "Auth"
                ],
                "summary": "RefreshToken",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.TokenReqBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Register",
                "tags": [
                    "Auth"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                }
            }
        },
        "/orders": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "CreateOrder",
                "tags": [
                    "Order"
                ],
                "summary": "CreateOrder",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.OrderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                }
            }
        },
        "/products": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get List Products",
                "tags": [
                    "Product"
                ],
                "summary": "Get List Products",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "int valid",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "int valid",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "string valid",
                        "name": "sortField",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "string valid",
                        "name": "sortValue",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ProductPagination"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Create Product",
                "tags": [
                    "Product"
                ],
                "summary": "Create Product",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CreateDataResponse"
                        }
                    }
                }
            }
        },
        "/products/categories": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get List Product Category",
                "tags": [
                    "Product Category"
                ],
                "summary": "Get List product category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "int valid",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "int valid",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "string valid",
                        "name": "sortField",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "string valid",
                        "name": "sortValue",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ProductCategoryPagination"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Create Product Category",
                "tags": [
                    "Product Category"
                ],
                "summary": "Create product category",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ProductCategoryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CreateDataResponse"
                        }
                    }
                }
            }
        },
        "/products/categories/{productCategoryID}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get product category By Id",
                "tags": [
                    "Product Category"
                ],
                "summary": "Get product category By Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "productCategoryID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ProductCategory"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Update product category By Id",
                "tags": [
                    "Product Category"
                ],
                "summary": "Update product category By Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "productCategoryID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateProductCategory"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.UpdateDataResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Delete product category By Id",
                "tags": [
                    "Product Category"
                ],
                "summary": "Delete product category By Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "productCategoryID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.UpdateDataResponse"
                        }
                    }
                }
            }
        },
        "/products/{productID}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get product By Id",
                "tags": [
                    "Product"
                ],
                "summary": "Get product By Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "productID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Product"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Update product By Id",
                "tags": [
                    "Product"
                ],
                "summary": "Update product By Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "productID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateProduct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.UpdateDataResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Delete product By Id",
                "tags": [
                    "Product"
                ],
                "summary": "Delete product By Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "productID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.DeleteDataResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get List Users",
                "tags": [
                    "User"
                ],
                "summary": "Get List Users",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "int valid",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "int valid",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "string valid",
                        "name": "sortField",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "string valid",
                        "name": "sortValue",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.UserPagination"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Create user",
                "tags": [
                    "User"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CreateDataResponse"
                        }
                    }
                }
            }
        },
        "/users/{userID}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get user By Id",
                "tags": [
                    "User"
                ],
                "summary": "Get user By Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.User"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Update user By Id",
                "tags": [
                    "User"
                ],
                "summary": "Update user By Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.UpdateDataResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Delete user By Id",
                "tags": [
                    "User"
                ],
                "summary": "Delete user By Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.DeleteDataResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "model.User": {
            "type": "object",
            "required": [
                "fullname",
                "password",
                "username"
            ],
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "fullname": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "request.LoginRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "example": "1234"
                },
                "username": {
                    "type": "string",
                    "example": "admin"
                }
            }
        },
        "request.OrderItem": {
            "type": "object",
            "required": [
                "amount",
                "id",
                "name",
                "price",
                "product_category_id"
            ],
            "properties": {
                "amount": {
                    "type": "integer",
                    "example": 10
                },
                "description": {
                    "type": "string",
                    "example": "apple"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "apple"
                },
                "price": {
                    "type": "number",
                    "example": 200
                },
                "product_category_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "request.OrderRequest": {
            "type": "object",
            "required": [
                "orders"
            ],
            "properties": {
                "orders": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/request.OrderItem"
                    }
                }
            }
        },
        "request.Product": {
            "type": "object",
            "required": [
                "amount",
                "name",
                "price",
                "product_category_id",
                "sale_close_date",
                "sale_open_date"
            ],
            "properties": {
                "amount": {
                    "type": "integer",
                    "example": 10
                },
                "description": {
                    "type": "string",
                    "example": "apple"
                },
                "name": {
                    "type": "string",
                    "example": "apple"
                },
                "price": {
                    "type": "number",
                    "example": 200
                },
                "product_category_id": {
                    "type": "integer",
                    "example": 1
                },
                "sale_close_date": {
                    "type": "string",
                    "example": "2021-12-26T00:00:00Z"
                },
                "sale_open_date": {
                    "type": "string",
                    "example": "2021-12-26T00:00:00Z"
                }
            }
        },
        "request.ProductCategoryRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "request.TokenReqBody": {
            "type": "object",
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "request.UpdateProduct": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer",
                    "example": 10
                },
                "description": {
                    "type": "string",
                    "example": "apple"
                },
                "name": {
                    "type": "string",
                    "example": "apple"
                },
                "price": {
                    "type": "number",
                    "example": 200
                },
                "product_category_id": {
                    "type": "integer",
                    "example": 1
                },
                "sale_close_date": {
                    "type": "string",
                    "example": "2021-12-26T00:00:00Z"
                },
                "sale_open_date": {
                    "type": "string",
                    "example": "2021-12-26T00:00:00Z"
                }
            }
        },
        "request.UpdateProductCategory": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "request.UserRequest": {
            "type": "object",
            "required": [
                "fullname",
                "password",
                "username"
            ],
            "properties": {
                "avatar": {
                    "type": "string",
                    "example": "admin"
                },
                "fullname": {
                    "type": "string",
                    "example": "admin test"
                },
                "password": {
                    "type": "string",
                    "example": "1234"
                },
                "username": {
                    "type": "string",
                    "example": "admin"
                }
            }
        },
        "response.CreateDataResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "create success"
                },
                "response_key": {
                    "type": "string"
                },
                "response_message": {
                    "type": "string"
                }
            }
        },
        "response.DeleteDataResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "delete success"
                },
                "response_key": {
                    "type": "string"
                },
                "response_message": {
                    "type": "string"
                }
            }
        },
        "response.Product": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "example": "apple"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string",
                    "example": "apple"
                }
            }
        },
        "response.ProductCategory": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "response.ProductCategoryPagination": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.ProductCategory"
                    }
                },
                "page": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "response_key": {
                    "type": "string"
                },
                "response_message": {
                    "type": "string"
                },
                "totalPage": {
                    "type": "integer"
                }
            }
        },
        "response.ProductPagination": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.Product"
                    }
                },
                "page": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "response_key": {
                    "type": "string"
                },
                "response_message": {
                    "type": "string"
                },
                "totalPage": {
                    "type": "integer"
                }
            }
        },
        "response.UpdateDataResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "update success"
                },
                "response_key": {
                    "type": "string"
                },
                "response_message": {
                    "type": "string"
                }
            }
        },
        "response.User": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "fullName": {
                    "description": "Password string ` + "`" + `json:\"password\"` + "`" + `",
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "response.UserPagination": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.User"
                    }
                },
                "page": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "response_key": {
                    "type": "string"
                },
                "response_message": {
                    "type": "string"
                },
                "totalPage": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Type \"Bearer\" followed by a space and JWT token.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
