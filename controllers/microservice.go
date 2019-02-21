package controllers

import (
	"net/http"
	"os"
	"strings"

	"git.urantiatech.com/cloudcms/cloudcms/api"
	"git.urantiatech.com/cloudcms/cloudcms/item"
	"github.com/urantiatech/beego"
	"golang.org/x/text/language"
)

// MicroserviceController definition
type MicroserviceController struct {
	beego.Controller
}

// Index request handler
func (mc *MicroserviceController) Index() {

	if Authenticate(mc.Ctx) != nil {
		// Redirect to login page
		mc.Redirect("/admin", http.StatusSeeOther)
		return
	}

	name := mc.Ctx.Input.Param(":name")
	mc.Data["Name"] = name

	mc.TplName = "page/microservice.tpl"
	mc.Data["Title"] = strings.Title(name) + " Microservice"
	mc.Data["Schema"] = Schema

	list, total, err := api.List(name, language.English.String(), "", "id", 10, 0, os.Getenv("CLOUDCMS_SVC"))
	if err != nil {
		mc.Data["Error"] = err.Error()
		return
	}

	mc.Data["List"] = list
	mc.Data["Total"] = total

}

// Edit request handler
func (mc *MicroserviceController) Edit() {
	if Authenticate(mc.Ctx) != nil {
		// Redirect to login page
		mc.Redirect("/admin", http.StatusSeeOther)
		return
	}

	mc.TplName = "page/microservice-edit.tpl"
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

// Save request handler
func (mc *MicroserviceController) Save() {
	if Authenticate(mc.Ctx) != nil {
		// Redirect to login page
		mc.Redirect("/admin", http.StatusSeeOther)
		return
	}

	name := mc.Ctx.Input.Param(":name")
	mc.Data["Name"] = name
	mc.Data["Title"] = strings.Title(name) + " Microservice"
	mc.Data["Schema"] = Schema

	contents := make(map[string]interface{})
	for i, field := range strings.Split(mc.GetString("fields"), ",") {

		switch Schema[name].Fields[i].Widget {
		case item.WidgetInput:
			contents[field] = mc.GetString(field)

		case item.WidgetDate:
			contents[field] = mc.GetString(field)

		case item.WidgetFile:
			contents[field] = mc.GetString(field)

		case item.WidgetTextarea:
			contents[field] = mc.GetString(field)

		case item.WidgetRichtext:
			contents[field] = mc.GetString(field)

		case item.WidgetCheckbox:
			contents[field] = mc.GetString(field)

		case item.WidgetRadio:
			contents[field] = mc.GetString(field)

		case item.WidgetSelect:
			contents[field] = mc.GetString(field)

		case item.WidgetSelectMultiple:
			contents[field] = mc.GetString(field)

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

	var err error
	if slug := mc.GetString("slug"); slug == "" {
		// Create Request
		slug = contents[mc.GetString("useforslug")].(string)
		_, err = api.Create(name, language.English.String(), slug, contents, os.Getenv("CLOUDCMS_SVC"))
	} else {
		// Update Request
		_, err = api.Update(name, mc.GetString("language"), slug, contents, os.Getenv("CLOUDCMS_SVC"))
	}
	if err != nil {
		mc.Data["Error"] = err.Error()
		mc.TplName = "page/microservice-edit.tpl"
		return
	}

	mc.Redirect("/admin/microservice/"+name, http.StatusSeeOther)
}
