package controllers

import (
	"net/http"
	"strings"

	"git.urantiatech.com/cloudcms/cloudcms/api"
	"github.com/astaxie/beego"
)

// ContentController definition
type ContentController struct {
	beego.Controller
}

// Index request handler
func (mc *ContentController) Index() {
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
	mc.Data["Flash"] = beego.ReadFromRequest(&mc.Controller).Data

	mc.TplName = "page/content-index.tpl"
	mc.Data["Title"] = strings.Title(name)
	mc.Data["Schema"] = Schema

	query := mc.GetString("q")
	mc.Data["Query"] = strings.TrimSpace(query)

	p, err := mc.GetInt("p")
	if err != nil || p <= 0 {
		p = 1
	}

	size, err := beego.AppConfig.Int("itemsperpage")
	if err != nil {
		size = 25
	}
	skip := (p - 1) * size

	mc.Data["Size"] = size
	mc.Data["Skip"] = skip

	if query == "" {
		list, total, err := api.List(name, GetLanguage(mc.Ctx), "", "-id", size, skip, CurrentCMS)
		if err != nil {
			mc.Data["Error"] = err.Error()
			return
		}
		mc.Data["List"] = list
		mc.Data["First"] = first(size, skip, int(total))
		mc.Data["Last"] = last(size, skip, int(total))
		mc.Data["CurrentPage"] = p
		mc.Data["Size"] = size
		mc.Data["Total"] = int(total)
		return
	}

	results, total, _, err := api.Search(name, GetLanguage(mc.Ctx), query, false, size, skip, CurrentCMS)
	if err != nil {
		mc.Data["Error"] = err.Error()
		return
	}
	mc.Data["List"] = results
	mc.Data["First"] = first(size, skip, int(total))
	mc.Data["Last"] = last(size, skip, int(total))
	mc.Data["CurrentPage"] = p
	mc.Data["Size"] = size
	mc.Data["Total"] = int(total)
}

func first(size, skip, total int) int {
	if total == 0 {
		return 0
	}
	return skip + 1
}

func last(size, skip, total int) int {
	if total == 0 {
		return 0
	}

	if total > skip+size {
		return skip + size
	}

	return total
}
