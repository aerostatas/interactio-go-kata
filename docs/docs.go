// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
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
        "/events": {
            "post": {
                "description": "Create a new event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Create a new event",
                "parameters": [
                    {
                        "description": "Event data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.EventCreate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/service.Event"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "service.Event": {
            "type": "object",
            "required": [
                "audioQuality",
                "date",
                "invitees",
                "languages",
                "name",
                "videoQuality"
            ],
            "properties": {
                "audioQuality": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "Low",
                        "High"
                    ]
                },
                "date": {
                    "type": "string",
                    "example": "2023-04-20T14:00:00Z"
                },
                "description": {
                    "type": "string",
                    "example": "EU Summit 2023"
                },
                "id": {
                    "type": "integer",
                    "example": 123
                },
                "invitees": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "one@email.com",
                        "two@email.com"
                    ]
                },
                "languages": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "Lithuanian",
                        "French"
                    ]
                },
                "name": {
                    "type": "string",
                    "example": "EU Summit"
                },
                "videoQuality": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "720p",
                        "1080p"
                    ]
                }
            }
        },
        "service.EventCreate": {
            "type": "object",
            "required": [
                "audioQuality",
                "date",
                "invitees",
                "languages",
                "name",
                "videoQuality"
            ],
            "properties": {
                "audioQuality": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "Low",
                        "High"
                    ]
                },
                "date": {
                    "type": "string",
                    "example": "2023-04-20T14:00:00Z"
                },
                "description": {
                    "type": "string",
                    "example": "EU Summit 2023"
                },
                "invitees": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "one@email.com",
                        "two@email.com"
                    ]
                },
                "languages": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "Lithuanian",
                        "French"
                    ]
                },
                "name": {
                    "type": "string",
                    "example": "EU Summit"
                },
                "videoQuality": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "720p",
                        "1080p"
                    ]
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Event API",
	Description:      "REST API that allows users to create events",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
