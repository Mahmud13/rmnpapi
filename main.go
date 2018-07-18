package main

import (
	"mnp_api/database"
	_ "mnp_api/routers"
	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	db := database.DB
	defer db.Close()
	beego.Run()
}
