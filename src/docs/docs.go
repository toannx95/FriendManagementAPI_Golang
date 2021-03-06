// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-04-27 17:39:39.2189216 +0700 +07 m=+0.048474201

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
            "name": "API Support",
            "email": "soberkoder@swagger.io"
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
        "/friends/block": {
            "post": {
                "description": "API to block updates from an email address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Friend"
                ],
                "summary": "CreateBlock",
                "parameters": [
                    {
                        "description": "requestDto",
                        "name": "requestDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RequestDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Success"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/friends/create-friend": {
            "post": {
                "description": "API to create friend connection between two email addresses",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Friend"
                ],
                "summary": "CreateFriend",
                "parameters": [
                    {
                        "description": "friendDto",
                        "name": "friendDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.FriendDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Success"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/friends/get-common-friends-list": {
            "post": {
                "description": "API to retrieve the common friends list between two email addresses",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Friend"
                ],
                "summary": "GetCommonFriends",
                "parameters": [
                    {
                        "description": "friendDto",
                        "name": "friendDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.FriendDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Success"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/friends/get-friends-list": {
            "post": {
                "description": "API to retrieve the friends list for an email address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Friend"
                ],
                "summary": "GetFriendsListByEmail",
                "parameters": [
                    {
                        "description": "emailDto",
                        "name": "emailDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.EmailDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Success"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/friends/get-receivers-list": {
            "post": {
                "description": "API to retrieve all email addresses that can receive updates from an email address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Friend"
                ],
                "summary": "GetReceiversList",
                "parameters": [
                    {
                        "description": "senderDto",
                        "name": "senderDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SenderDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Success"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/friends/subscribe": {
            "post": {
                "description": "API to subscribe to updates from an email address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Friend"
                ],
                "summary": "CreateSubscribe",
                "parameters": [
                    {
                        "description": "requestDto",
                        "name": "requestDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RequestDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Success"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "API to get all emails",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "GetAllUsers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/users/create-user": {
            "post": {
                "description": "API to create new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "CreateUser",
                "parameters": [
                    {
                        "description": "emailDto",
                        "name": "emailDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.EmailDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Success"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.EmailDto": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "dto.FriendDto": {
            "type": "object",
            "properties": {
                "friends": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "dto.RequestDto": {
            "type": "object",
            "properties": {
                "requestor": {
                    "type": "string"
                },
                "target": {
                    "type": "string"
                }
            }
        },
        "dto.SenderDto": {
            "type": "object",
            "properties": {
                "sender": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "response.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "friends": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "response.Success": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
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
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Friend Management API",
	Description: "This is a sample service for managing Friend Management",
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
