package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type ErrorController struct {
	web.Controller
}

type ResponseError struct {
	Code int32 `json:"code"`
	Message string `json:"message"`
}

var NotFoundError = ResponseError{Code: 404, Message: "Not Found"}
var InternalServerError = ResponseError{Code: 500, Message: "Internal Server Error"}

func (c *ErrorController) Error404() {
	c.Data["json"] = &NotFoundError
	c.ServeJSON()
}

func (c *ErrorController) Error500() {
	c.Data["json"] = &InternalServerError
	c.ServeJSON()
}

func (c *ErrorController) ErrorDb() {
	c.Data["content"] = "database is now down"
	c.TplName = "dberror.tpl"
}
