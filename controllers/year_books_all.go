package controllers

import (
	"agriculture-api/logic/yearbooks"
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
)

type YearBooksAllController struct {
	beego.Controller
}

type PostResponse struct {
	Data struct {
		Id                      int    `json:"id"`
		County                  string `json:"county"`
		Population              string `json:"population"`
		Gdp                     string `json:"gdp"`
		GdpIncr                 string `json:"gdp_incr"`
		CultivatedArea          int    `json:"cultivated_area"`
		CultivatedAreaIncr      string `json:"cultivated_area_incr"`
		FarmlandArea            int    `json:"farmland_area"`
		FarmlandAreaIncr        string `json:"farmland_area_incr"`
		SownArea                int
		SownArea_incr           string
		SownArea_incr_sign      int
		GrainYield              int
		GrainYield_incr         string
		GrainYield_incr_sign    int
		OilProduction           int
		OilProduction_incr      string
		OilProduction_incr_sign int
	} `json:"data"`
}

func (this *YearBooksAllController) Get() {
	defer this.ServeJSON()
	var city = this.GetString("city")
	//token := this.Ctx.Input.Header("Authorization")
	//if token == "" {
	//	this.Data["json"] = map[string]string{
	//		"code":  "403",
	//		"error": "not find token !",
	//	}
	//	return
	//}
	//info, err := utils.ValidateToken(token)
	//if err != nil {
	//	fmt.Println(err)
	//	this.Data["json"] = map[string]string{
	//		"code":  "403",
	//		"error": "token is err !",
	//	}
	//	return
	//}
	//fmt.Println(info)
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

func (this *YearBooksAllController) Post() {
	defer this.ServeJSON()
	var postResponse yearbooks.YarBooksResponse
	data := this.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &postResponse)
	if err != nil {
		this.Data["json"] = map[string]string{
			"code":  "500",
			"error": "json.Unmarshal is err !",
		}
	}
	yearBooks, err := yearbooks.UpdateYearBooks(postResponse)
	if err != nil {
		log.Println(err)
		this.Data["json"] = map[string]string{
			"code":  "500",
			"error": "update yearbooks error",
		}
		return
	}
	this.Data["json"] = map[string]interface{}{
		"code": "200",
		"data": yearBooks,
	}
	return
}
