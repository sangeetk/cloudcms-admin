package controllers

import (
	"net/http"
	"os"
	"strings"

	"git.urantiatech.com/cloudcms/cloudcms/api"
	"golang.org/x/text/language"
)

// Editor request handler
func (mc *ContentController) Editor() {
	if Authenticate(mc.Ctx) != nil {
		// Redirect to login page
		mc.Redirect("/admin", http.StatusSeeOther)
		return
	}

	mc.TplName = "page/content-editor.tpl"
	name := mc.Ctx.Input.Param(":name")
	mc.Data["Name"] = name

	slug := mc.GetString("slug")
	var c interface{}
	var err error

	if slug == "" {
		mc.Data["Title"] = "Add " + strings.Title(name)
	} else {
		mc.Data["Title"] = "Edit " + strings.Title(name)
		c, err = api.Read(name, language.English.String(), slug, os.Getenv("CLOUDCMS_SVC"))
	}
	if err != nil {
		mc.Data["Error"] = err.Error()
		return
	}
	if c != nil {
		mc.Data["Content"] = c.(map[string]interface{})
	}

	mc.Data["Schema"] = Schema
	mc.Data["Fields"] = Schema[name].Fields
}
