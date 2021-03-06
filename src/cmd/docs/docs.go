// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Nam Nguyen",
            "email": "nvnam.c@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/buy/:id": {
            "post": {
                "description": "Buy wager",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wagers"
                ],
                "summary": "Buy wager",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "wager id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "values",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.BuyWagerRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.BuyWagerResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    }
                }
            }
        },
        "/wagers": {
            "get": {
                "description": "List wagers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wagers"
                ],
                "summary": "List wagers",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "cursor",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "place",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.ListWagersResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    }
                }
            },
            "post": {
                "description": "Place wager",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wagers"
                ],
                "summary": "Place wager",
                "parameters": [
                    {
                        "description": "body",
                        "name": "values",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.PlaceWagerRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.ListWagersResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.BuyWagerRequest": {
            "type": "object",
            "properties": {
                "buying_price": {
                    "type": "number"
                }
            }
        },
        "dtos.BuyWagerResponse": {
            "type": "object",
            "properties": {
                "bought_at": {
                    "type": "integer"
                },
                "buying_price": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "wager_id": {
                    "type": "integer"
                }
            }
        },
        "dtos.ListWagersRequest": {
            "type": "object",
            "properties": {
                "cursor": {
                    "type": "integer"
                },
                "place": {
                    "type": "integer"
                }
            }
        },
        "dtos.ListWagersResponse": {
            "type": "object",
            "properties": {
                "amount_sold": {
                    "type": "integer"
                },
                "current_selling_price": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "odds": {
                    "type": "integer"
                },
                "percentage_sold": {
                    "type": "number"
                },
                "placed_at": {
                    "type": "integer"
                },
                "selling_percentage": {
                    "type": "number"
                },
                "selling_price": {
                    "type": "number"
                },
                "total_wager_value": {
                    "type": "number"
                }
            }
        },
        "dtos.PlaceWagerRequest": {
            "type": "object",
            "properties": {
                "odds": {
                    "type": "integer"
                },
                "selling_percentage": {
                    "type": "integer"
                },
                "selling_price": {
                    "type": "number"
                },
                "total_wager_value": {
                    "type": "number"
                }
            }
        },
        "errors.AppError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/prpcl/v1",
	Schemes:     []string{},
	Title:       "prpcl",
	Description: "prpcl API documentation",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
