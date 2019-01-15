package models

import "time"

/* 订单 table_name = order */
type OrderHouse struct {
	Id          int       `json:"order_id"`               //订单编号
	User        *User     `orm:"rel(fk)" json:"user_id"`  //下单的用户编号
	House       *House    `orm:"rel(fk)" json:"house_id"` //预定的房间编号
	Begin_date  time.Time `orm:"type(datetime)"`          //预定的起始时间
	End_date    time.Time `orm:"type(datetime)"`          //预定的结束时间
	Days        int       //预定总天数
	House_price int       //房屋的单价
	Amount      int       //订单总金额
	Status      string    `orm:"default(WAIT_ACCEPT)"` //订单状态
	Comment     string    `orm:"size(512)"`            //订单评论
	Ctime       time.Time `orm:"auto_now_add;type(datetime)" json:"ctime"`
}
