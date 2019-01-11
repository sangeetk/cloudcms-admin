package routers

import (
	"git.urantiatech.com/cloudcms/cloudcms-admin/controllers"
	"github.com/urantiatech/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
