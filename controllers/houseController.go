package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"lovehome/models"
	"strconv"
)

type HouseController struct {
	beego.Controller
}

func (c *HouseController) RetData(resp map[string]interface{}) {
	c.Data["json"] = &resp
	c.ServeJSON()
}

// 获取房源
func (c *HouseController) GetHouses() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)

	houses := []models.House{}
	uid := c.GetSession("user_id")
	o := orm.NewOrm()
	qs := o.QueryTable("house")
	num, err := qs.Filter("user_id", uid.(int)).All(&houses)
	if err != nil {
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
		return
	}

	if num == 0 {
		resp["errno"] = models.RECODE_NODATA
		resp["errmsg"] = models.RecodeText(models.RECODE_NODATA)
		return
	}

	respData := make(map[string]interface{})
	respData["houses"] = &houses

	resp["data"] = &respData
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
}

// 发布房源
func (c *HouseController) PostHouse() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)

	// 获取前端数据
	req := make(map[string]interface{})
	// 解析body中数据
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	// 插入数据到数据库中
	//房间信息
	house := models.House{}
	house.Title = req["title"].(string)
	price, _ := strconv.Atoi(req["price"].(string))
	house.Price = price
	house.Address = req["address"].(string)
	room_count, _ := strconv.Atoi(req["room_count"].(string))
	house.Room_count = room_count
	house.Unit = req["unit"].(string)
	house.Beds = req["beds"].(string)
	min_days, _ := strconv.Atoi(req["min_days"].(string))
	house.Min_days = min_days
	max_days, _ := strconv.Atoi(req["max_days"].(string))
	house.Max_days = max_days

	// 设备信
	facility := []*models.Facility{}
	for _, fid := range req["facility"].([]interface{}) {
		beego.Info("fid:", fid)
		f_id, _ := strconv.Atoi(fid.(string))
		fac := &models.Facility{Id: f_id}
		facility = append(facility, fac)
	}

	area_id, _ := strconv.Atoi(req["area_id"].(string))
	area := models.Area{Id: area_id}
	house.Area = &area

	// t填充用户信息
	uid := c.GetSession("user_id")
	user := models.User{Id: uid.(int)}
	house.User = &user

	o := orm.NewOrm()
	house_id, err := o.Insert(&house)
	if err != nil {
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
		return
	}
	house.Id = int(house_id)
	m2m := o.QueryM2M(&house, "Facilities")
	num, err := m2m.Add(facility)
	if err != nil || num == 0 {
		resp["errno"] = models.RECODE_NODATA
		resp["errmsg"] = models.RecodeText(models.RECODE_NODATA)
		return
	}
	respData := make(map[string]interface{})
	respData["house_id"] = &house_id
	resp["data"] = &respData
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
}

func (c *HouseController) HouseById() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)

	uid := c.GetSession("user_id")
	user := models.User{Id: uid.(int)}
	house_id := c.Ctx.Input.Param(":id")
	hid, _ := strconv.Atoi(house_id)

	o := orm.NewOrm()
	house := models.House{Id: hid}
	house.User = &user

	if err := o.Read(&house); err != nil {
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
		return
	}
	// 载入关联字段
	o.LoadRelated(&house, "Area")
	o.LoadRelated(&house, "User")
	o.LoadRelated(&house, "Images")
	o.LoadRelated(&house, "Facilities")

	facs := []string{}
	for _, fac := range house.Facilities {
		fid := strconv.Itoa(fac.Id)
		facs = append(facs, fid)
	}
	respData := make(map[string]interface{})
	respData["acreage"] = house.Acreage
	respData["address"] = house.Address
	respData["beds"] = house.Beds
	respData["capacity"] = house.Capacity
	respData["deposit"] = house.Deposit
	respData["images"] = house.Images
	respData["min_days"] = house.Min_days
	respData["max_days"] = house.Max_days
	respData["price"] = house.Price
	respData["facilities"] = facs

	resp["house"] = &house
	resp["data"] = &respData

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
}
