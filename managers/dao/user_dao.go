package dao

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"reflect"
	"ssnbee/base"
	"ssnbee/models/entity"
	"ssnbee/models/query"
	"ssnbee/utils"
)

type userDao struct {
	base.GenericDao
}

var UserDao = new(userDao)

func (mgr *userDao) GetPagedList(query query.UserQuery) (pager utils.Pager) {
	pageNum := query.PageNum
	pageSize := query.PageSize
	o := orm.NewOrm()
	user := new([]entity.User)
	table := o.QueryTable("user")
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

func (mgr *userDao) GetList() interface{} {
	o := orm.NewOrm()
	user := new([]entity.User)
	o.QueryTable("user").All(user)
	return user
}

func (mgr *userDao) DeleteById(id int) bool {
	o := orm.NewOrm()
	i, err := o.Delete(&entity.User{Id: id})
	if err == nil {
		logs.Debug(i)
	}
	return i > 0
}

func (mgr *userDao) SelectById(id int) (model entity.User, err error) {
	o := orm.NewOrm()
	ob := entity.User{Id: id}
	err = o.Read(&ob)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
		return
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
		return
	}
	return ob, err
}

func (mgr *userDao) Update(sysUser entity.User) bool {
	orm.Debug = true
	var o orm.Ormer
	o = orm.NewOrm()
	if i, err := o.Update(&sysUser); err == nil {
		logs.Debug(i)
		return true
	}
	return false
}

func (mgr *userDao) Insert(sysUser entity.User) (model entity.User) {
	var o orm.Ormer
	o = orm.NewOrm()
	_, err := o.Insert(&sysUser)
	if err == nil {
		model = sysUser
	}
	return model
}
