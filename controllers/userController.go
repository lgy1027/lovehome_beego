package controllers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"lovehome_beego/models"
	"lovehome_beego/utils"
	"os"
)

var url_prefix = "http://ipfs.io/ipfs/"

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
func (c *UserController) PostAvatar() {
	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_REQERR
	resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
	defer c.RetData(resp)

	f, h, err := c.GetFile("avatar")
	//获取文件后缀
	//suffix := path.Ext(h.Filename)
	//filename := h.Filename
	c.SaveToFile("avatar", "static/upload/"+h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建

	//删除文件
	defer os.Remove("static/upload/" + h.Filename)

	f, err = os.Open("static/upload/" + h.Filename)
	defer f.Close()

	fs, err := ioutil.ReadAll(f)
	hash, err := utils.UploadIPFS(string(fs))
	beego.Info("hash:", hash)
	if err != nil {
		resp["errno"] = models.RECODE_UNKNOWERR
		resp["errmsg"] = models.RecodeText(models.RECODE_UNKNOWERR)
		return
	}

	if err != nil {
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)

		return
	}
	//hash,err := utils.UploadIPFS()

	//获取用户id
	uid := c.GetSession("user_id")
	o := orm.NewOrm()
	user := models.User{}
	qs := o.QueryTable("user")
	qs.Filter("id", uid).One(&user)
	user.Avatar_url = url_prefix + hash
	_, err = o.Update(&user)
	if err != nil {
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)

		return
	}
	respData := make(map[string]interface{})
	respData["avatar_url"] = url_prefix + hash
	resp["data"] = respData
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
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

	UserName := make(map[string]string)
	json.Unmarshal(c.Ctx.Input.RequestBody, &UserName)
	// 获取用户id
	uid := c.GetSession("user_id")
	o := orm.NewOrm()
	user := models.User{Id: uid.(int)}
	if o.Read(&user) == nil {
		user.Name = UserName["name"]
		if _, err := o.Update(&user, "name"); err == nil {
			c.SetSession("name", UserName["name"])
			resp["data"] = UserName
			resp["errno"] = models.RECODE_OK
			resp["errmsg"] = models.RecodeText(models.RECODE_OK)
			return
		}
	}
	resp["errno"] = models.RECODE_DATAERR
	resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
}

func (c *UserController) UpdateCert() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)

	uid := c.GetSession("user_id")
	realName := make(map[string]string)
	json.Unmarshal(c.Ctx.Input.RequestBody, &realName)
	o := orm.NewOrm()
	user := models.User{Id: uid.(int)}
	if o.Read(&user) == nil {
		user.Real_name = realName["real_name"]
		user.Id_card = realName["id_card"]
		if _, err := o.Update(&user); err == nil {
			c.SetSession("user_id", uid)

			resp["errno"] = models.RECODE_OK
			resp["errmsg"] = models.RecodeText(models.RECODE_OK)
			return
		}
	}
	resp["errno"] = models.RECODE_DATAERR
	resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
}
