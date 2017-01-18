package routers

import (
	"github.com/astaxie/beego"
	"github.com/hudangwei/apidoctool/controllers/api"
)

func init() {
	beego.Router("/search", &api.SearchController{})
}
