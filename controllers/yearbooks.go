package controllers

import (
	"agriculture-api/logic/yearbooks"
	"github.com/astaxie/beego"
	"log"
	"regexp"
)

type YearBooksController struct {
	beego.Controller
}

func (this *YearBooksController) Get() {
	re := regexp.MustCompile(`[/][^?]*`)
	uri := re.FindAllString(this.Ctx.Request.RequestURI, -1)[0]
	defer this.ServeJSON()
	var city = this.GetString("city")
	if city == "" {
		this.Data["json"] = map[string]string{
			"code":  "500",
			"error": "missing parameter city",
		}
		return
	}
	code := yearbooks.SwitchChangeCity(city)
	if code == 0 {
		this.Data["json"] = map[string]string{
			"code":  "500",
			"error": "city is not found",
		}
		return
	}
	if uri == "/getYearBooks" {
		yearBooks, err := yearbooks.GetYearBooks(code)
		if err != nil {
			log.Println(err)
			this.Data["json"] = map[string]string{
				"code":  "500",
				"error": "get yearbooks error",
			}
			return
		}
		this.Data["json"] = map[string]interface{}{
			"code": "200",
			"data": yearBooks,
		}
		return
	} else if uri == "/getYearBooksAll" {
		yearBooks, err := yearbooks.GetYearBooksAll(code)
		if err != nil {
			log.Println(err)
			this.Data["json"] = map[string]string{
				"code":  "500",
				"error": "get yearbooks error",
			}
			return
		}
		this.Data["json"] = map[string]interface{}{
			"code": "200",
			"data": yearBooks,
		}
		return
	}
}
