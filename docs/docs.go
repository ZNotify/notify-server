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
            "name": "Issues",
            "url": "https://github.com/ZNotify/server/issues"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "https://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "Provide UI",
                "produces": [
                    "text/html"
                ],
                "summary": "Web Index",
                "responses": {
                    "200": {
                        "description": "html",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/alive": {
            "get": {
                "description": "Check if the server is alive",
                "produces": [
                    "text/plain"
                ],
                "summary": "Server Heartbeat",
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/check": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Check if the user_id is valid",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Response-bool"
                        }
                    }
                }
            }
        },
        "/docs": {
            "get": {
                "produces": [
                    "text/plain"
                ],
                "summary": "Redirect to docs",
                "responses": {
                    "301": {
                        "description": "Moved Permanently",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/{user_id}/record": {
            "get": {
                "description": "Get recent 30days message record of user",
                "produces": [
                    "application/json"
                ],
                "summary": "Get record",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Response-array_entity_Message"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/types.UnauthorizedResponse"
                        }
                    }
                }
            }
        },
        "/{user_id}/s": {
            "put": {
                "description": "Send notification to user_id",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Send notification",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "content",
                        "name": "string",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Response-entity_Message"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.BadRequestResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/types.UnauthorizedResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Send notification to user_id",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Send notification",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "content",
                        "name": "string",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Response-entity_Message"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.BadRequestResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/types.UnauthorizedResponse"
                        }
                    }
                }
            }
        },
        "/{user_id}/send": {
            "put": {
                "description": "Send notification to user_id",
                "produces": [
                    "application/json"
                ],
                "summary": "Send notification",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "title",
                        "name": "title",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "content",
                        "name": "content",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "long",
                        "name": "long",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Response-entity_Message"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.BadRequestResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/types.UnauthorizedResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Send notification to user_id",
                "produces": [
                    "application/json"
                ],
                "summary": "Send notification",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "title",
                        "name": "title",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "content",
                        "name": "content",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "long",
                        "name": "long",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Response-entity_Message"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.BadRequestResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/types.UnauthorizedResponse"
                        }
                    }
                }
            }
        },
        "/{user_id}/token/{device_id}": {
            "put": {
                "description": "Create or update token of device",
                "produces": [
                    "application/json"
                ],
                "summary": "Create or update token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "device_id should be a valid UUID",
                        "name": "device_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "TelegramHost",
                            "WebSocketHost",
                            "FCM",
                            "WebPush"
                        ],
                        "type": "string",
                        "description": "channel can be used. Sometimes less than document.",
                        "name": "channel",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Response-bool"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.BadRequestResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/types.UnauthorizedResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete token of device",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "device_id",
                        "name": "device_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Response-bool"
                        }
                    }
                }
            }
        },
        "/{user_id}/{id}": {
            "get": {
                "description": "Get message record detail of a message",
                "produces": [
                    "application/json"
                ],
                "summary": "Get message record detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Response-entity_Message"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.BadRequestResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/types.UnauthorizedResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/types.NotFoundResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete message record with id",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete message record",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Response-bool"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/types.UnauthorizedResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Message": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "long": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "types.BadRequestResponse": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "code": {
                    "type": "integer",
                    "default": 400
                }
            }
        },
        "types.NotFoundResponse": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "code": {
                    "type": "integer",
                    "default": 404
                }
            }
        },
        "types.Response-array_entity_Message": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Message"
                    }
                },
                "code": {
                    "type": "integer",
                    "default": 200
                }
            }
        },
        "types.Response-bool": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "boolean"
                },
                "code": {
                    "type": "integer",
                    "default": 200
                }
            }
        },
        "types.Response-entity_Message": {
            "type": "object",
            "properties": {
                "body": {
                    "$ref": "#/definitions/entity.Message"
                },
                "code": {
                    "type": "integer",
                    "default": 200
                }
            }
        },
        "types.UnauthorizedResponse": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "code": {
                    "type": "integer",
                    "default": 401
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
	Title:            "Notify API",
	Description:      "This is Znotify api server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
