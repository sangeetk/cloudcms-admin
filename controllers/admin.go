package controllers

import (
	"errors"
	"net/http"
	"os"

	"git.urantiatech.com/cloudcms/cloudcms/api"
	"github.com/urantiatech/beego"
	"github.com/urantiatech/beego/context"
	"golang.org/x/text/language"
)

// AdminController definition
type AdminController struct {
	beego.Controller
}

const authToken = "some-random-string"

// Schema stores definition of content types
var Schema map[string]api.ContentType

// Languages supported
var Languages []string

// Get request handler
func (ac *AdminController) Get() {
	if err := Authenticate(ac.Ctx); err != nil {
		ac.TplName = "page/login.tpl"
		ac.Data["Title"] = "Admin Login"
		ac.Data["Error"] = err.Error()
		return
	}

	lang := ac.GetString("lang")
	if lang == "" {
		lang = language.English.String()
	}

	// Set LanguageCode in Cookie
	signKey := beego.AppConfig.String("signkey")
	ac.SetSecureCookie(signKey, "LanguageCode", lang)

	if dst := ac.GetString("dst"); dst != "" {
		ac.Redirect(dst, http.StatusSeeOther)
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
		Languages, Schema, _ = api.Schema(os.Getenv("CLOUDCMS_SVC"))
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

// GetLanguage gets language code from Cookie
func GetLanguage(c *context.Context) string {
	// Check LanguageCode Cookie
	signkey := beego.AppConfig.String("signkey")
	lang, found := c.GetSecureCookie(signkey, "LanguageCode")
	if !found || lang == "" {
		return language.English.String()
	}
	return lang
}
