// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/admin/bill": {
            "get": {
                "description": "Get all billboards from data base",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin/billboard"
                ],
                "summary": "GetAll",
                "operationId": "get-billboards",
                "responses": {}
            },
            "post": {
                "description": "Create the billboard and add it to data base",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin/billboard"
                ],
                "summary": "Create",
                "operationId": "create-billboard",
                "parameters": [
                    {
                        "description": " height / width / display_type / location_id / price",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/admin/bill/{id}": {
            "get": {
                "description": "Get the billboard from data base",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin/billboard"
                ],
                "summary": "GetById",
                "operationId": "get-billboard",
                "responses": {}
            },
            "put": {
                "description": "Update",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin/billboard"
                ],
                "summary": "UpdateById",
                "operationId": "update-billboard",
                "responses": {}
            }
        },
        "/auth/sign-in": {
            "post": {
                "description": "login to account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Sign In",
                "operationId": "login-account",
                "parameters": [
                    {
                        "description": "username / password",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.SignInRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/auth/sign-up": {
            "post": {
                "description": "register to site",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Sign Up",
                "operationId": "create-account",
                "parameters": [
                    {
                        "description": "username / email / password / name / surname",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "dtos.SignInRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.Product": {
            "type": "object",
            "properties": {
                "display_type": {
                    "type": "integer"
                },
                "height": {
                    "type": "integer"
                },
                "locationId": {
                    "type": "integer"
                },
                "price": {
                    "type": "integer"
                },
                "productId": {
                    "type": "integer"
                },
                "width": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password",
                "surname",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "headers"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "HappyBill",
	Description:      "This is a server for order billboards.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
