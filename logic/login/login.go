package login

import (
	"github.com/astaxie/beego/orm"
)

func IsUser(username string) (bool, error) {
	o := orm.NewOrm()
	var userCount int
	err := o.Raw("SELECT count(*) FROM user WHERE username=?", username).QueryRow(&userCount)
	if err != nil {
		return false, err
	}
	if userCount > 0 {
		return true, nil
	}
	return false, nil
}

func Login(username string, password string) (bool, error) {
	o := orm.NewOrm()
	var userCount int
	err := o.Raw("SELECT count(*) FROM user WHERE username=? AND password=?", username, password).QueryRow(&userCount)
	if err != nil {
		return false, err
	}
	if userCount > 0 {
		return true, nil
	}
	return false, nil
}
