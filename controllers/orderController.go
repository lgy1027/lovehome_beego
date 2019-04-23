package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"lovehome_beego/models"
)

// 订单
type OrderController struct {
	beego.Controller
}

func (c *OrderController) RetData(resp map[string]interface{}) {
	c.Data["json"] = &resp
	c.ServeJSON()
}

// 订单
func (c *OrderController) GetOrderData() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)

	uid := c.GetSession("user_id")

	// 根据url获取当前操作的角色
	role := c.GetString("role")
	if role == "custom" {
		o := orm.NewOrm()
		orders := []models.OrderHouse{}
		qs := o.QueryTable("OrderHouse")
		user := models.User{Id: uid.(int)}
		qs.Filter("user_id", uid.(int)).All(&orders)
		for _, order := range orders {
			order.User = &user
			o.LoadRelated(order, "User")
		}

		respData := make(map[string]interface{})
		respData["orders"] = orders
		resp["data"] = respData
		resp["errno"] = models.RECODE_OK
		resp["errmsg"] = models.RecodeText(models.RECODE_OK)
		return
	}

	if role == "landlord" {

	}

	if role == "" {
		resp["errno"] = models.RECODE_ROLEERR
		resp["errmsg"] = models.RecodeText(models.RECODE_ROLEERR)
		return
	}
}
