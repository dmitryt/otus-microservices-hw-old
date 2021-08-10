package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/dmitryt/otus-microservices-hw/hw02_k8s/models"
	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
)

// Operations about Users.
type UserController struct {
	beego.Controller
}

func prepareError(err error) (result *Response) {
	var code int
	var message string

	pqErr, _ := err.(*pq.Error)
	switch err.(type) {
	case *strconv.NumError:
		code = http.StatusBadRequest
	case validator.ValidationErrors:
		code = http.StatusBadRequest
		message = err.Error()
	case *pq.Error:
		code = http.StatusBadRequest
		if pqErr.Constraint != "" {
			message = pqErr.Detail
		}
	default:
		if errors.Is(err, orm.ErrNoRows) {
			code = http.StatusNotFound
		} else {
			code = http.StatusBadGateway
		}
	}
	logs.Info("Response ERR", err)

	if message == "" {
		return NewResponseWithDefaultMessage(code)
	}

	return NewResponse(code, message)
}

func (u *UserController) sendResponse(r interface{}) {
	u.Data["json"] = r
	_ = u.ServeJSON()
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post].
func (u *UserController) Post() {
	user, err := models.AddUser(u.Ctx.Input.RequestBody)
	if err != nil {
		u.sendResponse(prepareError(err))

		return
	}
	u.sendResponse(user)
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get].
func (u *UserController) GetAll() {
	users, err := models.GetAllUsers()
	if err != nil {
		u.sendResponse(prepareError(err))

		return
	}
	u.sendResponse(users)
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get].
func (u *UserController) Get() {
	uid, err := u.GetInt64(":uid")
	if err != nil {
		u.sendResponse(prepareError(err))

		return
	}
	user, err := models.GetUser(uid)
	if err != nil {
		u.sendResponse(prepareError(err))

		return
	}
	u.sendResponse(user)
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put].
func (u *UserController) Put() {
	uid, err := u.GetInt64(":uid")
	if err != nil {
		u.sendResponse(prepareError(err))

		return
	}
	user, err := models.UpdateUser(uid, u.Ctx.Input.RequestBody)
	if err != nil {
		u.sendResponse(prepareError(err))

		return
	}
	u.sendResponse(user)
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete].
func (u *UserController) Delete() {
	uid, err := u.GetInt64(":uid")
	if err != nil {
		u.sendResponse(prepareError(err))

		return
	}
	_, err = models.GetUser(uid)
	if err != nil {
		u.sendResponse(prepareError(err))

		return
	}
	err = models.DeleteUser(uid)
	if err != nil {
		u.sendResponse(prepareError(err))

		return
	}
	u.sendResponse(NewResponseWithDefaultMessage(http.StatusNoContent))
}
