package utils

import "github.com/astaxie/beego"

func InitTmpl(ctrl beego.Controller)  {

	ctrl.Layout = "shared/layout.html"
	//其他的部分
	ctrl.LayoutSections = make(map[string]string)
	ctrl.LayoutSections["footerSection"] ="shared/footer.html"
	ctrl.LayoutSections["navSection"] ="shared/nav.html"
}