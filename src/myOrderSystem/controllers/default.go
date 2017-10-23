package controllers

import (
	"github.com/astaxie/beego"
	"myOrderSystem/models"
	"github.com/astaxie/beego/logs"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	tmp :=&models.UserInfo{
		"yangwei123",
		"18112188839",
		"8821399",
	}
	res,err:=models.UserMesUpData("order_system","user_info",tmp)
	if err!=""{
		logs.Info(err)
	}
	logs.Info("插入状态",res)
}
