package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
	name string
	age int
	data map[string] string
}

func (ct *MainController) Show(){
	ct.age=10
	ct.name="zs"
	ct.data=map[string]string{"a1":"df","b1":"ssdf"}
	fmt.Print(ct.data)
}

func (ct *MainController) show(arr []int,mps map[string]string ){

	for i:=range arr{
		print(i)
	}
	for k,v:=range mps{
		fmt.Print(k)
		fmt.Print(v)
	}

}

func(ct *MainController) mutil(arg...int)  {

	fmt.Print(arg)
}


func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"

	var ctrl MainController


	var data = map[string]string{"a":"df","b":"ssdf"}

	ctrl1:=MainController{age:111,name:"lise",data:data}
	fmt.Print(ctrl1.data)
	ctrl1.Show()
	fmt.Print(ctrl1.data)

	ctrl1.show([]int{1,2,54,656},map[string]string{"aa":"bb"})


	ctrl2:=new(MainController)
	var ctrl3 MainController=MainController{}
	ctrl3.name="sdf"
	fmt.Print(ctrl)
	fmt.Print("\n")
	fmt.Print(ctrl1)
	fmt.Print("\n")
	fmt.Print(*ctrl2)
	fmt.Print("\n")
	fmt.Print(ctrl3)
	fmt.Print("\n")

	fmt.Printf("%T",ctrl3)


}
