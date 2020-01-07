package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"ssnbee/managers"
	"ssnbee/models/query"
	"ssnbee/utils"
	"strconv"
	"strings"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) URLMapping() {
	//c.Mapping("Detail", c.Detail)
	//c.Mapping("AllBlock", c.AllBlock)
}

// @router /user/list/?:pageNum
func (ctrl *UserController) Get()  {
	ctrl.TplName ="user/list.html"
	ctrl.Layout = "shared/layout.html"

	//其他的部分
	ctrl.LayoutSections = make(map[string]string)
	ctrl.LayoutSections["navSection"] = "shared/nav.html"

	pageNum, err := strconv.Atoi(ctrl.Ctx.Input.Param(":pageNum"))
	if err!=nil{
		pageNum=1
	}

	manager := managers.NewUserManager() //new(managers.SysUserManager)
	pageSize, _ := beego.GetConfig("Int", "pagesize", 10)
	var query query.UserQuery
	ctrl.ParseForm(&query)
	query.PageNum=pageNum
	query.PageSize=pageSize.(int)
	pager := manager.GetPagedList(query)

	ctrl.Data["userList"] = pager.List
	ctrl.Data["query"] =query
	ctrl.Data["pageHtml"], _, _ =utils.LimitPage(pageNum,pager.Total,ctrl.Ctx.Request.URL.RawQuery,"/user/list")

}

// @router /user/edit [get]
func (ctrl *UserController) Detail()  {
	ctrl.TplName ="user/edit.html"
	ctrl.Layout = "shared/layout.html"

	//其他的部分
	ctrl.LayoutSections = make(map[string]string)
	ctrl.LayoutSections["navSection"] = "shared/nav.html"

	sysUserManager :=managers.NewUserManager() //new(managers.SysUserManager)
	id, _ := ctrl.GetInt("id", 0)
	entity,_ := sysUserManager.SelectById(id)
	ctrl.Data["entity"] = entity
}

// @router /user/save [post]
func (ctrl *UserController) Post() {
	_, header, e := ctrl.GetFile("the_file")
	utils.AssertErr(e)
	split := strings.Split(header.Filename, ".")
	ext :=split[1]
	fmt.Printf(header.Filename)
	ctrl.SaveToFile("the_file",fmt.Sprintf("e:/sss.%v",ext))
	ctrl.Redirect("/user/list", 302)
}
