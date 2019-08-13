package controllers

import (
	"net/http"
	"strings"

	"git.urantiatech.com/cloudcms/cloudcms/api"
	"git.urantiatech.com/pkg/lang"
	"github.com/astaxie/beego"
	"golang.org/x/text/language"
)

// Editor request handler
func (mc *ContentController) Editor() {
	mc.Data["Languages"] = Languages
	mc.Data["LanguageCode"] = GetLanguage(mc.Ctx)
	mc.Data["URI"] = mc.Ctx.Request.URL.String()
	mc.Data["CurrentCMS"] = CurrentCMS

	if Authenticate(mc.Ctx) != nil {
		// Redirect to login page
		mc.Redirect("/admin", http.StatusSeeOther)
		return
	}

	mc.TplName = "page/content-editor.tpl"
	name := mc.Ctx.Input.Param(":name")
	mc.Data["Name"] = name
	mc.Data["Flash"] = beego.ReadFromRequest(&mc.Controller).Data
	slug := mc.GetString("slug")

	if slug == "" {
		// New
		mc.Data["Title"] = "Add " + " " + strings.Title(name)
		mc.Data["SubmitButton"] = "Add " + " " + strings.Title(name)
	} else {
		c, _ := api.Read(name, GetLanguage(mc.Ctx), slug, CurrentCMS)
		if c != nil {
			// Edit
			mc.Data["Content"] = c.(map[string]interface{})
			mc.Data["Title"] = "Edit " + " " + strings.Title(name)
			mc.Data["SubmitButton"] = "Update " + " " + strings.Title(name)
		} else {
			// Trsnalation
			c, _ = api.Read(name, language.English.String(), slug, CurrentCMS)
			if c != nil {
				enContent := c.(map[string]interface{})
				enContent["language"] = GetLanguage(mc.Ctx)
				delete(enContent, "id")
				delete(enContent, "status")
				delete(enContent, "created_at")
				delete(enContent, "updated_at")
				// delete(enContent, "slug")
				mc.Data["Content"] = enContent
			}
			mc.Data["Title"] = "Add " + lang.CodeToName(GetLanguage(mc.Ctx)) + " Translation"
			mc.Data["TranslationSlug"] = slug
			mc.Data["SubmitButton"] = "Add " + lang.CodeToName(GetLanguage(mc.Ctx)) + " Translation"
		}
	}

	mc.Data["Schema"] = Schema
	mc.Data["Fields"] = Schema[name].Fields
}
