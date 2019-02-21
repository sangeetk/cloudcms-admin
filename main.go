package main

import (
	"strings"

	_ "git.urantiatech.com/cloudcms/cloudcms-admin/routers"
	"git.urantiatech.com/cloudcms/cloudcms-admin/views"
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
	beego.AddFuncMap("hasPrefix", strings.HasPrefix)
	beego.AddFuncMap("contentTextValue", views.ContentTextValue)
	beego.AddFuncMap("contentDateValue", views.ContentDateValue)
	beego.AddFuncMap("contentTagsValue", views.ContentTagsValue)

	beego.Run()
}
