package controllers

import (
	"github.com/astaxie/beego"
)

// Controller基类继承封装
type BaseController struct {
	beego.Controller
}

// run before get
func (this *BaseController) Prepare() {

}

// run after finished
func (this *BaseController) Finish() {

}
