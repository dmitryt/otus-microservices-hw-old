package main

import (
	_ "github.com/dmitryt/otus-microservices-hw/hw02_k8s/routers"
	"github.com/dmitryt/otus-microservices-hw/hw02_k8s/controllers"
	"github.com/beego/beego/v2/client/orm"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)

	orm.RegisterDataBase("default", "postgres", "root:root@/orm_test?charset=utf8")
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
