// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-04-06 12:02:28.815599 +0300 IDT m=+0.025459525

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "1.0"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/watcher/checkin": {
            "post": {
                "description": "get request to perform checkin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Watcher"
                ],
                "summary": "Perform checkin via timewatch.co.il",
                "operationId": "checkin",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "user_data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.UserData"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Max delay in minutes until request is sent (random)",
                        "name": "max_delay",
                        "in": "query"
                    }
                ]
            }
        },
        "/v1/watcher/checkout": {
            "post": {
                "description": "get request to perform checkout",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Watcher"
                ],
                "summary": "Perform checkout via timewatch.co.il",
                "operationId": "checkout",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "user_data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.UserData"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Max delay in minutes until request is sent (random)",
                        "name": "max_delay",
                        "in": "query"
                    }
                ]
            }
        }
    },
    "definitions": {
        "model.UserData": {
            "type": "object",
            "required": [
                "company",
                "password",
                "user"
            ],
            "properties": {
                "company": {
                    "type": "string",
                    "example": "5687"
                },
                "password": {
                    "type": "string",
                    "example": "Baba"
                },
                "user": {
                    "type": "string",
                    "example": "1234"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
