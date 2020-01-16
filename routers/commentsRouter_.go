package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["ssnbee/controllers:UserController"] = append(beego.GlobalControllerRouter["ssnbee/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Chat",
			Router:           `/user/chat`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ssnbee/controllers:UserController"] = append(beego.GlobalControllerRouter["ssnbee/controllers:UserController"],
		beego.ControllerComments{
			Method:           "ChatRoom",
			Router:           `/user/chatroom`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ssnbee/controllers:UserController"] = append(beego.GlobalControllerRouter["ssnbee/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Detail",
			Router:           `/user/edit`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ssnbee/controllers:UserController"] = append(beego.GlobalControllerRouter["ssnbee/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/user/list/?:pageNum`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ssnbee/controllers:UserController"] = append(beego.GlobalControllerRouter["ssnbee/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/user/save`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
