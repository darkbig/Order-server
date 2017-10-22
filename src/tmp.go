package main

import (
	"Infrastructure/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController)Get(){
	this.Ctx.WriteString("我的第一个网页")
}
func main(){
	//注册路由
	beego.Router("/",&HomeController{})
	beego.Run()
}

