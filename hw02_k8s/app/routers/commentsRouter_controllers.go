package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {
	beego.GlobalControllerRouter["github.com/dmitryt/otus-microservices-hw/hw02_k8s/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/dmitryt/otus-microservices-hw/hw02_k8s/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		})

	beego.GlobalControllerRouter["github.com/dmitryt/otus-microservices-hw/hw02_k8s/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/dmitryt/otus-microservices-hw/hw02_k8s/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		})

	beego.GlobalControllerRouter["github.com/dmitryt/otus-microservices-hw/hw02_k8s/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/dmitryt/otus-microservices-hw/hw02_k8s/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/:uid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		})

	beego.GlobalControllerRouter["github.com/dmitryt/otus-microservices-hw/hw02_k8s/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/dmitryt/otus-microservices-hw/hw02_k8s/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:uid",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		})

	beego.GlobalControllerRouter["github.com/dmitryt/otus-microservices-hw/hw02_k8s/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/dmitryt/otus-microservices-hw/hw02_k8s/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:uid",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		})
}
