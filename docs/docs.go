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
                "summary": "Get user credential",
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
                "summary": "Forgot password",
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
                "summary": "Login user google",
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
                "summary": "Login user",
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
                "summary": "Register new user",
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
                "summary": "Register new doctor",
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
                "summary": "Register new pharmacist",
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
                "summary": "Reset password",
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
                "summary": "Check system status",
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
                "summary": "Get user profile",
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
                "summary": "Edit user profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "EditUserProfileRequest",
                        "name": "EditUserProfileRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.EditUserProfileRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.EditUserProfileResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
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
        "dto.EditUserProfileRequest": {
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
        "dto.EditUserProfileResponse": {
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
