// Package docs_swagger Code generated by swaggo/swag. DO NOT EDIT
package docs_swagger

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
        "/api/v1/songs": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Метод получения данных библиотеки с фильтрацией по всем полям и пагинацией",
                "operationId": "getPageSongs",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Название группы",
                        "name": "group",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Название песни",
                        "name": "song",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Дата выхода песни",
                        "name": "release_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Текст книги",
                        "name": "text",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Ссылка на песню",
                        "name": "link",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Лимит",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Смещение",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Список песен",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.JSONSongModel"
                            }
                        }
                    },
                    "400": {
                        "description": "Неверный запрос",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Песни не найдены",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Метод добавления новой песни",
                "operationId": "addNewSong",
                "parameters": [
                    {
                        "description": "Группа и название песни",
                        "name": "nemSong",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.NewSongDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Успешное создание песни"
                    },
                    "400": {
                        "description": "Неверный запрос",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Песня не найдена",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "409": {
                        "description": "Песня уже существует",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/songs/{id}": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Метод изменения данных песни",
                "operationId": "updateByID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор песни",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Параметры песни",
                        "name": "newSongParams",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SongParamDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешное обновление данных песни"
                    },
                    "400": {
                        "description": "Неверный запрос",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Песня не найдена",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Метод удаления песни",
                "operationId": "deleteByID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор песни",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешное удаление песни"
                    },
                    "400": {
                        "description": "Неверный запрос",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Песни не найдены",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/songs/{id}/verses": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Метод получения данных библиотеки с фильтрацией по всем полям и пагинацией",
                "operationId": "getSongByVerses",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор песни",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Лимит",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Смещение",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Куплеты",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Неверный запрос",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Песня не найдена",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.ErrorResponse": {
            "type": "object",
            "properties": {
                "error_msg": {
                    "type": "string"
                }
            }
        },
        "dto.NewSongDTO": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                }
            }
        },
        "dto.SongParamDTO": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "releaseDate": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "models.JSONSongModel": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "release_date": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                },
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
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "LibrarySongs API",
	Description:      "API Server for LibrarySongs Application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
