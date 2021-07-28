package main

import (
	"log"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/dmitryt/otus-microservices-hw/hw02_k8s/controllers"
	_ "github.com/dmitryt/otus-microservices-hw/hw02_k8s/routers"
	"github.com/dmitryt/otus-microservices-hw/hw02_k8s/utils"
	_ "github.com/lib/pq"
)

func init() {
	orm.Debug = true

	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		log.Fatal("Error during registering the DB driver", err)
	}
	sdn, err := utils.GetSQLDSN()
	if err != nil {
		log.Fatal("DB config is invalid", err)
	}
	err = orm.RegisterDataBase("default", "postgres", sdn)
	if err != nil {
		log.Fatal("Error during registering the default DB", err)
	}
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
