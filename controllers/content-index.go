package controllers

import (
	"net/http"
	"os"
	"strings"

	"git.urantiatech.com/cloudcms/cloudcms/api"
	"github.com/urantiatech/beego"
	"golang.org/x/text/language"
)

// ContentController definition
type ContentController struct {
	beego.Controller
}

// Index request handler
func (mc *ContentController) Index() {
	if Authenticate(mc.Ctx) != nil {
		// Redirect to login page
		mc.Redirect("/admin", http.StatusSeeOther)
		return
	}

	name := mc.Ctx.Input.Param(":name")
	mc.Data["Name"] = name
	mc.Data["Flash"] = beego.ReadFromRequest(&mc.Controller).Data

	mc.TplName = "page/content-index.tpl"
	mc.Data["Title"] = strings.Title(name)
	mc.Data["Schema"] = Schema

	list, total, err := api.List(name, language.English.String(), "", "id", 10, 0, os.Getenv("CLOUDCMS_SVC"))
	if err != nil {
		mc.Data["Error"] = err.Error()
		return
	}

	mc.Data["List"] = list
	mc.Data["Total"] = total
}
