package models

import (
	"log"
	db "lovehome/gin_demo/database"
)

type Area struct {
	Id     int      `json:"id" form:"id"`                    //区域编号
	Name   string   `orm:"size(32)" json:"name" form:"name"` //区域名字
	Houses []*House `orm:"reverse(many)" json:"houses"`      //区域所有的房屋
}

func (this *Area) GetArea() (areas []Area, err error) {
	areas = make([]Area, 0)
	rows, err := db.SqlDB.Query("SELECT id, name FROM area")
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	for rows.Next() {
		var area Area
		rows.Scan(&area.Id, &area.Name)
		areas = append(areas, area)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}
