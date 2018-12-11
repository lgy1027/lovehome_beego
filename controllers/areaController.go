package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
)

// 区域管理
type AreaController struct {
	beego.Controller
}

func (c *AreaController) Get(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

}
