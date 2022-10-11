package controllers

import (
	"agriculture-api/logic/yearbooks"
	"github.com/astaxie/beego"
)

type PopulationController struct {
	beego.Controller
}

func (this *PopulationController) Get() {
	var city = this.GetString("city")
	if city == "" {
		this.Data["json"] = map[string]string{
			"code":  "500",
			"error": "missing parameter city",
		}
		this.ServeJSON()
		return
	}
	if city == "neimenggu" {
		province, err := yearbooks.GetPopulationProvince(15)
		if err != nil {
			this.Data["json"] = map[string]string{
				"code":  "500",
				"error": "get yearbooks error",
			}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{
			"code": "200",
			"data": province,
		}
		this.ServeJSON()
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
	yearBooks, err := yearbooks.GetPopulation(code)
	if err != nil {
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
