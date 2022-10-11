package yearbooks

import "github.com/astaxie/beego/orm"

type PopulationResult struct {
	Id         int    `json:"id"`
	County     string `json:"county"`
	Population int    `json:"population"`
}

type PopulationProvinceResult struct {
	Id         int    `json:"id"`
	City       string `json:"city"`
	Population int    `json:"population"`
}

func GetPopulationProvince(province int) ([]PopulationProvinceResult, error) {
	o := orm.NewOrm()
	var yearBooksResults []PopulationProvinceResult
	_, err := o.Raw("SELECT id,city,sum(population) as population FROM yearbooks WHERE province=? group by code", province).QueryRows(&yearBooksResults)
	if err != nil {
		return []PopulationProvinceResult{}, err
	}
	return yearBooksResults, nil
}

func GetPopulation(code int) ([]PopulationResult, error) {
	o := orm.NewOrm()
	var yearBooksResults []PopulationResult
	_, err := o.Raw("SELECT id,county,population FROM yearbooks WHERE code=?", code).QueryRows(&yearBooksResults)
	if err != nil {
		return []PopulationResult{}, err
	}
	return yearBooksResults, nil
}
