package controllers

import (
	"log"
	"net/http"

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

	dc.TplName = "page/dashboard.tpl"
	dc.Data["Title"] = "Dashboard"
}
