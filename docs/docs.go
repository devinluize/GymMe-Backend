// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "dasdasdas",
            "url": "asdasdas",
            "email": "tes@gmail.com"
        },
        "license": {
            "name": "Gymme"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/bookmark": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Remove Bookmark",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bookmark"
                ],
                "summary": "Remove Bookmark",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "information_id",
                        "name": "information_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Bookmark"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    }
                }
            }
        },
        "/api/information": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get All Information By Pagination",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Information"
                ],
                "summary": "Get All Information By Pagination",
                "parameters": [
                    {
                        "type": "string",
                        "description": "sort_by",
                        "name": "sort_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "sort_of",
                        "name": "sort_of",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.InformationEntities"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create New Information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Information"
                ],
                "summary": "Create New Information",
                "parameters": [
                    {
                        "description": "Insert Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/MenuPayloads.InformationInsertPayloads"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.StandarAPIResponses"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update Information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Information"
                ],
                "summary": "Update Information",
                "parameters": [
                    {
                        "description": "Update Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/MenuPayloads.InformationUpdatePayloads"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    }
                }
            }
        },
        "/api/information/by-id/{information_id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get Information By Information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Information"
                ],
                "summary": "Get Information By Information",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "information_id",
                        "name": "information_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.InformationEntities"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    }
                }
            }
        },
        "/api/information/delete/{information_id}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete Information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Information"
                ],
                "summary": "Delete Information",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "information_id",
                        "name": "information_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.StandarAPIResponses"
                        }
                    }
                }
            }
        },
        "/api/profile": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get Profile Detail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "summary": "Get Profile Detail",
                "parameters": [
                    {
                        "description": "user detail Headers",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/MenuPayloads.ProfilePayloadRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get Profile Detail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "summary": "Get Profile Detail",
                "parameters": [
                    {
                        "description": "user detail Headers",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/MenuPayloads.ProfilePayloadRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    }
                }
            }
        },
        "/api/profile/{user_id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get Profile Detail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "summary": "Get Profile Detail",
                "parameters": [
                    {
                        "type": "integer",
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
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    }
                }
            }
        },
        "/api/user/login2": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Login by header",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login by header",
                "parameters": [
                    {
                        "description": "Insert Header Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payloads.LoginPaylods"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    }
                }
            }
        },
        "/api/user/register": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "register by heade",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "register by heade",
                "parameters": [
                    {
                        "description": "Insert Header Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payloads.RegisterPayloads"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    }
                }
            }
        },
        "/api/weight": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create Weight History",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Weight"
                ],
                "summary": "Create Weight History",
                "parameters": [
                    {
                        "description": "Insert Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/MenuPayloads.WeightHistoryPayloads"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.StandarAPIResponses"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    }
                }
            }
        },
        "/api/weight/delete/{weight_id}/{user_id}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create Weight History",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Weight"
                ],
                "summary": "Create Weight History",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "weight_id",
                        "name": "weight_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
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
                            "$ref": "#/definitions/responses.StandarAPIResponses"
                        }
                    }
                }
            }
        },
        "/api/weight/{user_id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get Weight Pagination History",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Weight"
                ],
                "summary": "Get Weight Pagination History",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "sort_by",
                        "name": "sort_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "sort_of",
                        "name": "sort_of",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.WeightHistoryEntities"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponses"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "MenuPayloads.InformationBodyDetail": {
            "type": "object",
            "properties": {
                "information_body_paragraph": {
                    "type": "string"
                },
                "information_image_content_path": {
                    "type": "string"
                }
            }
        },
        "MenuPayloads.InformationInsertPayloads": {
            "type": "object",
            "properties": {
                "information_body_paragraph": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/MenuPayloads.InformationBodyDetail"
                    }
                },
                "information_header": {
                    "type": "string"
                },
                "information_id": {
                    "type": "integer"
                },
                "information_image_content_path_1": {
                    "type": "string"
                },
                "information_image_content_path_2": {
                    "type": "string"
                },
                "information_image_content_path_3": {
                    "type": "string"
                },
                "information_image_content_path_4": {
                    "type": "string"
                },
                "information_image_content_path_5": {
                    "type": "string"
                },
                "information_type_id": {
                    "type": "integer"
                }
            }
        },
        "MenuPayloads.InformationUpdatePayloads": {
            "type": "object",
            "properties": {
                "information_body_content": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/MenuPayloads.InformationBodyDetail"
                    }
                },
                "information_id": {
                    "type": "integer"
                }
            }
        },
        "MenuPayloads.ProfilePayloadRequest": {
            "type": "object",
            "properties": {
                "user_detail_id": {
                    "type": "integer"
                },
                "user_gender": {
                    "type": "string"
                },
                "user_height": {
                    "type": "number"
                },
                "user_id": {
                    "type": "integer"
                },
                "user_profile_description": {
                    "type": "string"
                },
                "user_profile_image": {
                    "type": "string"
                },
                "user_weight": {
                    "type": "number"
                }
            }
        },
        "MenuPayloads.WeightHistoryPayloads": {
            "type": "object",
            "properties": {
                "user_id": {
                    "type": "integer"
                },
                "user_weight": {
                    "type": "number"
                },
                "user_weight_time": {
                    "type": "string"
                }
            }
        },
        "entities.Bookmark": {
            "type": "object",
            "properties": {
                "bookmarkType": {
                    "$ref": "#/definitions/entities.BookmarkType"
                },
                "bookmark_id": {
                    "type": "integer"
                },
                "bookmark_type_id": {
                    "type": "integer"
                },
                "information": {
                    "$ref": "#/definitions/entities.InformationEntities"
                },
                "information_id": {
                    "type": "integer"
                },
                "user": {
                    "$ref": "#/definitions/entities.Users"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "entities.BookmarkType": {
            "type": "object",
            "properties": {
                "bookmark_type_id": {
                    "type": "integer"
                },
                "bookmark_type_name": {
                    "type": "string"
                }
            }
        },
        "entities.InformationBodyEntities": {
            "type": "object",
            "properties": {
                "information_body_id": {
                    "type": "integer"
                },
                "information_body_paragraph": {
                    "type": "string"
                },
                "information_id": {
                    "type": "integer"
                },
                "information_image_content_path": {
                    "type": "string"
                }
            }
        },
        "entities.InformationEntities": {
            "type": "object",
            "properties": {
                "informationType": {
                    "$ref": "#/definitions/entities.InformationType"
                },
                "information_body": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.InformationBodyEntities"
                    }
                },
                "information_created_by_user_id": {
                    "type": "integer"
                },
                "information_date_created": {
                    "type": "string"
                },
                "information_header": {
                    "type": "string"
                },
                "information_id": {
                    "type": "integer"
                },
                "information_type_id": {
                    "type": "integer"
                }
            }
        },
        "entities.InformationType": {
            "type": "object",
            "properties": {
                "information_type_id": {
                    "type": "integer"
                },
                "information_type_name": {
                    "type": "string"
                }
            }
        },
        "entities.UserDetail": {
            "type": "object",
            "properties": {
                "user_detail_id": {
                    "type": "integer"
                },
                "user_gender": {
                    "type": "string"
                },
                "user_height": {
                    "type": "number"
                },
                "user_id": {
                    "type": "integer"
                },
                "user_phone_number": {
                    "type": "string"
                },
                "user_profile_description": {
                    "type": "string"
                },
                "user_profile_image": {
                    "type": "string"
                },
                "user_weight": {
                    "type": "number"
                }
            }
        },
        "entities.Users": {
            "type": "object",
            "properties": {
                "user_detail": {
                    "description": "IsVIP        bool       ` + "`" + `gorm:\"column:is_vip\" json:\"is_vip\"` + "`" + `",
                    "allOf": [
                        {
                            "$ref": "#/definitions/entities.UserDetail"
                        }
                    ]
                },
                "user_email": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                },
                "user_name": {
                    "type": "string"
                },
                "user_password": {
                    "type": "string"
                }
            }
        },
        "entities.WeightHistoryEntities": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/entities.Users"
                },
                "user_id": {
                    "type": "integer"
                },
                "user_weight": {
                    "type": "number"
                },
                "user_weight_time": {
                    "type": "string"
                },
                "weight_history_id": {
                    "type": "integer"
                }
            }
        },
        "payloads.LoginPaylods": {
            "type": "object",
            "properties": {
                "useremail": {
                    "type": "string"
                },
                "userpasword": {
                    "type": "string"
                }
            }
        },
        "payloads.RegisterPayloads": {
            "type": "object",
            "properties": {
                "user_email": {
                    "type": "string"
                },
                "user_gender": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                },
                "user_password": {
                    "type": "string"
                },
                "user_phone_number": {
                    "type": "string"
                }
            }
        },
        "responses.ErrorResponses": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "responses.StandarAPIResponses": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                },
                "success": {
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Gym Me Backend Thesis",
	Description:      "Gym Me BackEnd Thesis",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
