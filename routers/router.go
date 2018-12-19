package routers

import (
	"github.com/astaxie/beego"
	"lovehome/controllers"
)

func init() {
	//首页
	beego.Router("/", &controllers.MainController{})
	// 区域获取
	beego.Router("/api/v1.0/areas", &controllers.AreaController{}, "get:GetArea")
	// 显示登陆注册按钮
	beego.Router("/api/v1.0/houses/index", &controllers.HouseIndexController{}, "get:GetHouseIndex")

	//api/v1.0/session
	beego.Router("/api/v1.0/session", &controllers.SessionController{}, "get:GetSessionData;delete:DeleteSessionData")

	//api/v1.0/users   注册功能
	beego.Router("/api/v1.0/users", &controllers.UserController{}, "post:Reg")

	//api/v1.0/sessions 登陆
	beego.Router("/api/v1.0/sessions", &controllers.SessionController{}, "post:Login")
	//api/v1.0/user/avatar
	beego.Router("/api/v1.0/sessions", &controllers.UserController{}, "post:PostAvatar")

}
