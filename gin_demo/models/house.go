package models

import "time"

/* 房屋信息 table_name = house */
type House struct {
	Id              int           `json:"house_id"`                                          //房屋编号
	User            *User         `orm:"rel(fk)" json:"user_id"`                             //房屋主人的用户编号
	Area            *Area         `orm:"rel(fk)" json:"area_id"`                             //归属地的区域编号
	Title           string        `orm:"size(64)" json:"title"`                              //房屋标题
	Price           int           `orm:"default(0)" json:"price"`                            //单价,单位:分
	Address         string        `orm:"size(512)" orm:"default("")" json:"address"`         //地址
	Room_count      int           `orm:"default(1)" json:"room_count"`                       //房间数目
	Acreage         int           `orm:"default(0)" json:"acreage"`                          //房屋总面积
	Unit            string        `orm:"size(32)" orm:"default("")" json:"unit"`             //房屋单元,如 几室几厅
	Capacity        int           `orm:"default(1)" json:"capacity"`                         //房屋容纳的总人数
	Beds            string        `orm:"size(64)" orm:"default("")" json:"beds"`             //房屋床铺的配置
	Deposit         int           `orm:"default(0)" json:"deposit"`                          //押金
	Min_days        int           `orm:"default(1)" json:"min_days"`                         //最好入住的天数
	Max_days        int           `orm:"default(0)" json:"max_days"`                         //最多入住的天数 0表示不限制
	Order_count     int           `orm:"default(0)" json:"order_count"`                      //预定完成的该房屋的订单数
	Index_image_url string        `orm:"size(256)" orm:"default("")" json:"index_image_url"` //房屋主图片路径
	Facilities      []*Facility   `orm:"reverse(many)" json:"facilities"`                    //房屋设施
	Images          []*HouseImage `orm:"reverse(many)" json:"img_urls"`                      //房屋的图片
	Orders          []*OrderHouse `orm:"reverse(many)" json:"orders"`                        //房屋的订单
	Ctime           time.Time     `orm:"auto_now_add;type(datetime)" json:"ctime"`
}
