package controllers

import (
	"agriculture-api/logic/login"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type User struct {
	UserName string
	PassWord string
}

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Post() {
	reJson := make(map[string]interface{})
	this.Data["json"] = reJson
	defer this.ServeJSON()
	valid := validation.Validation{}
	user := User{}
	user.UserName = this.GetString("username")
	user.PassWord = this.GetString("password")
	valid.Required(user.UserName, "UserName").Message("用户名不能为空！")
	valid.MaxSize(user.UserName, 16, "UserName").Message("用户名不能超过16位！")
	valid.Required(user.PassWord, "PassWord").Message("密码不能为空！")
	valid.MinSize(user.PassWord, 6, "PassWord").Message("密码最小6位！")
	if valid.HasErrors() {
		reJson["code"] = "500"
		errMsg := []string{}
		for _, error := range valid.Errors {
			errMsg = append(errMsg, error.Message)
		}
		reJson["msg"] = errMsg
		return
	}
	isUser, err := login.IsUser(user.UserName)
	if err != nil {
		reJson["code"] = "500"
		reJson["msg"] = err
		return
	}
	if !isUser {
		reJson["code"] = "500"
		reJson["msg"] = "用户不存在！"
		return
	}
	isLogin, err := login.Login(user.UserName, user.PassWord)
	if err != nil {
		reJson["code"] = "500"
		reJson["msg"] = err
		return
	}
	if !isLogin {
		reJson["code"] = "500"
		reJson["msg"] = "密码错误！"
		return
	}
	reJson["code"] = "200"
	reJson["msg"] = "登录成功"
}
