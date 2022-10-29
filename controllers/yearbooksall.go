package controllers

import (
	"agriculture-api/logic/yearbooks"
	"agriculture-api/utils"
	"fmt"
	"github.com/astaxie/beego"
	"log"
)

type YearBooksAllController struct {
	beego.Controller
}

func (this *YearBooksAllController) Get() {
	defer this.ServeJSON()
	var city = this.GetString("city")
	token := this.Ctx.Input.Header("Authorization")
	if token == "" {
		this.Data["json"] = map[string]string{
			"code":  "500",
			"error": "not find token !",
		}
		return
	}
	info, err := utils.ValidateToken(token)
	if err != nil {
		fmt.Println(err)
		this.Data["json"] = map[string]string{
			"code":  "500",
			"error": "token is err !",
		}
		return
	}
	fmt.Println(info)
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
