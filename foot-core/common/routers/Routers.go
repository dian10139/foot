package routers

import (
	"github.com/astaxie/beego"
	"tesou.io/platform/foot-parent/foot-core/module/index/controller"
	controller2 "tesou.io/platform/foot-parent/foot-core/module/match/controller"
	controller3 "tesou.io/platform/foot-parent/foot-core/module/wechat/controller"
)

type Routers struct {
}

func init() {
	beego.Router("/", &controller.IndexController{})

	//match
	beego.AutoRouter(&controller2.MatchController{})
	beego.AutoRouter(&controller2.MatchLastController{})

	//wechat
	beego.AutoRouter(&controller3.WeChatController{})

}
