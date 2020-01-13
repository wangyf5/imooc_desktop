package routers

import (
	"imooc/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login",&controllers.MainController{},"Post:RegPost")
}
