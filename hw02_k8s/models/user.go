package models

import (
	"encoding/json"

	"github.com/beego/beego/v2/client/orm"
	"github.com/dmitryt/otus-microservices-hw/hw02_k8s/utils"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	orm.RegisterModel(new(User))
	validate = validator.New()
	validate.RegisterValidation("phone", utils.ValidatePhone)
}

type User struct {
	ID        int64  `orm:"unique;column(id)" json:"id"`
	Username  string `validate:"required" json:"username"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `validate:"email" json:"email,omitempty"`
	Phone     string `validate:"phone" json:"phone,omitempty"`
}

func (u *User) TableName() string {
	// db table name
	return "users"
}

func AddUser(input []byte) (u *User, err error) {
	var parsedUser User
	err = json.Unmarshal(input, &parsedUser)
	if err != nil {
		return
	}
	err = validate.Struct(parsedUser)
	if err != nil {
		return
	}
	o := orm.NewOrm()
	id, err := o.Insert(&parsedUser)
	if err != nil {
		return
	}
	u = &User{ID: id}
	err = o.Read(u)

	return
}

func GetUser(uid int64) (u User, err error) {
	o := orm.NewOrm()
	u.ID = uid
	err = o.Read(&u)

	return
}

func GetAllUsers() (users []User, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(new(User)).All(&users)

	return
}

func UpdateUser(uid int64, input []byte) (u *User, err error) {
	var parsedUser User
	if err = json.Unmarshal(input, &parsedUser); err != nil {
		return
	}
	err = validate.Struct(parsedUser)
	if err != nil {
		return
	}
	o := orm.NewOrm()
	if err = o.Read(&User{ID: uid}); err != nil {
		return
	}
	parsedUser.ID = uid
	_, err = o.Update(&parsedUser)

	return &parsedUser, err
}

func DeleteUser(uid int64) (err error) {
	o := orm.NewOrm()
	_, err = o.Delete(&User{ID: uid})

	return
}
