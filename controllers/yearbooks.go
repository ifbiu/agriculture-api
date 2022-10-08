package controllers

import "github.com/astaxie/beego"

type YearBooksController struct {
	beego.Controller
}

func (this *YearBooksController) Get() {
	this.Data["json"] = map[string]string{
		"msg": "yearbooks",
	}
	this.ServeJSON()
}
