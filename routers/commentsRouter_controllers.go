package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["mnp_api/controllers:CustomerController"] = append(beego.GlobalControllerRouter["mnp_api/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mnp_api/controllers:CustomerController"] = append(beego.GlobalControllerRouter["mnp_api/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mnp_api/controllers:CustomerController"] = append(beego.GlobalControllerRouter["mnp_api/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:customerID`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mnp_api/controllers:CustomerController"] = append(beego.GlobalControllerRouter["mnp_api/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:customerId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mnp_api/controllers:CustomerController"] = append(beego.GlobalControllerRouter["mnp_api/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:customerId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mnp_api/controllers:CustomerController"] = append(beego.GlobalControllerRouter["mnp_api/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "GetByImei",
			Router: `/by-imei/:imei`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mnp_api/controllers:CustomerController"] = append(beego.GlobalControllerRouter["mnp_api/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "GetAllOrders",
			Router: `/orders/:customerId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mnp_api/controllers:CustomerController"] = append(beego.GlobalControllerRouter["mnp_api/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "GetPendingOrders",
			Router: `/pending-orders/:customerId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mnp_api/controllers:OrderController"] = append(beego.GlobalControllerRouter["mnp_api/controllers:OrderController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mnp_api/controllers:OrderController"] = append(beego.GlobalControllerRouter["mnp_api/controllers:OrderController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mnp_api/controllers:OrderController"] = append(beego.GlobalControllerRouter["mnp_api/controllers:OrderController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:orderID`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mnp_api/controllers:OrderController"] = append(beego.GlobalControllerRouter["mnp_api/controllers:OrderController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:orderId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mnp_api/controllers:RetailerController"] = append(beego.GlobalControllerRouter["mnp_api/controllers:RetailerController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mnp_api/controllers:RetailerController"] = append(beego.GlobalControllerRouter["mnp_api/controllers:RetailerController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mnp_api/controllers:RetailerController"] = append(beego.GlobalControllerRouter["mnp_api/controllers:RetailerController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:retailerID`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mnp_api/controllers:RetailerController"] = append(beego.GlobalControllerRouter["mnp_api/controllers:RetailerController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:retailerId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mnp_api/controllers:RetailerController"] = append(beego.GlobalControllerRouter["mnp_api/controllers:RetailerController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:retailerId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mnp_api/controllers:RetailerController"] = append(beego.GlobalControllerRouter["mnp_api/controllers:RetailerController"],
		beego.ControllerComments{
			Method: "GetByCode",
			Router: `/by-code/:code`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mnp_api/controllers:RetailerController"] = append(beego.GlobalControllerRouter["mnp_api/controllers:RetailerController"],
		beego.ControllerComments{
			Method: "GetByImei",
			Router: `/by-imei/:imei`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mnp_api/controllers:RetailerController"] = append(beego.GlobalControllerRouter["mnp_api/controllers:RetailerController"],
		beego.ControllerComments{
			Method: "Import",
			Router: `/import`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mnp_api/controllers:RetailerController"] = append(beego.GlobalControllerRouter["mnp_api/controllers:RetailerController"],
		beego.ControllerComments{
			Method: "Nearby",
			Router: `/nearby`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mnp_api/controllers:RetailerController"] = append(beego.GlobalControllerRouter["mnp_api/controllers:RetailerController"],
		beego.ControllerComments{
			Method: "GetAllOrders",
			Router: `/orders/:retailerId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mnp_api/controllers:RetailerController"] = append(beego.GlobalControllerRouter["mnp_api/controllers:RetailerController"],
		beego.ControllerComments{
			Method: "GetPending",
			Router: `/pending-orders/:retailerId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mnp_api/controllers:ServiceController"] = append(beego.GlobalControllerRouter["mnp_api/controllers:ServiceController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mnp_api/controllers:ServiceController"] = append(beego.GlobalControllerRouter["mnp_api/controllers:ServiceController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mnp_api/controllers:ServiceController"] = append(beego.GlobalControllerRouter["mnp_api/controllers:ServiceController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:serviceID`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mnp_api/controllers:ServiceController"] = append(beego.GlobalControllerRouter["mnp_api/controllers:ServiceController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:serviceId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mnp_api/controllers:ServiceController"] = append(beego.GlobalControllerRouter["mnp_api/controllers:ServiceController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:serviceId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mnp_api/controllers:TestController"] = append(beego.GlobalControllerRouter["mnp_api/controllers:TestController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
