package controllers

import (
	"github.com/astaxie/beego"
	"lovehome_beego/models"
)

type HouseIndexController struct {
	beego.Controller
}

// 转换成json数据返回
func (c *HouseIndexController) RetData(resp map[string]interface{}) {
	c.Data["json"] = &resp
	c.ServeJSON()
}

func (c *HouseIndexController) GetHouseIndex() {
	resp := make(map[string]interface{})

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	c.RetData(resp)
}
