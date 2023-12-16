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
        "/dev/func/add": {
            "post": {
                "tags": [
                    "鉴权接口-方法"
                ],
                "summary": "新增功能",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AccessToken",
                        "name": "AccessToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "新增功能信息",
                        "name": "FuncAddRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.DevFuncAddRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/dev/func/delete": {
            "delete": {
                "tags": [
                    "鉴权接口-方法"
                ],
                "summary": "删除功能",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AccessToken",
                        "name": "AccessToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "identity",
                        "name": "identity",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/dev/func/update": {
            "put": {
                "tags": [
                    "鉴权接口-方法"
                ],
                "summary": "修改功能",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AccessToken",
                        "name": "AccessToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "修改功能信息",
                        "name": "FuncUpdateRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.DevFuncUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/dev/menu/add": {
            "post": {
                "tags": [
                    "鉴权接口-方法"
                ],
                "summary": "新增菜单",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AccessToken",
                        "name": "AccessToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "新增菜单信息",
                        "name": "MenuAddRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.DevMenuAddRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/dev/menu/delete": {
            "delete": {
                "tags": [
                    "鉴权接口-方法"
                ],
                "summary": "删除菜单",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AccessToken",
                        "name": "AccessToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "identity",
                        "name": "identity",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/dev/menu/update": {
            "put": {
                "tags": [
                    "鉴权接口-方法"
                ],
                "summary": "修改菜单",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AccessToken",
                        "name": "AccessToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "修改菜单信息",
                        "name": "MenuUpdateRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.DevMenuUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/functions": {
            "get": {
                "tags": [
                    "登录信息认证-方法"
                ],
                "summary": "获取功能列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AccessToken",
                        "name": "AccessToken",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login/password": {
            "post": {
                "tags": [
                    "公共方法"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "登录信息",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.LoginPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/menus": {
            "get": {
                "tags": [
                    "登录信息认证-方法"
                ],
                "summary": "获取菜单列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AccessToken",
                        "name": "AccessToken",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/set/func/list": {
            "get": {
                "tags": [
                    "鉴权接口-方法"
                ],
                "summary": "获取功能列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AccessToken",
                        "name": "AccessToken",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/set/menu/list": {
            "get": {
                "tags": [
                    "鉴权接口-方法"
                ],
                "summary": "获取菜单列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AccessToken",
                        "name": "AccessToken",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/set/role/create": {
            "post": {
                "tags": [
                    "鉴权接口-方法"
                ],
                "summary": "角色创建",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AccessToken",
                        "name": "AccessToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "角色信息",
                        "name": "RoleCreateRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.SetRoleCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/set/role/delete": {
            "delete": {
                "tags": [
                    "鉴权接口-方法"
                ],
                "summary": "角色删除",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AccessToken",
                        "name": "AccessToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "identity",
                        "name": "identity",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/set/role/detail": {
            "get": {
                "tags": [
                    "鉴权接口-方法"
                ],
                "summary": "角色详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AccessToken",
                        "name": "AccessToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "identity",
                        "name": "identity",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/set/role/list": {
            "get": {
                "tags": [
                    "鉴权接口-方法"
                ],
                "summary": "获取角色列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AccessToken",
                        "name": "AccessToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "size",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "keyword",
                        "name": "keyword",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/set/role/update": {
            "put": {
                "tags": [
                    "鉴权接口-方法"
                ],
                "summary": "角色修改",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AccessToken",
                        "name": "AccessToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "所修改用户信息",
                        "name": "RoleUpdateRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.SetRoleUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/set/role/update/admin": {
            "put": {
                "tags": [
                    "鉴权接口-方法"
                ],
                "summary": "修改角色的管理员身份",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AccessToken",
                        "name": "AccessToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "管理员设置",
                        "name": "RoleUpdateAdminRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.SetRoleUpdateAdminRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/set/user/add": {
            "post": {
                "tags": [
                    "鉴权接口-方法"
                ],
                "summary": "管理员创建",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AccessToken",
                        "name": "AccessToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "新增管理员信息",
                        "name": "UserAddRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.SetUserAddRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/set/user/delete": {
            "delete": {
                "tags": [
                    "鉴权接口-方法"
                ],
                "summary": "管理员删除",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AccessToken",
                        "name": "AccessToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "identity",
                        "name": "identity",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/set/user/list": {
            "get": {
                "tags": [
                    "鉴权接口-方法"
                ],
                "summary": "管理员列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AccessToken",
                        "name": "AccessToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "size",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "keyword",
                        "name": "keyword",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/set/user/update": {
            "put": {
                "tags": [
                    "鉴权接口-方法"
                ],
                "summary": "管理员修改",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AccessToken",
                        "name": "AccessToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "新增管理员信息",
                        "name": "UserAddRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.SetUserUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/info": {
            "get": {
                "tags": [
                    "登录信息认证-方法"
                ],
                "summary": "获取用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AccessToken",
                        "name": "AccessToken",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/password/change": {
            "put": {
                "tags": [
                    "登录信息认证-方法"
                ],
                "summary": "修改密码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AccessToken",
                        "name": "AccessToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "新旧密码",
                        "name": "PasswordChangeRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.UserPasswordChangeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "service.DevFuncAddRequest": {
            "type": "object",
            "properties": {
                "menu_identity": {
                    "description": "菜单唯一标识",
                    "type": "string"
                },
                "name": {
                    "description": "功能名称",
                    "type": "string"
                },
                "sort": {
                    "description": "排序",
                    "type": "integer"
                },
                "uri": {
                    "description": "请求地址",
                    "type": "string"
                }
            }
        },
        "service.DevFuncUpdateRequest": {
            "type": "object",
            "properties": {
                "identity": {
                    "description": "功能唯一标识，必填",
                    "type": "string"
                },
                "menu_identity": {
                    "description": "菜单唯一标识",
                    "type": "string"
                },
                "name": {
                    "description": "功能名称",
                    "type": "string"
                },
                "sort": {
                    "description": "排序",
                    "type": "integer"
                },
                "uri": {
                    "description": "请求地址",
                    "type": "string"
                }
            }
        },
        "service.DevMenuAddRequest": {
            "type": "object",
            "properties": {
                "level": {
                    "description": "菜单等级，{0：目录，1：菜单，2：按钮}",
                    "type": "integer"
                },
                "name": {
                    "description": "菜单名称",
                    "type": "string"
                },
                "parent_identity": {
                    "description": "父级唯一标识，不填默认为顶级菜单",
                    "type": "string"
                },
                "path": {
                    "description": "路径",
                    "type": "string"
                },
                "sort": {
                    "description": "排序",
                    "type": "integer"
                }
            }
        },
        "service.DevMenuUpdateRequest": {
            "type": "object",
            "properties": {
                "identity": {
                    "description": "菜单唯一标识，必填",
                    "type": "string"
                },
                "level": {
                    "description": "菜单等级，{0：目录，1：菜单，2：按钮}",
                    "type": "integer"
                },
                "name": {
                    "description": "菜单名称",
                    "type": "string"
                },
                "parent_identity": {
                    "description": "父级唯一标识，不填默认为顶级菜单",
                    "type": "string"
                },
                "path": {
                    "description": "路径",
                    "type": "string"
                },
                "sort": {
                    "description": "排序",
                    "type": "integer"
                }
            }
        },
        "service.LoginPasswordRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "service.SetRoleCreateRequest": {
            "type": "object",
            "properties": {
                "is_admin": {
                    "type": "integer"
                },
                "menu_identities": {
                    "description": "被授权的菜单列表",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "sort": {
                    "type": "integer"
                }
            }
        },
        "service.SetRoleUpdateAdminRequest": {
            "type": "object",
            "properties": {
                "identity": {
                    "description": "角色唯一标识",
                    "type": "string"
                },
                "is_admin": {
                    "description": "是否是超管【0-否 1-是】",
                    "type": "integer"
                }
            }
        },
        "service.SetRoleUpdateRequest": {
            "type": "object",
            "properties": {
                "identity": {
                    "description": "角色唯一标识",
                    "type": "string"
                },
                "is_admin": {
                    "type": "integer"
                },
                "menu_identities": {
                    "description": "被授权的菜单列表",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "sort": {
                    "type": "integer"
                }
            }
        },
        "service.SetUserAddRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "role_identity": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "service.SetUserUpdateRequest": {
            "type": "object",
            "properties": {
                "identity": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "role_identity": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "service.UserPasswordChangeRequest": {
            "type": "object",
            "properties": {
                "new_password": {
                    "description": "新密码",
                    "type": "string"
                },
                "old_password": {
                    "description": "旧密码",
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
