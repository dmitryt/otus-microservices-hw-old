// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/dmitryt/otus-microservices-hw/hw02_k8s/controllers"
)

func init() {
	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSGet("/health", func(ctx *context.Context) {
			ctx.Output.JSON(map[string]string{"status": "OK"}, false, false)
		}),
	)
	beego.AddNamespace(ns)
}
