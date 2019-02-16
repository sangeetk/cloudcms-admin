package controllers

import (
	"errors"
	"net/http"
	"os"

	"git.urantiatech.com/cloudcms/cloudcms/api"
	"github.com/urantiatech/beego"
	"github.com/urantiatech/beego/context"
)

// AdminController definition
type AdminController struct {
	beego.Controller
}

const authToken = "some-random-string"

// Schema stores definition of content types
var Schema map[string]api.ContentType

// Get request handler
func (ac *AdminController) Get() {
	if err := Authenticate(ac.Ctx); err != nil {
		ac.TplName = "page/login.tpl"
		ac.Data["Title"] = "Admin Login"
		ac.Data["Error"] = err.Error()
		return
	}

	// Redirect to Dashboard if already logged in
	ac.Redirect("/admin/dashboard", http.StatusSeeOther)
}

// Post request handler
func (ac *AdminController) Post() {
	username := ac.GetString("username")
	password := ac.GetString("password")

	if username != beego.AppConfig.String("adminuser") ||
		password != beego.AppConfig.String("adminpass") {
		ac.TplName = "page/login.tpl"
		ac.Data["Title"] = "Admin Login"
		ac.Data["Error"] = "Invalid username or password"
		return
	}

	// Set Auth Cookie
	signkey := beego.AppConfig.String("signkey")
	ac.SetSecureCookie(signkey, "AuthCookie", authToken)

	if Schema == nil {
		Schema, _ = api.Schema(os.Getenv("CLOUDCMS_SVC"))
	}

	// Redirect to Dashboard after login
	ac.Redirect("/admin/dashboard", http.StatusSeeOther)
}

// Authenticate looks for AuthCookie
func Authenticate(c *context.Context) error {
	// Check Auth Cookie
	signkey := beego.AppConfig.String("signkey")
	val, found := c.GetSecureCookie(signkey, "AuthCookie")
	if !found || val != authToken {
		return errors.New("Please login to continue")
	}
	return nil
}
