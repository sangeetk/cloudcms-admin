package main

import (
	"os"
	"strings"

	_ "git.urantiatech.com/cloudcms/cloudcms-admin/routers"
	"git.urantiatech.com/cloudcms/cloudcms-admin/views"
	"git.urantiatech.com/pkg/lang"
	"github.com/urantiatech/beego"
)

func main() {
	beego.AddFuncMap("appendField", views.AppendField)
	beego.AddFuncMap("timeToString", views.TimeToString)
	beego.AddFuncMap("unixTimeToString", views.UnixTimeToString)
	beego.AddFuncMap("timeToDateString", views.TimeToDate)
	beego.AddFuncMap("unixTimeToDateString", views.UnixTimeToDate)
	beego.AddFuncMap("title", strings.Title)
	beego.AddFuncMap("status", views.Status)
	beego.AddFuncMap("lowercase", strings.ToLower)
	beego.AddFuncMap("uppercase", strings.ToUpper)
	beego.AddFuncMap("hasPrefix", strings.HasPrefix)
	beego.AddFuncMap("contentTextValue", views.ContentTextValue)
	beego.AddFuncMap("contentDateValue", views.ContentDateValue)
	beego.AddFuncMap("contentTagsValue", views.ContentTagsValue)
	beego.AddFuncMap("contentFile", views.ContentFile)
	beego.AddFuncMap("currentDate", views.CurrentDate)
	beego.AddFuncMap("langCodeToName", lang.CodeToName)
	beego.AddFuncMap("langNameToCode", lang.NameToCode)
	beego.AddFuncMap("getenv", os.Getenv)
	beego.AddFuncMap("trimPrefix", strings.TrimPrefix)
	beego.AddFuncMap("drive", views.CloudDrive)

	beego.Run()
}
