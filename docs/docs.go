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
        "contact": {
            "name": "Gabriel de Freitas Pinheiro",
            "url": "https://github.com/GabrielFreitasP",
            "email": "gabrieldefreitaspinheiro@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/search": {
            "post": {
                "description": "Search Roman numeral from text",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Roman Numeral"
                ],
                "summary": "Search Roman numeral",
                "parameters": [
                    {
                        "description": "Roman Numeral Search",
                        "name": "romanNumeralSearch",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RomanNumeralSearch"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.RomanNumeral"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httpErrors.RestError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httpErrors.RestError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpErrors.RestError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "httpErrors.RestError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "models.RomanNumeral": {
            "type": "object",
            "properties": {
                "number": {
                    "type": "string"
                },
                "value": {
                    "type": "integer"
                }
            }
        },
        "models.RomanNumeralSearch": {
            "type": "object",
            "properties": {
                "text": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Smallest Roman Numeral API",
	Description:      "Smallest Roman Numeral API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
