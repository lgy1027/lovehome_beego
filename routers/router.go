package routers

import (
	"github.com/astaxie/beego"
	"lovehome_beego/controllers"
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
	//api/v1.0/user/avatar 起别名
	beego.Router("/api/v1.0/user/avatar", &controllers.UserController{}, "post:PostAvatar")
	//api/v1.0/user 展示用户信息
	beego.Router("/api/v1.0/user", &controllers.UserController{}, "get:UserInffo")
	//api/v1.0/user/name  更新用户姓名
	beego.Router("/api/v1.0/user/name", &controllers.UserController{}, "put:UpdateName")
	//api/v1.0/user/auth 实名认证
	beego.Router("/api/v1.0/user/auth", &controllers.UserController{}, "get:UserInffo;post:UpdateCert")
	//api/v1.0/user/houses 查看房源
	beego.Router("/api/v1.0/user/houses", &controllers.HouseController{}, "get:GetHouses;POST:PostHouse")
	//api/v1.0/houses 发布房源
	beego.Router("/api/v1.0/houses", &controllers.HouseController{}, "post:PostHouse")
	//api/v1.0/houses/1
	beego.Router("/api/v1.0/houses/?:id", &controllers.HouseController{}, "get:HouseById")
	//api/v1.0/user/orders 订单
	beego.Router("/api/v1.0/user/orders", &controllers.OrderController{}, "get:GetOrderData")

}
