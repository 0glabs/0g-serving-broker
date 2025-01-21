// Code generated by swaggo/swag. DO NOT EDIT
package doc

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
        "/quote": {
            "get": {
                "description": "This endpoint allows you to get a quote",
                "tags": [
                    "quote"
                ],
                "operationId": "getQuote",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/{userAddress}/task": {
            "get": {
                "description": "This endpoint allows you to list tasks by user address",
                "tags": [
                    "task"
                ],
                "operationId": "listTask",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user address",
                        "name": "userAddress",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "latest tasks",
                        "name": "latest",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schema.Task"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "This endpoint allows you to create a fine-tuning task",
                "tags": [
                    "task"
                ],
                "operationId": "createTask",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user address",
                        "name": "userAddress",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.Task"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content - success without response body"
                    }
                }
            }
        },
        "/user/{userAddress}/task/{taskID}": {
            "get": {
                "description": "This endpoint allows you to get a task by ID",
                "tags": [
                    "task"
                ],
                "operationId": "getTask",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user address",
                        "name": "userAddress",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "task ID",
                        "name": "taskID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.Task"
                        }
                    }
                }
            }
        },
        "/user/{userAddress}/task/{taskID}/log": {
            "get": {
                "description": "This endpoint allows you to get the progress log of a task by ID",
                "produces": [
                    "application/octet-stream"
                ],
                "tags": [
                    "task"
                ],
                "operationId": "getTaskProgress",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user address",
                        "name": "userAddress",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "task ID",
                        "name": "taskID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "progress.log",
                        "schema": {
                            "type": "file"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "schema.Task": {
            "type": "object",
            "required": [
                "datasetHash",
                "fee",
                "nonce",
                "preTrainedModelHash",
                "serviceName",
                "signature",
                "trainingParams",
                "userAddress"
            ],
            "properties": {
                "createdAt": {
                    "type": "string",
                    "readOnly": true
                },
                "datasetHash": {
                    "type": "string"
                },
                "deliverIndex": {
                    "type": "integer",
                    "readOnly": true
                },
                "fee": {
                    "type": "string"
                },
                "id": {
                    "type": "string",
                    "readOnly": true
                },
                "nonce": {
                    "type": "string"
                },
                "preTrainedModelHash": {
                    "type": "string"
                },
                "serviceName": {
                    "type": "string"
                },
                "signature": {
                    "type": "string"
                },
                "status": {
                    "type": "string",
                    "readOnly": true
                },
                "trainingParams": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string",
                    "readOnly": true
                },
                "userAddress": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.2.0",
	Host:             "localhost:3080",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "0G Serving Provider Broker API",
	Description:      "These APIs allows customers to interact with the 0G Compute Fine Tune Service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
