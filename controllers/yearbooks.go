package controllers

import (
	"agriculture-api/logic/yearbooks"
	"github.com/astaxie/beego"
	"log"
)

type YearBooksController struct {
	beego.Controller
}

func (this *YearBooksController) Get() {
	var city = this.GetString("city")
	if city == "" {
		this.Data["json"] = map[string]string{
			"code":  "500",
			"error": "missing parameter city",
		}
		this.ServeJSON()
		return
	}
	code := yearbooks.SwitchChangeCity(city)
	if code == 0 {
		this.Data["json"] = map[string]string{
			"code":  "500",
			"error": "city is not found",
		}
		this.ServeJSON()
		return
	}
	yearBooks, err := yearbooks.GetYearBooks(code)
	if err != nil {
		log.Println(err)
		this.Data["json"] = map[string]string{
			"code":  "500",
			"error": "get yearbooks error",
		}
		this.ServeJSON()
		return
	}
	this.Data["json"] = map[string]interface{}{
		"code": "200",
		"data": yearBooks,
	}
	this.ServeJSON()
}
