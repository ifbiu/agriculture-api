package main

import (
	_ "agriculture-api/db"
	_ "agriculture-api/routers"
	"agriculture-api/utils"
	"github.com/astaxie/beego"
)

func init() {
	utils.CorsDomain()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
