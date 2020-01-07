package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"ssnbee/managers"
	"ssnbee/models/entity"
	"ssnbee/models/query"
	"ssnbee/utils"
	"strconv"
	"time"
)

type SysuserController struct {
	beego.Controller
}
func (ctrl *SysuserController) Get() {
	ctrl.Layout = "shared/layout.html"
	ctrl.TplName = "sysuser/list.html"

	//其他的部分
	ctrl.LayoutSections = make(map[string]string)
	ctrl.LayoutSections["navSection"] = "shared/nav.html"

	pageNum, err := strconv.Atoi(ctrl.Ctx.Input.Param(":pageNum"))
	if err!=nil{
		pageNum=1
	}

	sysUserManager := new(managers.SysUserManager)
	pageSize, _ := beego.GetConfig("Int", "pagesize", 10)
	var query query.SysUserQuery
	ctrl.ParseForm(&query)
	query.PageNum=pageNum
	query.PageSize=pageSize.(int)
	pager := sysUserManager.GetPagedList(query)

	ctrl.Data["sysuserList"] = pager.List
	ctrl.Data["query"] =query
	ctrl.Data["pageHtml"], _, _ =utils.LimitPage(pageNum,pager.Total,ctrl.Ctx.Request.URL.RawQuery,"/sysuser/list")

}

// 详情
func (ctrl *SysuserController) Detail() {
	ctrl.Layout = "shared/layout.html"
	ctrl.TplName = "sysuser/edit.html"

	//其他的部分
	ctrl.LayoutSections = make(map[string]string)
	ctrl.LayoutSections["navSection"] = "shared/nav.html"

	sysUserManager := new(managers.SysUserManager)
	id, _ := ctrl.GetInt("id", 0)
	entity,_ := sysUserManager.SelectById(id)
	ctrl.Data["entity"] = entity

}

// 保存
func (ctrl *SysuserController) Post() {
	var ob entity.SysUser
	//json.Unmarshal(ctrl.Ctx.Input.RequestBody, &ob)
	ctrl.ParseForm(&ob)
	logs.Debug(ob)
	sysUserManager := new(managers.SysUserManager)
	model,_ := sysUserManager.SelectById(ob.Id)
	if &model == nil || model.Id <= 0 { //新增
		ob.CreateDate = time.Now()
		sysUserManager.Insert(ob)
	} else {
		model.RealName = ob.RealName
		model.Phone=ob.Phone
		model.Remark=ob.Remark
		sysUserManager.Update(model)
	}
	ctrl.Redirect("/sysuser/list", 302)
}

// 删除
func (ctrl *SysuserController) Delete() {
	id, _ := ctrl.GetInt("id", 0)
	sysUserManager := new(managers.SysUserManager)
	succ := sysUserManager.DeleteById(id)
	res := make(map[string]interface{})
	if succ{
		res["errcode"] = 0
	}else
	{
		res["errcode"] = 1
		res["errmsg"] = "删除失败！"
	}
	bytes, _ := json.Marshal(res)
	ctrl.Ctx.WriteString(string(bytes))


}
