package routers

import (
	"github.com/astaxie/beego"
	"github.com/hiprice/beewechat/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//beego.Router("/home", &controllers.WeChatController{})
	beego.Router("/receiver", &controllers.WeChatController{})

	// beego.Router("/test", &controllers.TestController{})
	// beego.Router("/b", &controllers.BController{})
}
