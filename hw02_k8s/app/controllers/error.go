package controllers

import (
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type ErrorController struct {
	web.Controller
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	NotFoundError       = Response{Code: 404, Message: "Not Found"}
	InternalServerError = Response{Code: 500, Message: "Internal Server Error"}
)

func NewResponseWithDefaultMessage(code int) *Response {
	return &Response{Code: code, Message: http.StatusText(code)}
}

func NewResponse(code int, msg string) *Response {
	return &Response{Code: code, Message: msg}
}

func (c *ErrorController) Error404() {
	c.Data["json"] = &NotFoundError
	_ = c.ServeJSON()
}

func (c *ErrorController) Error500() {
	c.Data["json"] = &InternalServerError
	_ = c.ServeJSON()
}

func (c *ErrorController) ErrorDB() {
	c.Data["content"] = "database is now down"
	c.TplName = "dberror.tpl"
}
