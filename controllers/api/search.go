package api

import "github.com/hudangwei/apidoctool/controllers"

type SearchController struct {
    controllers.BaseController
}

func (this *SearchController) Get()  {
    this.Data["json"] = map[string]interface{}{"data": "just is a test case"}
    this.ServeJSON()
}