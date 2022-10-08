package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type WeatherController struct {
	beego.Controller
}

type Daily struct {
	TempMax       string `json:"tempMax"`
	TempMin       string `json:"tempMin"`
	IconDay       string `json:"iconDay"`
	TextDay       string `json:"textDay"`
	FxDate        string `json:"fxDate"`
	Weekday       int    `json:"weekday"`
	WeekdayFormat string `json:"weekdayFormat"`
}

type weatherResultStruct struct {
	Code  string  `json:"code"`
	Daily []Daily `json:"daily"`
}

func (this *WeatherController) Get() {
	var key = "20b4d42efc414cd79fd73b7840604190"
	var city = this.GetString("city")
	if city == "" {
		this.Data["json"] = map[string]string{
			"code":  "500",
			"error": "missing parameter city",
		}
		this.ServeJSON()
		return
	}
	location := switchCityToLocation(city)
	if city == "nofound" {
		this.Data["json"] = map[string]string{
			"code":  "500",
			"error": "city is not found",
		}
		this.ServeJSON()
		return
	}
	response, err := http.Get("https://devapi.qweather.com/v7/weather/7d?location=" + location + "&key=" + key)
	if err != nil {
		log.Println("http get error")
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("ioutil close error")
			return
		}
	}(response.Body)
	body, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		log.Println("ioutil read error")
		return
	}
	var weatherData = weatherResultStruct{}
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		log.Println("json transformation error")
		return
	}
	if weatherData.Code != "200" {
		this.Data["json"] = map[string]string{
			"code": weatherData.Code,
			"err":  "服务异常！",
		}
		this.ServeJSON()
		return
	}
	t := time.Now()
	weekday := int(t.Weekday())
	for i, daily := range weatherData.Daily {
		fxDateAll := strings.Split(daily.FxDate, "-")
		if len(fxDateAll) != 3 {
			log.Println("fxDate split error")
			return
		}
		if fxDateAll[2][0] == '0' {
			fxDateAll[2] = fxDateAll[2][1:]
		}
		weatherData.Daily[i].FxDate = fxDateAll[2] + "日"
		// 获取周几
		weatherData.Daily[i].Weekday = weekday
		switch weekday {
		case 0:
			weatherData.Daily[i].WeekdayFormat = "周日"
		case 1:
			weatherData.Daily[i].WeekdayFormat = "周一"
		case 2:
			weatherData.Daily[i].WeekdayFormat = "周二"
		case 3:
			weatherData.Daily[i].WeekdayFormat = "周三"
		case 4:
			weatherData.Daily[i].WeekdayFormat = "周四"
		case 5:
			weatherData.Daily[i].WeekdayFormat = "周五"
		case 6:
			weatherData.Daily[i].WeekdayFormat = "周六"

		}
		// 星期加加
		if weekday+1 > 6 {
			weekday = 0
		} else {
			weekday = weekday + 1
		}
	}
	this.Data["json"] = weatherData
	this.ServeJSON()
}

func switchCityToLocation(city string) string {
	var res = "nofound"
	switch city {
	case "huhehaote":
		res = "101080101"
	case "baotou":
		res = "101080201"
	case "wulanchabu":
		res = "101080405"
	case "tongliao":
		res = "101080501"
	case "chifeng":
		res = "101080601"
	case "eerduosi":
		res = "101080701"
	case "bayannaoer":
		res = "101080811"
	case "xilinguole":
		res = "101080902"
	case "hulunbeier":
		res = "101081013"
	case "xinganmeng":
		res = "101081108"
	case "alashan":
		res = "101081213"
	case "wuhai":
		res = "101080301"
	}
	return res
}
