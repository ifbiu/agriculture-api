package yearbooks

import (
	"github.com/astaxie/beego/orm"
)

type CountyData struct {
	Id                    int    `json:"id"`
	County                string `json:"county"`
	Population            string `json:"population"`
	Gdp                   string `json:"gdp"`
	GdpIncr               string `json:"gdp_incr"`
	CultivatedArea        int    `json:"cultivated_area"`
	CultivatedAreaIncr    string `json:"cultivated_area_incr"`
	FarmlandArea          int    `json:"farmland_area"`
	FarmlandArea_incr     string `json:"farmland_area_incr"`
	SownArea              int    `json:"sown_area"`
	SownAreaIncr          string `json:"sown_area_incr"`
	SownAreaIncrSign      int    `json:"sown_area_incr_sign"`
	GrainYield            int    `json:"grain_yield"`
	GrainYieldIncr        string `json:"grain_yield_incr"`
	GrainYieldIncrSign    int    `json:"grain_yield_incr_sign"`
	OilProduction         int    `json:"oil_production"`
	OilProductionIncr     string `json:"oil_production_incr"`
	OilProductionIncrSign int    `json:"oil_production_incr_sign"`
}

type DiffCountyResult struct {
	CountyData1 CountyData
	CountyData2 CountyData
}

func GetDiffCounty(county1 string, county2 string) (DiffCountyResult, error) {

	o := orm.NewOrm()
	var diffCountyResult DiffCountyResult
	var countyData1 CountyData
	var countyData2 CountyData
	err := o.Raw("SELECT id,county,population,"+
		"gdp,gdp_incr,"+
		"cultivated_area,cultivated_area_incr,"+
		"farmland_area,farmland_area_incr,"+
		"sown_area,sown_area_incr,"+
		"grain_yield,grain_yield_incr,"+
		"oil_production,oil_production_incr "+
		"FROM yearbooks WHERE ccode=?", county1).QueryRow(&countyData1)
	if err != nil {
		return DiffCountyResult{}, err
	}
	err = o.Raw("SELECT id,county,population,"+
		"gdp,gdp_incr,"+
		"cultivated_area,cultivated_area_incr,"+
		"farmland_area,farmland_area_incr,"+
		"sown_area,sown_area_incr,"+
		"grain_yield,grain_yield_incr,"+
		"oil_production,oil_production_incr "+
		"FROM yearbooks WHERE ccode=?", county2).QueryRow(&countyData2)
	if err != nil {
		return DiffCountyResult{}, err
	}
	diffCountyResult.CountyData1 = countyData1
	diffCountyResult.CountyData2 = countyData2
	return diffCountyResult, nil
}
