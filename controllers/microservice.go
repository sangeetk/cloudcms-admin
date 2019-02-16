package controllers

import (
	"net/http"
	"strings"

	// "git.urantiatech.com/cloudcms/cloudcms/api"
	"github.com/urantiatech/beego"
)

// MicroserviceController definition
type MicroserviceController struct {
	beego.Controller
}

// Get request handler
func (mc *MicroserviceController) Get() {

	if Authenticate(mc.Ctx) != nil {
		// Redirect to login page
		mc.Redirect("/admin", http.StatusSeeOther)
		return
	}

	name := mc.Ctx.Input.Param(":name")

	mc.TplName = "page/microservice.tpl"
	mc.Data["Title"] = strings.Title(name) + " Microservice"
	mc.Data["Schema"] = Schema
}
