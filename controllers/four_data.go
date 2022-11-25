package controllers

import (
	"agriculture-api/logic/yearbooks"
	"github.com/astaxie/beego"
)

type FourDataController struct {
	beego.Controller
}

func (this *FourDataController) Get() {
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
	fourData, err := yearbooks.GetFourData(code)
	if err != nil {
		this.Data["json"] = map[string]string{
			"code":  "500",
			"error": "get fourData error",
		}
		this.ServeJSON()
		return
	}
	fourDataSum := yearbooks.GetFourDataSum(fourData)
	this.Data["json"] = map[string]interface{}{
		"code": "200",
		"data": fourDataSum,
	}
	this.ServeJSON()
}
