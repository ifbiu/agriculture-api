package yearbooks

import "github.com/astaxie/beego/orm"

type YearBooksAllResult struct {
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

func GetYearBooksAll(code int) ([]YearBooksAllResult, error) {
	o := orm.NewOrm()
	var yearBooksAllResult []YearBooksAllResult
	_, err := o.Raw("SELECT id,county,population,"+
		"gdp,gdp_incr,"+
		"cultivated_area,cultivated_area_incr,"+
		"farmland_area,farmland_area_incr,"+
		"sown_area,sown_area_incr,"+
		"grain_yield,grain_yield_incr,"+
		"oil_production,oil_production_incr "+
		"FROM yearbooks WHERE code=?", code).QueryRows(&yearBooksAllResult)
	if err != nil {
		return []YearBooksAllResult{}, err
	}
	yearBooksAllResult = fitterYearBooksAll(yearBooksAllResult)
	return yearBooksAllResult, nil
}

func fitterYearBooksAll(yearBooksResults []YearBooksAllResult) []YearBooksAllResult {
	for i, yearBooksResult := range yearBooksResults {
		str1 := yearBooksResult.OilProductionIncr
		if str1 == "" || str1 == "0" {
			yearBooksResults[i].OilProductionIncr = ""
			yearBooksResults[i].OilProductionIncrSign = 0
		} else if str1[0] == '-' {
			yearBooksResults[i].OilProductionIncr = str1[1:]
			yearBooksResults[i].OilProductionIncrSign = -1
		} else {
			yearBooksResults[i].OilProductionIncrSign = 1
		}

		str2 := yearBooksResult.OilProductionIncr
		if str2 == "" || str2 == "0" {
			yearBooksResults[i].GrainYieldIncr = ""
			yearBooksResults[i].GrainYieldIncrSign = 0
		} else if str2[0] == '-' {
			yearBooksResults[i].GrainYieldIncr = str2[1:]
			yearBooksResults[i].GrainYieldIncrSign = -1
		} else {
			yearBooksResults[i].GrainYieldIncrSign = 1
		}

		str3 := yearBooksResult.SownAreaIncr
		if str3 == "" || str3 == "0" {
			yearBooksResults[i].SownAreaIncr = ""
			yearBooksResults[i].SownAreaIncrSign = 0
		} else if str3[0] == '-' {
			yearBooksResults[i].SownAreaIncr = str3[1:]
			yearBooksResults[i].SownAreaIncrSign = -1
		} else {
			yearBooksResults[i].SownAreaIncrSign = 1
		}
	}
	return yearBooksResults
}
