package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"git.urantiatech.com/cloudcms/cloudcms/api"
	"git.urantiatech.com/cloudcms/cloudcms/item"
	"github.com/urantiatech/beego"
	"golang.org/x/text/language"
)

// Save request handler
func (mc *ContentController) Save() {
	if Authenticate(mc.Ctx) != nil {
		// Redirect to login page
		mc.Redirect("/admin", http.StatusSeeOther)
		return
	}

	var err error
	var date time.Time
	flash := beego.NewFlash()

	name := mc.Ctx.Input.Param(":name")
	mc.Data["Name"] = name
	mc.Data["Title"] = strings.Title(name)
	mc.Data["Schema"] = Schema
	mc.Data["Flash"] = beego.ReadFromRequest(&mc.Controller).Data

	contents := make(map[string]interface{})
	for i, field := range strings.Split(mc.GetString("fields"), ",") {

		switch Schema[name].Fields[i].Widget {
		case item.WidgetDate:
			date, err = time.Parse("2006-01-02", mc.GetString(field))
			if err != nil {
				flash.Error(fmt.Sprintf("Error: %s", err.Error()))
				flash.Store(&mc.Controller)
				if slug := mc.GetString("slug"); slug != "" {
					mc.Redirect("/admin/content/"+name+"/editor?slug="+slug, http.StatusSeeOther)
				} else {
					mc.Redirect("/admin/content/"+name+"/editor", http.StatusSeeOther)
				}
				return
			}
			contents[field] = date.Format(time.RFC3339)

		case item.WidgetTags:
			var tags []string
			for _, tag := range strings.Split(mc.GetString(field), ",") {
				if t := strings.TrimSpace(tag); t != "" {
					tags = append(tags, t)
				}
			}
			contents[field] = tags

		default:
			contents[field] = mc.GetString(field)
		}
	}

	mc.Data["Header"] = item.Header{}

	if slug := mc.GetString("slug"); slug == "" {
		// Create Request
		slug = contents[mc.GetString("useforslug")].(string)
		_, err = api.Create(name, language.English.String(), slug, contents, os.Getenv("CLOUDCMS_SVC"))
	} else {
		// Update Request
		_, err = api.Update(name, mc.GetString("language"), slug, contents, os.Getenv("CLOUDCMS_SVC"))
	}
	if err != nil {
		flash.Error(fmt.Sprintf("Error: %s", err.Error()))
		flash.Store(&mc.Controller)
		if slug := mc.GetString("slug"); slug != "" {
			mc.Redirect("/admin/content/"+name+"/editor?slug="+slug, http.StatusSeeOther)
		} else {
			mc.Redirect("/admin/content/"+name+"/editor", http.StatusSeeOther)
		}
		return
	}

	flash.Notice(fmt.Sprintf("%s saved", strings.Title(name)))
	flash.Store(&mc.Controller)
	mc.Redirect("/admin/content/"+name, http.StatusSeeOther)
}
