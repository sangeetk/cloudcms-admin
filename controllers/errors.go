package controllers

import (
	"github.com/astaxie/beego"
)

// ErrorController handles 404, 401, 403, 500, 503 errors
type ErrorController struct {
	beego.Controller
}

// Error404 handles 404
func (ec *ErrorController) Error404() {
	ec.Data["content"] = "page not found"
	ec.TplName = "page/404.tpl"
}

// Error500 handles 500
func (ec *ErrorController) Error500() {
	ec.Data["content"] = "internal server error"
	ec.TplName = "page/500.tpl"
}
