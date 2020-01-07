package routers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/plugins/cors"
	"ssnbee/controllers"
	"ssnbee/models"
	"strings"
)


func init() {

	// 最后一个参数必须设置为false 不然无法打印数据
	beego.InsertFilter("/*", beego.FinishRouter, FilterLog, false)

	beego.InsertFilter("/*", beego.BeforeRouter, FilterAdminAuth, false)

	// 这段代码放在router.go文件的init()的开头
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		//AllowOrigins: 	  []string{"http://"+beego.AppConfig.String("front_end_domain")+":"+beego.AppConfig.String("front_end_port")},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

    beego.Router("/", &controllers.MainController{},)
	beego.Router("/home/index", &controllers.HomeController{},"get:Index")
	beego.Router("/home/logout", &controllers.HomeController{},"get:Logout")
	beego.Router("/home/login", &controllers.HomeController{},"get:Login;post:PostLogin")


	beego.Router("/sysuser/list/?:pageNum", &controllers.SysuserController{})
	beego.Router("/sysuser/save", &controllers.SysuserController{})
	beego.Router("/sysuser/delete", &controllers.SysuserController{},"get:Delete")
	beego.Router("/sysuser/edit", &controllers.SysuserController{},"get:Detail")


	beego.Include(&controllers.UserController{})
	//beego.Router("/user/list/?:pageNum", &controllers.UserController{})

	beego.Router("/about",&controllers.AboutController{},"get:Get")


}

func testNs()  {
	//初始化 namespace
	ns :=
		beego.NewNamespace("/v1",
			beego.NSCond(func(ctx *context.Context) bool {
				if ctx.Input.Domain() == "api.beego.me" {
					return true
				}
				return false
			}),
			//beego.NSBefore(auth),
			beego.NSGet("/notallowed", func(ctx *context.Context) {
				ctx.Output.Body([]byte("notAllowed"))
			}),
			beego.NSRouter("/version", &controllers.HomeController{}, "get:ShowAPIVersion"),
			beego.NSRouter("/changepassword", &controllers.SysuserController{}),
			beego.NSNamespace("/shop",
				//beego.NSBefore(sentry),
				beego.NSGet("/:id", func(ctx *context.Context) {
					ctx.Output.Body([]byte("notAllowed"))
				}),
			),
			beego.NSNamespace("/cms",
				beego.NSInclude(
					&controllers.MainController{},
					&controllers.UserController{},
					&controllers.HomeController{},
				),
			),
		)
	//注册 namespace
	beego.AddNamespace(ns)
}



// 添加日志拦截器
var FilterLog = func(ctx *context.Context) {
	url, _ := json.Marshal(ctx.Input.Data()["RouterPattern"])
	params, _ := json.Marshal(ctx.Request.Form)
	outputBytes, _ := json.Marshal(ctx.Input.Data()["json"])
	divider := " - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -"
	topDivider := "┌" + divider
	middleDivider := "├" + divider
	bottomDivider := "└" + divider
	outputStr := "\n" + topDivider + "\n│ 请求地址:" + string(url) + "\n" + middleDivider + "\n│ 请求参数:" + string(params) + "\n│ 返回数据:" + string(outputBytes) + "\n" + bottomDivider
	logs.Info(outputStr)
}

// 后台权限校验
var FilterAdminAuth = func(ctx *context.Context) {
	session := ctx.Input.Session(models.ADMIN_SESSION_KEY)
	if session ==nil && !strings.Contains(ctx.Request.RequestURI,"login"){
		//fmt.Sprintf(url, rq.Encode())
		ctx.Redirect(302, "/home/login?r_url="+ctx.Request.RequestURI)
	}
	logs.Info("dddd")
}
//实现了如何实现自己的路由规则:
/*var UrlManager = func(ctx *context.Context) {
	// 数据库读取全部的 url mapping 数据
	urlMapping := model.GetUrlMapping()
	for baseurl,rule:=range urlMapping {
		if baseurl == ctx.Request.RequestURI {
			ctx.Input.RunController = rule.controller
			ctx.Input.RunMethod = rule.method
			break
		}
	}
}

beego.InsertFilter("/*",beego.BeforeRouter,UrlManager)
*/