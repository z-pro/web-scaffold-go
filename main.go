package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"runtime"
	"ssnbee/models/entity"
	_ "ssnbee/routers"
	"time"

	// 使用 go get github.com/astaxie/beego/orm 获取包
	"github.com/astaxie/beego/orm"
	//mysql数据库驱动
	// 使用 go get github.com/go-sql-driver/mysql 获取
	//下划线
	_ "github.com/go-sql-driver/mysql"

	_ "ssnbee/managers/mq/consume"
)

//初始化操作
//1.完成模型、数据库、驱动的注册
//2.创建表
func init() {
	fmt.Println("init")
	// 注册数据库
	//第三个参数：对应的链接字符串
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp("+beego.AppConfig.String("dbhost")+":3306)/blog?charset=utf8")
	// 注册Model
	orm.RegisterModel(new(entity.SysUser))
	orm.RegisterModel(new(entity.User))
	// 注册驱动
	// 参数1   driverName
	// 参数2   数据库类型，这个用来设置 driverName 对应的数据库类型
	// mysql / sqlite3 / postgres 这三种是默认已经注册过的，所以可以无需设置
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 创建表格，只会创建一次，重复会跳过
	//第一个参数，表的别名，默认为“default”
	//第二个参数，如果当前建表出错，跳过执行下一个
	//第三个参数，如果表存在给提醒，改为false不会提醒
	//如果没有这个函数，会报错，提示主键不存在
	orm.RunSyncdb("default", false, true)

	/*	//切换数据库
		orm.RegisterDataBase("db1", "mysql", "root:root@/orm_db2?charset=utf8")
		orm.RegisterDataBase("db2", "sqlite3", "data.db")
		o1 := orm.NewOrm()
		o1.Using("db1")
		o2 := orm.NewOrm()
		o2.Using("db2")*/
}

func main() {
	//配置全局
	//beego.BConfig.WebConfig.Session.SessionOn=true
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 86400 //设置Session有效期,单位秒
	beego.Run()

}
func testMaxCPU() {
	cpuNum := runtime.NumCPU() //获得当前设备的cpu核心数
	fmt.Println("cpu核心数:", cpuNum)
	runtime.GOMAXPROCS(cpuNum)
}

func testStoper() {

	// 创建一个计时器, 2秒后触发
	stopper := time.NewTimer(time.Second * 2)
	ticker := time.NewTicker(time.Millisecond * 500)
	for {
		select {
		case <-stopper.C:
			fmt.Print("ssss")
			goto stopTag
		case <-ticker.C:
			fmt.Println(12)
		}

	}

stopTag:
	fmt.Print("over")

}

func testAfterFun() {
	exit := make(chan int)
	fmt.Println("开始")
	time.AfterFunc(time.Second*3, func() {
		exit <- 1
	})
	<-exit
	fmt.Println("end")
}

func testChan() {
	var ch = make(chan int, 3)
	ch <- 1
	ch <- 3
	ch <- 3
	<-ch
	<-ch
	<-ch
	select {
	case <-ch:
		fmt.Println("1")
	case <-time.After(time.Second * 3):
		fmt.Println("3333")
	}
	fmt.Println(len(ch))
}
