// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"mnp_api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/retailers",
			beego.NSInclude(
				&controllers.RetailerController{},
			),
		),
		beego.NSNamespace("/customers",
			beego.NSInclude(
				&controllers.CustomerController{},
			),
		),
		beego.NSNamespace("/orders",
			beego.NSInclude(
				&controllers.OrderController{},
			),
		),
		beego.NSNamespace("/services",
			beego.NSInclude(
				&controllers.ServiceController{},
			),
		),
		beego.NSNamespace("/tests",
			beego.NSInclude(
				&controllers.TestController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
