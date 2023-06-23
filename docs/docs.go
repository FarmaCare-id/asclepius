// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/auth/credential": {
            "get": {
                "description": "Put all mandatory parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Get User Credential",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/auth/forgot-password": {
            "post": {
                "description": "Put all mandatory parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Forgot Password",
                "parameters": [
                    {
                        "description": "ForgotPasswordRequest",
                        "name": "ForgotPasswordRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ForgotPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ForgotPasswordRequest"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Put all mandatory parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login User Google",
                "parameters": [
                    {
                        "description": "LoginRequest",
                        "name": "LoginRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.LoginResponse"
                        }
                    }
                }
            }
        },
        "/auth/login-google": {
            "post": {
                "description": "Put all mandatory parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login User",
                "parameters": [
                    {
                        "description": "GoogleLoginRequest",
                        "name": "GoogleLoginRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.GoogleLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.LoginResponse"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "post": {
                "description": "Put all mandatory parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Logout User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Put all mandatory parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register New User",
                "parameters": [
                    {
                        "description": "CreateUserRequest",
                        "name": "CreateUserRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateUserResponse"
                        }
                    }
                }
            }
        },
        "/auth/register/admin": {
            "post": {
                "description": "Put all mandatory parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register New Admin",
                "parameters": [
                    {
                        "description": "CreateAdminRequest",
                        "name": "CreateAdminRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateAdminRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateAdminResponse"
                        }
                    }
                }
            }
        },
        "/auth/register/doctor": {
            "post": {
                "description": "Put all mandatory parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register New Doctor",
                "parameters": [
                    {
                        "description": "CreateDoctorRequest",
                        "name": "CreateDoctorRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateDoctorRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateDoctorResponse"
                        }
                    }
                }
            }
        },
        "/auth/register/pharmacist": {
            "post": {
                "description": "Put all mandatory parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register New Pharmacist",
                "parameters": [
                    {
                        "description": "CreatePharmacistRequest",
                        "name": "CreatePharmacistRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreatePharmacistRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CreatePharmacistResponse"
                        }
                    }
                }
            }
        },
        "/auth/reset-password": {
            "post": {
                "description": "Put all mandatory parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Reset Password",
                "parameters": [
                    {
                        "description": "ResetPasswordRequest",
                        "name": "ResetPasswordRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ResetPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ResetPasswordRequest"
                        }
                    }
                }
            }
        },
        "/feedback": {
            "get": {
                "description": "Put all mandatory parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Feedback"
                ],
                "summary": "Get All Feedback",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/feedback/category": {
            "get": {
                "description": "Put all mandatory parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Feedback"
                ],
                "summary": "Get All Feedback Category",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/feedback/category/create": {
            "post": {
                "description": "Put all mandatory parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Feedback"
                ],
                "summary": "Create Feedback Category",
                "parameters": [
                    {
                        "description": "CreateFeedbackCategoryRequest",
                        "name": "CreateFeedbackCategoryRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateFeedbackCategoryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateFeedbackCategoryResponse"
                        }
                    }
                }
            }
        },
        "/feedback/category/{id}": {
            "get": {
                "description": "Put all mandatory parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Feedback"
                ],
                "summary": "Get Feedback Category By Id",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/feedback/create": {
            "post": {
                "description": "Put all mandatory parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Feedback"
                ],
                "summary": "Create Feedback",
                "parameters": [
                    {
                        "description": "CreateFeedbackRequest",
                        "name": "CreateFeedbackRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateFeedbackRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateFeedbackResponse"
                        }
                    }
                }
            }
        },
        "/healthcheck": {
            "get": {
                "description": "Put all mandatory parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Healthcheck"
                ],
                "summary": "Check System Status",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.Status"
                            }
                        }
                    }
                }
            }
        },
        "/management/drug": {
            "get": {
                "description": "Put all mandatory parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Management"
                ],
                "summary": "Get All Drug",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/management/drug/create": {
            "post": {
                "description": "Put all mandatory parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Management"
                ],
                "summary": "Create Drug",
                "parameters": [
                    {
                        "description": "CreateDrugRequest",
                        "name": "CreateDrugRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateDrugRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.CreateDrugResponse"
                            }
                        }
                    }
                }
            }
        },
        "/profile": {
            "get": {
                "description": "Put all mandatory parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "summary": "Get User Profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/profile/edit": {
            "put": {
                "description": "Put all mandatory parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "summary": "Edit User Profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "EditUserRequest",
                        "name": "EditUserRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.EditUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.EditUserResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "dto.CreateAdminRequest": {
            "type": "object",
            "required": [
                "email",
                "fullname",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.CreateAdminResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "dto.CreateDoctorRequest": {
            "type": "object",
            "required": [
                "email",
                "fullname",
                "no_sip",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "no_sip": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.CreateDoctorResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "no_sip": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "dto.CreateDrugRequest": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dto.CreateDrugResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
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
        "dto.CreateFeedbackCategoryRequest": {
            "type": "object",
            "required": [
                "description",
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
        "dto.CreateFeedbackCategoryResponse": {
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
        "dto.CreateFeedbackRequest": {
            "type": "object",
            "required": [
                "detail",
                "feedback_category_id",
                "issue"
            ],
            "properties": {
                "detail": {
                    "type": "string"
                },
                "feedback_category_id": {
                    "type": "integer"
                },
                "issue": {
                    "type": "string"
                }
            }
        },
        "dto.CreateFeedbackResponse": {
            "type": "object",
            "required": [
                "feedback_category_id"
            ],
            "properties": {
                "detail": {
                    "type": "string"
                },
                "feedback_category_id": {
                    "type": "integer"
                },
                "issue": {
                    "type": "string"
                }
            }
        },
        "dto.CreatePharmacistRequest": {
            "type": "object",
            "required": [
                "email",
                "fullname",
                "no_sipa",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "no_sipa": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.CreatePharmacistResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "no_sipa": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "dto.CreateUserRequest": {
            "type": "object",
            "required": [
                "email",
                "fullname",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.CreateUserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "dto.EditUserRequest": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "fullname": {
                    "type": "string"
                },
                "height": {
                    "type": "number"
                },
                "no_sip": {
                    "type": "string"
                },
                "no_sipa": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "specialist": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "weight": {
                    "type": "number"
                }
            }
        },
        "dto.EditUserResponse": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "fullname": {
                    "type": "string"
                },
                "height": {
                    "type": "number"
                },
                "no_sip": {
                    "type": "string"
                },
                "no_sipa": {
                    "type": "string"
                },
                "specialist": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "weight": {
                    "type": "number"
                }
            }
        },
        "dto.ForgotPasswordRequest": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "dto.GoogleLoginRequest": {
            "type": "object",
            "required": [
                "token"
            ],
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "dto.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "dto.ResetPasswordRequest": {
            "type": "object",
            "required": [
                "password",
                "token"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "dto.Status": {
            "type": "object",
            "properties": {
                "data": {},
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "FarmaCare Server",
	Description:      "API definition for FarmaCare Server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
