package controllers

import (
	"net/http"

	"github.com/urantiatech/beego"
)

// LogoutController definition
type LogoutController struct {
	beego.Controller
}

// Get request handler
func (lc *LogoutController) Get() {
	// Set Empty Auth Cookie
	signkey := beego.AppConfig.String("signkey")
	lc.SetSecureCookie(signkey, "AuthCookie", "")

	// Redirect to login page
	lc.Redirect("/admin", http.StatusSeeOther)
}
