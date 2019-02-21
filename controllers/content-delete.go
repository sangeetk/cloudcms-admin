package controllers

import (
	"fmt"
	"net/http"
	"os"

	"git.urantiatech.com/cloudcms/cloudcms/api"
	"github.com/urantiatech/beego"
)

// Delete request handler
func (mc *ContentController) Delete() {
	mc.Data["Languages"] = Languages
	mc.Data["LanguageCode"] = GetLanguage(mc.Ctx)
	mc.Data["URI"] = mc.Ctx.Request.URL.String()

	if Authenticate(mc.Ctx) != nil {
		// Redirect to login page
		mc.Redirect("/admin", http.StatusSeeOther)
		return
	}

	name := mc.Ctx.Input.Param(":name")
	mc.Data["Name"] = name

	slug := mc.GetString("slug")
	mc.Data["Flash"] = beego.ReadFromRequest(&mc.Controller).Data

	if slug == "" {
		mc.Redirect("/admin/content/"+name, http.StatusSeeOther)
		return
	}

	_, err := api.Delete(name, GetLanguage(mc.Ctx), slug, os.Getenv("CLOUDCMS_SVC"))
	if err != nil {
		flash := beego.NewFlash()
		flash.Error(fmt.Sprintf("Error: %s", err.Error()))
		flash.Store(&mc.Controller)
	}

	mc.Redirect("/admin/content/"+name, http.StatusSeeOther)
}
