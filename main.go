package main

import (
	_ "git.urantiatech.com/cloudcms/cloudcms-admin/routers"
	"git.urantiatech.com/cloudcms/cloudcms-admin/views"
	"github.com/urantiatech/beego"
)

func main() {
	beego.AddFuncMap("title", views.Title)
	beego.Run()
}
