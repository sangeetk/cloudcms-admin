package controllers

import (
	"net/http"
	"os"

	"git.urantiatech.com/cloudcms/cloudcms/api"
	"golang.org/x/text/language"
)

// Delete request handler
func (mc *ContentController) Delete() {
	if Authenticate(mc.Ctx) != nil {
		// Redirect to login page
		mc.Redirect("/admin", http.StatusSeeOther)
		return
	}

	name := mc.Ctx.Input.Param(":name")
	mc.Data["Name"] = name

	slug := mc.GetString("slug")

	if slug == "" {
		mc.Redirect("/admin/content/"+name, http.StatusSeeOther)
		return
	}

	_, err := api.Delete(name, language.English.String(), slug, os.Getenv("CLOUDCMS_SVC"))
	if err != nil {
		mc.Data["Error"] = err.Error()
	}
	mc.Redirect("/admin/content/"+name, http.StatusSeeOther)
}
