package routers

import (
	"git.urantiatech.com/cloudcms/cloudcms-admin/controllers"
	"github.com/astaxie/beego"
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
		beego.Router("/admin/content/:name", &controllers.ContentController{}, "get:Index;post:Save;delete:Delete")
		beego.Router("/admin/content/:name/editor", &controllers.ContentController{}, "get:Editor")
		beego.Router("/admin/content/:name/delete", &controllers.ContentController{}, "get:Delete")
		beego.Router("/admin/logout", &controllers.LogoutController{})

		// Admin static files
		beego.SetStaticPath("/admin/theme", "static")
	}
	// beego.ErrorController(&controllers.ErrorController{})

}
