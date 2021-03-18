package routers

import (
	"github.com/astaxie/beego"
	"online-exam-server-rec/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
