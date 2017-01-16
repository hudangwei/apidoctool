package main

import (
	"github.com/astaxie/beego"
	_ "github.com/hudangwei/apidoctool/initial"
    _ "github.com/hudangwei/apidoctool/routers"
)

func main() {
	beego.Run()
}
