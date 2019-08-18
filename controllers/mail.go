package controllers

import (
	"net/http"

	"git.urantiatech.com/cloudcms/cloudcms/api"
	"github.com/astaxie/beego"
)

// MailController definition
type MailController struct {
	beego.Controller
}

// Get request handler
func (mc *MailController) Get() {
	if Authenticate(mc.Ctx) != nil {
		// Redirect to login page
		mc.Redirect("/admin", http.StatusSeeOther)
		return
	}

	if Schema == nil {
		Languages, Schema, _ = api.Schema(CurrentCMS)
	}

	mc.TplName = "page/mail.tpl"
	mc.Data["Title"] = "Mail"
	mc.Data["Schema"] = Schema
	mc.Data["Languages"] = Languages
	mc.Data["LanguageCode"] = GetLanguage(mc.Ctx)
	mc.Data["URI"] = mc.Ctx.Request.URL.String()
	mc.Data["Name"] = "mail"
}
