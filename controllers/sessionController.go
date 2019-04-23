package controllers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"lovehome_beego/models"
)

type SessionController struct {
	beego.Controller
}

func (this *SessionController) RetData(resp map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

// 判断用户是否已登陆
func (c *SessionController) GetSessionData() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)

	resp["errno"] = models.RECODE_DBERR
	resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
	user := models.User{}
	name := c.GetSession("name")
	if name != nil {
		user.Name = name.(string)
		resp["errno"] = models.RECODE_OK
		resp["errmsg"] = models.RecodeText(models.RECODE_OK)
		resp["data"] = user
	}
}

// 退出登陆删除session
func (c *SessionController) DeleteSessionData() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)
	c.DelSession("name")

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
}

// 登陆
func (c *SessionController) Login() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)

	// 获取前端参数
	json.Unmarshal(c.Ctx.Input.RequestBody, &resp)

	//2.判断是否合法
	if resp["mobile"] == nil || resp["password"] == nil {
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
		return
	}

	//3.与数据库匹配判断账号密码正确

	o := orm.NewOrm()
	var user models.User
	mobile := resp["mobile"].(string)
	qs := o.QueryTable("user")
	err := qs.Filter("mobile", mobile).One(&user)

	if err != nil {
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
		return
	}

	if user.Password_hash != fmt.Sprintf("%x", md5.Sum([]byte(resp["password"].(string)))) {
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
		return
	}

	//4.添加session
	c.SetSession("name", resp["mobile"])
	c.SetSession("mobile", resp["mobile"])
	c.SetSession("user_id", user.Id)

	//5.返回json数据给前端
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
}
