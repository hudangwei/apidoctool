package routers

import (
	"github.com/astaxie/beego"
	"github.com/hudangwei/apidoctool/controllers"
	"github.com/hudangwei/apidoctool/controllers/index"
)

func init() {
	beego.Router("/", &index.MainController{})
	beego.ErrorController(&controllers.ErrorController{})
}
