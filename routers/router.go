package routers

import (
	"git.urantiatech.com/cloudcms/cloudcms-admin/controllers"
	"github.com/urantiatech/beego"
)

func init() {

	// Static files
	beego.SetStaticPath("/fonts", "static/fonts")
	beego.SetStaticPath("/uploads", "uploads")

	if beego.AppConfig.String("adminuser") != "" &&
		beego.AppConfig.String("adminpass") != "" {
		// Admin routes
		beego.Router("/admin", &controllers.AdminController{})
		beego.Router("/admin/dashboard", &controllers.DashboardController{})
		beego.Router("/admin/microservice/:name", &controllers.MicroserviceController{}, "get:Index;post:Save")
		beego.Router("/admin/microservice/:name/edit", &controllers.MicroserviceController{}, "get:Edit")
		beego.Router("/admin/logout", &controllers.LogoutController{})

		// Admin static files
		beego.SetStaticPath("/badmin", "static/badmin")
	}
	// beego.ErrorController(&controllers.ErrorController{})

}
