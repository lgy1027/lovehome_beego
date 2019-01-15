package models

/* 设施信息 table_name = "facility"*/
type Facility struct {
	Id     int      `json:"fid"`     //设施编号
	Name   string   `orm:"size(32)"` //设施名字
	Houses []*House `orm:"rel(m2m)"` //都有哪些房屋有此设施
}
