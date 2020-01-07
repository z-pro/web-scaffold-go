package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"ssnbee/managers"
	"ssnbee/models"
	"ssnbee/models/entity"
	"ssnbee/utils"
	"time"
)

type HomeController struct {
	beego.Controller
}

func (ctrl *HomeController) Login() {
	ctrl.TplName = "home/login.html"
}

func (ctrl *HomeController) PostLogin() {
	//var loginModel LoginModel
	//json.Unmarshal(ctrl.Ctx.Input.Params() ,&loginModel)
	userName := ctrl.GetString("username") //ctrl.Ctx.Request.Form["username"]
	password := ctrl.GetString("password") //ctrl.Ctx.Request.Form["password"]


	res := make(map[string]interface{})
	if checkLogin(userName,password){
		res["errcode"] = 0
		ctrl.SetSession(models.ADMIN_SESSION_KEY,userName)
	}else {
		res["errcode"] = 1
		res["errmsg"] = "用户名密码错误！"
	}

	bytes, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}
	//addSysUser(models.SysUser{UserName:"zs",Password:"323232323",CreateDate:time.Now()})
	user := entity.SysUser{UserName: "32234234", Id: 3,CreateDate:time.Now()}
	sysUserManager := new(managers.SysUserManager)
	sysUserManager.Update(user)
	ctrl.Ctx.WriteString(string(bytes))
}

func (ctrl *HomeController) Logout() {
	session := ctrl.GetSession(models.ADMIN_SESSION_KEY)
	if session!=nil{
		ctrl.DelSession(models.ADMIN_SESSION_KEY)
	}
	ctrl.Redirect("/home/login",302)
}


func (ctrl *HomeController) Index() {

	ctrl.Layout = "shared/layout.html"
	ctrl.TplName = "home/index.html"

	//其他的部分
	ctrl.LayoutSections = make(map[string]string)
	ctrl.LayoutSections["navSection"] ="shared/nav.html"
}


func checkLogin(username string,password string) bool{
	var o orm.Ormer
	o = orm.NewOrm()
	sysUser := entity.SysUser{UserName: username,Password:utils.EnWithMd5(password)}
	err := o.Read(&sysUser,"user_name","password")
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
		return false
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
		return false
	} else {
		fmt.Println(sysUser.Id, sysUser.RealName)
		return true
	}
}





