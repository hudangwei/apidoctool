package initial

import (
	"github.com/astaxie/beego"
	"fmt"
)

func InitEnv() {
	runmode := beego.AppConfig.String("runmode")
	if runmode == "dev" {
		//beego.SetStaticPath("/static", "static")
		fmt.Println(runmode)
	}
	fmt.Println(runmode)
}
