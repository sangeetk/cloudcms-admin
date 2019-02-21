package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"git.urantiatech.com/cloudcms/cloudcms/api"
	"github.com/urantiatech/beego"
	"golang.org/x/text/language"
)

// Delete request handler
func (mc *ContentController) Delete() {
	if Authenticate(mc.Ctx) != nil {
		// Redirect to login page
		mc.Redirect("/admin", http.StatusSeeOther)
		return
	}

	flash := beego.NewFlash()

	name := mc.Ctx.Input.Param(":name")
	mc.Data["Name"] = name

	slug := mc.GetString("slug")
	mc.Data["Flash"] = beego.ReadFromRequest(&mc.Controller).Data

	if slug == "" {
		mc.Redirect("/admin/content/"+name, http.StatusSeeOther)
		return
	}

	_, err := api.Delete(name, language.English.String(), slug, os.Getenv("CLOUDCMS_SVC"))
	if err != nil {
		mc.Data["Error"] = err.Error()
	}
	flash.Notice(fmt.Sprintf("%s deleted", strings.Title(name)))
	flash.Store(&mc.Controller)
	mc.Redirect("/admin/content/"+name, http.StatusSeeOther)
}
