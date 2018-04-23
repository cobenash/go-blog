package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type UserController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "index.tpl"
}

// 登录
func (this *UserController) Get() {
	this.Data["Website"] = "TESTPAGE"
	this.Data["Email"] = "victor.yang@hellosanta.com.tw"
	this.TplName = "main/login.tpl"
}
