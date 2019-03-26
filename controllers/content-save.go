package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"git.urantiatech.com/cloudcms/cloudcms/api"
	"git.urantiatech.com/cloudcms/cloudcms/item"
	"github.com/urantiatech/beego"
)

// Save request handler
func (mc *ContentController) Save() {
	mc.Data["Languages"] = Languages
	mc.Data["LanguageCode"] = GetLanguage(mc.Ctx)
	mc.Data["URI"] = mc.Ctx.Request.URL.String()

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

		case item.WidgetFile:
			file, header, err := mc.GetFile(field)
			if err != nil {
				// Use file from English version if available
				if mc.GetString(field+".name") != "" {
					enFile := item.File{
						Name: mc.GetString(field + ".name"),
						Type: mc.GetString(field + ".type"),
						URI:  mc.GetString(field + ".uri"),
					}
					enFile.Size, _ = mc.GetInt64(field + ".size")
					contents[field] = enFile
				}
				err = nil
			} else {
				dst := item.File{
					Name: header.Filename,
					Type: Schema[name].Fields[i].FileType,
					Size: header.Size,
				}
				log.Println(dst)
				// copy the uploaded file to the destination file
				if _, err := io.Copy(&dst, file); err != nil {
					flash.Error(fmt.Sprintf("Error: %s", err.Error()))
					flash.Store(&mc.Controller)
					if slug := mc.GetString("slug"); slug != "" {
						mc.Redirect("/admin/content/"+name+"/editor?slug="+slug, http.StatusSeeOther)
					} else {
						mc.Redirect("/admin/content/"+name+"/editor", http.StatusSeeOther)
					}
					return
				}
				contents[field] = dst
			}

		default:
			contents[field] = mc.GetString(field)
		}
	}

	mc.Data["Header"] = item.Header{}

	slug := mc.GetString("slug")
	translationslug := mc.GetString("translationslug")

	if slug == "" && translationslug == "" {
		// Create Request
		slug = contents[mc.GetString("useforslug")].(string)
		_, err = api.Create(name, GetLanguage(mc.Ctx), slug, contents, CurrentCMS)
	} else if slug == "" && translationslug != "" {
		// New Translation Request
		_, err = api.Create(name, mc.GetString("language"), translationslug, contents, CurrentCMS)
	} else {
		// Update Reuest
		_, err = api.Update(name, mc.GetString("language"), slug, contents, CurrentCMS)
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

	mc.Redirect("/admin/content/"+name, http.StatusSeeOther)
}
