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
        "/demo-docker/api/v1/hello": {
            "get": {
                "description": "HelloGet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "hello"
                ],
                "summary": "HelloGet",
                "operationId": "HelloGet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/hello.HelloResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResp"
                        }
                    }
                }
            },
            "post": {
                "description": "HelloPost",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "hello"
                ],
                "summary": "HelloGet",
                "operationId": "HelloPost",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/hello.HelloPostParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/hello.HelloResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResp"
                        }
                    }
                }
            }
        },
        "/demo-docker/api/v1/users": {
            "get": {
                "description": "获取用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "ListUser",
                "operationId": "ListUser",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "手机号",
                        "name": "mobile",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "姓名",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页数1开始",
                        "name": "pageOffset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ListUserResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResp"
                        }
                    }
                }
            },
            "post": {
                "description": "创建用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "CreateOrUpdateUser",
                "operationId": "CreateOrUpdateUser",
                "parameters": [
                    {
                        "description": "活动",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.CreateOrUpdateUserBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.ListUserResp": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.User"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "hello.HelloPostParam": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "hello.HelloResp": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "结果",
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "key": {
                    "type": "integer"
                },
                "mobile": {
                    "description": "电话",
                    "type": "string"
                },
                "mobileType": {
                    "description": "电话运营商",
                    "allOf": [
                        {
                            "$ref": "#/definitions/types.MobileType"
                        }
                    ]
                },
                "nickName": {
                    "description": "昵称",
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                },
                "userType": {
                    "description": "用户角色",
                    "allOf": [
                        {
                            "$ref": "#/definitions/types.UserType"
                        }
                    ]
                }
            }
        },
        "types.MobileType": {
            "type": "string",
            "enum": [
                "移动",
                "联通"
            ],
            "x-enum-comments": {
                "USER_TYPE_D": "管理员"
            },
            "x-enum-varnames": [
                "USER_TYPE_Y",
                "USER_TYPE_D"
            ]
        },
        "types.UserType": {
            "type": "integer",
            "enum": [
                0,
                1,
                2
            ],
            "x-enum-comments": {
                "USER_TYPE__ADMIN": "管理员",
                "USER_TYPE__USER": "用户"
            },
            "x-enum-varnames": [
                "USER_TYPE_UNKNOWN",
                "USER_TYPE__ADMIN",
                "USER_TYPE__USER"
            ]
        },
        "user.CreateOrUpdateUserBody": {
            "type": "object",
            "properties": {
                "mobile": {
                    "description": "手机号",
                    "type": "string"
                },
                "name": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "utils.ErrorResp": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
