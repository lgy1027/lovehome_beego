package models

/* 房屋图片 table_name = "house_image"*/
type HouseImage struct {
	Id    int    `json:"house_image_id"`         //图片id
	Url   string `orm:"size(256)" json:"url"`    //图片url
	House *House `orm:"rel(fk)" json:"house_id"` //图片所属房屋编号
}
