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

// 头像上传
//TODO 后期使用ipfs保存文件，等待完善
func (c *UserController) PostAvatar() {
	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_REQERR
	resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
	defer c.RetData(resp)

	//avatar,_,err := c.GetFile("avatar")

	//if err != nil {
	//	resp["errno"] = models.RECODE_REQERR
	//	resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
	//
	//	return
	//}
	//hash,err := utils.UploadIPFS()
	// 获取文件后缀
	//suffix := path.Ext(hd.Filename)

	// 获取用户id
	//uid := c.GetSession("user_id")
	//o := orm.NewOrm()
	//user := models.User{}
	//qs := o.QueryTable("user")
	//qs.Filter("id",uid).One(&user)
	//user.Avatar_url = "待插入"
	//_,err = o.Update(&user)
	//if err != nil {
	//	resp["errno"] = models.RECODE_REQERR
	//	resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
	//
	//	return
	//}

}

// 获取个人信息
func (c *UserController) UserInffo() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)

	// 获取用户id
	uid := c.GetSession("user_id")
	o := orm.NewOrm()
	user := models.User{Id: uid.(int)}
	err := o.Read(&user)
	if err != nil {
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}
	resp["data"] = &user
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
}

// 更新名字
func (c *UserController) UpdateName() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)

	username := make(map[string]string)
	json.Unmarshal(c.Ctx.Input.RequestBody, &username)
	// 获取用户id
	uid := c.GetSession("user_id")
	o := orm.NewOrm()
	user := models.User{Id: uid.(int)}
	err := o.Read(&user)

	user.Name = username["name"]
	_, err = o.Update(&user)
	if err != nil {
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}

	c.SetSession("name", username["name"])
	resp["data"] = username
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
}
