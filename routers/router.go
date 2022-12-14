// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"agriculture-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/weather", &controllers.WeatherController{})
	beego.Router("/getYearBooks", &controllers.YearBooksController{})
	beego.Router("/getPopulation", &controllers.PopulationController{})
	beego.Router("/getFourData", &controllers.FourDataController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/getYearBooksAll", &controllers.YearBooksAllController{})
	beego.Router("/getDiffCounty", &controllers.DiffCountyController{})
}
