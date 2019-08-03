package controllers

import (
	"net/http"
	"strings"

	"git.urantiatech.com/cloudcms/cloudcms/api"
	"git.urantiatech.com/pkg/lang"
	"github.com/urantiatech/beego"
)

// ContentController definition
type ContentController struct {
	beego.Controller
}

// Index request handler
func (mc *ContentController) Index() {
	mc.Data["Languages"] = Languages
	mc.Data["LanguageCode"] = GetLanguage(mc.Ctx)
	mc.Data["URI"] = mc.Ctx.Request.URL.String()

	if Authenticate(mc.Ctx) != nil {
		// Redirect to login page
		mc.Redirect("/admin", http.StatusSeeOther)
		return
	}

	name := mc.Ctx.Input.Param(":name")
	mc.Data["Name"] = name
	mc.Data["Flash"] = beego.ReadFromRequest(&mc.Controller).Data

	mc.TplName = "page/content-index.tpl"
	mc.Data["Title"] = lang.CodeToName(GetLanguage(mc.Ctx)) + " " + strings.Title(name)
	mc.Data["Schema"] = Schema

	list, total, err := api.List(name, GetLanguage(mc.Ctx), "", "-id", 10, 0, CurrentCMS)
	if err != nil {
		mc.Data["Error"] = err.Error()
		return
	}

	mc.Data["List"] = list
	mc.Data["Total"] = total
}
