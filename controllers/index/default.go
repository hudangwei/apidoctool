package index

import (
	"github.com/hudangwei/apidoctool/controllers"
)

// 首页
type MainController struct {
	controllers.BaseController
}

func (this *MainController) Get() {
	this.TplName = "index.tpl"
}
