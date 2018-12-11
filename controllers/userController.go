package controllers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"lovehome/models"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) RetData(resp map[string]interface{}) {
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *UserController) Reg() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)

	//获取前端传过来的json数据
	json.Unmarshal(c.Ctx.Input.RequestBody, &resp)

	//插入数据库
	o := orm.NewOrm()
	user := models.User{}
	user.Password_hash = fmt.Sprintf("%x", md5.Sum([]byte(resp["password"].(string))))
	user.Name = resp["mobile"].(string)
	user.Mobile = resp["mobile"].(string)

	id, err := o.Insert(&user)
	if err != nil {
		resp["errno"] = 4002
		resp["errmsg"] = "注册失败"
		return
	}

	beego.Info("reg success ,id = ", id)
	resp["errno"] = 0
	resp["errmsg"] = "注册成功"

	c.SetSession("name", user.Name)
}
