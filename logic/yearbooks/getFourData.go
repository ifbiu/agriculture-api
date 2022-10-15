package yearbooks

import "github.com/astaxie/beego/orm"

type FourDataResult struct {
	Id             int `json:"id"`
	CultivatedArea int `json:"cultivated_area"`
	SownArea       int `json:"sown_area"`
	GrainYield     int `json:"grain_yield"`
	OilProduction  int `json:"oil_production"`
}

type FourDataSumResult struct {
	Data              []FourDataResult `json:"data"`
	CultivatedAreaSum int              `json:"cultivated_area_sum"`
	SownAreaSum       int              `json:"sown_area_sum"`
	GrainYieldSum     int              `json:"grain_yield_sum"`
	OilProductionSum  int              `json:"oil_production_sum"`
}

func GetFourData(code int) ([]FourDataResult, error) {
	o := orm.NewOrm()
	var fourDataResult []FourDataResult
	_, err := o.Raw("SELECT id,cultivated_area,sown_area,grain_yield,oil_production FROM yearbooks WHERE code=?", code).QueryRows(&fourDataResult)
	if err != nil {
		return []FourDataResult{}, err
	}
	return fourDataResult, nil
}

func GetFourDataSum(fourDataAll []FourDataResult) FourDataSumResult {
	if len(fourDataAll) == 0 {
		return FourDataSumResult{}
	}
	fourDataSum := FourDataSumResult{}
	fourDataSum.Data = fourDataAll
	for i := 0; i < len(fourDataAll); i++ {
		fourDataSum.SownAreaSum += fourDataAll[i].SownArea
		fourDataSum.GrainYieldSum += fourDataAll[i].GrainYield
		fourDataSum.OilProductionSum += fourDataAll[i].OilProduction
		fourDataSum.CultivatedAreaSum += fourDataAll[i].CultivatedArea
	}
	return fourDataSum
}
