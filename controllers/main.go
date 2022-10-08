package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["json"] = map[string]string{
		"Welcome to": "agriculture-api~~",
		"Author":     "Candide",
	}
	this.ServeJSON()
}
