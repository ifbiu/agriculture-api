package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

type CORSOptions struct {
	AllowDomain      []string // Used for allowing requests from custom domains
	AllowOrigin      string   // Access-Control-Allow-Origin
	AllowCredentials string   // Access-Control-Allow-Credentials
	ExposeHeaders    string   // Access-Control-Expose-Headers
	MaxAge           int      // Access-Control-Max-Age
	AllowMethods     string   // Access-Control-Allow-Methods
	AllowHeaders     string   // Access-Control-Allow-Headers
}

func CorsDomain(){
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:true,
		AllowMethods:[]string{"GET","POST","PUT","DELETE","OPTIONS"},
		AllowHeaders:[]string{"Origin","Authorization","Access-Control-Allow-Origin","Access-Control-Allow-Headers","Content-Type"},
		ExposeHeaders:[]string{"Content-Length","Access-Control-Allow-Origin","Access-Control-Allow-Headers","Content-Type"},
		AllowCredentials:true,
	}))
}