package controllers

import (
	"agriculture-api/logic/yearbooks"
	"agriculture-api/utils"
	"fmt"
	"github.com/astaxie/beego"
	"log"
)

type DiffCountyController struct {
	beego.Controller
}

func (this *DiffCountyController) Get() {
	defer this.ServeJSON()
	var county1 = this.GetString("county1")
	var county2 = this.GetString("county2")
	token := this.Ctx.Input.Header("Authorization")
	if token == "" {
		this.Data["json"] = map[string]string{
			"code":  "403",
			"error": "not find token !",
		}
		return
	}
	_, err := utils.ValidateToken(token)
	if err != nil {
		fmt.Println(err)
		this.Data["json"] = map[string]string{
			"code":  "403",
			"error": "token is err !",
		}
		return
	}
	fmt.Println(county1)
	fmt.Println(county2)
	if county1 == "" {
		this.Data["json"] = map[string]string{
			"code":  "500",
			"error": "missing parameter county1",
		}
		return
	}
	if county2 == "" {
		this.Data["json"] = map[string]string{
			"code":  "500",
			"error": "missing parameter county2",
		}
		return
	}
	diffData, err := yearbooks.GetDiffCounty(county1, county2)
	if err != nil {
		log.Println(err)
		this.Data["json"] = map[string]string{
			"code":  "500",
			"error": "get diffData error",
		}
		return
	}
	fmt.Println(diffData)
	this.Data["json"] = map[string]interface{}{
		"code": "200",
		"data": diffData,
	}
	return
}
