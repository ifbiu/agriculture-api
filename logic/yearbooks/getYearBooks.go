package yearbooks

import "github.com/astaxie/beego/orm"

type YearBooksResult struct {
	Id      int    `json:"id"`
	County  string `json:"county"`
	Gdp     string `json:"gdp"`
	GdpIncr string `json:"gdp_incr"`
}

func GetYearBooks(code int) ([]YearBooksResult, error) {
	o := orm.NewOrm()
	var yearBooksResults []YearBooksResult
	_, err := o.Raw("SELECT id,county,gdp,gdp_incr FROM yearbooks WHERE code=?", code).QueryRows(&yearBooksResults)
	if err != nil {
		return []YearBooksResult{}, err
	}
	return yearBooksResults, nil
}
