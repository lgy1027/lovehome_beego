package controllers

import (
	"github.com/astaxie/beego"
	"lovehome/models"
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

	resp["errno"] = models.RECODE_DATAERR
	resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
	c.RetData(resp)
}
