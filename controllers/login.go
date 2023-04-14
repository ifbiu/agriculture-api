package controllers

import (
	"agriculture-api/logic/login"
	"agriculture-api/utils"
	"encoding/json"
	"fmt"
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
	data := this.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &user)
	valid.Required(user.UserName, "UserName").Message("用户名不能为空！")
	valid.MaxSize(user.UserName, 16, "UserName").Message("用户名不能超过16位！")
	valid.Required(user.PassWord, "PassWord").Message("密码不能为空！")
	valid.MinSize(user.PassWord, 6, "PassWord").Message("密码最小6位！")
	if valid.HasErrors() {
		reJson["code"] = "500"
		errMsg := ""
		for _, error := range valid.Errors {
			errMsg = error.Message + " "
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
	user1 := utils.User{"1", user.UserName}
	token, err := utils.GenerateToken(&user1, 0)
	if err != nil {
		fmt.Println(err)
		reJson["code"] = "500"
		reJson["msg"] = "token获取失败！"
		return
	}
	reJson["code"] = "200"
	reJson["msg"] = "登录成功"
	reJson["token"] = token
}
