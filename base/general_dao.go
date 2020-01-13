package base

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"reflect"
	"ssnbee/utils"
)

type GenericDao struct {
}

func (mgr *GenericDao) GetPagedList(user interface{}, query BasePageQuery, tableName string) (pager utils.Pager) {
	pageNum := query.PageNum
	pageSize := query.PageSize
	o := orm.NewOrm()
	//user := new([]entity.SysUser)
	/*	o.QueryTable("sys_user").Limit(pageSize, (pageNum-1)*pageSize).All(user)
		TotalCount, _ := o.QueryTable("sys_user").Count()*/
	table := o.QueryTable(tableName)

	types := reflect.TypeOf(query)
	values := reflect.ValueOf(query)
	for i := 0; i < types.NumField(); i++ {
		// 获取每个成员的结构体字段类型
		fieldType := types.Field(i)
		if fieldType.Type.String() == "string" {
			v := values.Field(i).String()
			if v != "" {
				table = table.Filter(fieldType.Name, v)
			}
		}
		// 获取interface{}类型的值, 通过类型断言转换
		//fmt.Printf("name: %v  tag: '%v'  %v  %v\n", fieldType.Name, fieldType.Tag,fieldType.Type,)
	}
	table.Limit(pageSize, (pageNum-1)*pageSize).All(user)
	TotalCount, _ := table.Count()

	pager.Total = int(TotalCount)
	pager.PageSize = pageSize
	pager.List = user
	fmt.Println(user)
	return pager

}

type GeneralDao interface {
	GetPagedList(pageNum int, pageSize int) (pager utils.Pager)
	GetList() interface{}
	DeleteById(id int) bool
	Update(sysUser interface{}) bool
	Insert(sysUser interface{}) (model interface{})
}
