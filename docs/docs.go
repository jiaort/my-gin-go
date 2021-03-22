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
        "termsOfService": "https://www.topgoer.com",
        "contact": {
            "name": "www.topgoer.com",
            "url": "https://www.topgoer.com",
            "email": "me@razeen.me"
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
        "/base/captcha": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Base"
                ],
                "summary": "生成验证码",
                "responses": {
                    "200": {
                        "description": "{\"code\":10001,\"data\":{},\"msg\":\"成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/base/verify": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Base"
                ],
                "summary": "验证验证码",
                "parameters": [
                    {
                        "description": "验证验证码",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CaptchaVerify"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":10001,\"data\":{},\"msg\":\"成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/detection/imgDetection": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Detection"
                ],
                "summary": "上传文件示例",
                "parameters": [
                    {
                        "type": "file",
                        "description": "上传文件示例",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":10001,\"data\":{},\"msg\":\"上传成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/getUserList": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "获取用户列表",
                "responses": {
                    "200": {
                        "description": "{\"code\":10001,\"data\":{},\"msg\":\"成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "用户登录",
                "responses": {
                    "200": {
                        "description": "{\"code\":10001,\"data\":{},\"msg\":\"成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "用户注册账号",
                "responses": {
                    "200": {
                        "description": "{\"code\":10001,\"data\":{},\"msg\":\"成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.CaptchaVerify": {
            "type": "object",
            "properties": {
                "captcha": {
                    "type": "string"
                },
                "captchaId": {
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
	Version:     "0.0.1",
	Host:        "10.0.24.5:8888",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "接口文档",
	Description: "图像识别服务",
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
