package controllers

import (
	"net/http"
	"os"

	"git.urantiatech.com/cloudcms/cloudcms/api"
	"github.com/urantiatech/beego"
)

// DashboardController definition
type DashboardController struct {
	beego.Controller
}

// Get request handler
func (dc *DashboardController) Get() {
	if Authenticate(dc.Ctx) != nil {
		// Redirect to login page
		dc.Redirect("/admin", http.StatusSeeOther)
		return
	}

	if Schema == nil {
		Languages, Schema, _ = api.Schema(os.Getenv("CLOUDCMS_SVC"))
	}

	dc.TplName = "page/dashboard.tpl"
	dc.Data["Title"] = "Dashboard"
	dc.Data["Schema"] = Schema
	dc.Data["Languages"] = Languages
	dc.Data["LanguageCode"] = GetLanguage(dc.Ctx)
	dc.Data["URI"] = dc.Ctx.Request.URL.String()
}
