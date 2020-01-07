package managers

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"reflect"
	"ssnbee/models"
	"ssnbee/models/entity"
	"ssnbee/models/query"
	"ssnbee/models/vo"
	"ssnbee/utils"
)

type SysUserManager struct {
	DBConf *models.DBConfig
}

/*
 * usermanger构造器
 */
func NewSysUserManager(dbConf *models.DBConfig) *SysUserManager {
	mgr := &SysUserManager{
		DBConf: dbConf,
	}
	mgr.initDB() //初始化orm
	return mgr
}

/**
  初始化db，注册默认数据库，同时将实体模型也注册上去
*/
func (mgr *SysUserManager) initDB() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", mgr.DBConf.Username, mgr.DBConf.Password, mgr.DBConf.Host, mgr.DBConf.Port, mgr.DBConf.Database)
	err := orm.RegisterDataBase("default", "mysql", conn, mgr.DBConf.MaxIdleConns, mgr.DBConf.MaxOpenConns)
	if err != nil {
		panic(err)
	}
	orm.RegisterModel(new(entity.SysUser))
}

/*func (mgr *SysUserManager) GetPagedList(pageNum int, pageSize int) (pager utils.Pager) {
	o := orm.NewOrm()
	user := new([]models.SysUser)
	o.QueryTable("sys_user").Limit(pageSize, (pageNum-1)*pageSize).All(user)
	TotalCount, _ := o.QueryTable("sys_user").Count()
	pager.Total = int(TotalCount)
	pager.PageSize = pageSize
	pager.List = user
	fmt.Println(user)
	return pager
}*/

func (mgr *SysUserManager) GetPagedList(query query.SysUserQuery) (pager utils.Pager) {
	pageNum := query.PageNum
	pageSize := query.PageSize
	o := orm.NewOrm()
	user := new([]entity.SysUser)
/*	o.QueryTable("sys_user").Limit(pageSize, (pageNum-1)*pageSize).All(user)
	TotalCount, _ := o.QueryTable("sys_user").Count()*/
	table := o.QueryTable("sys_user")

	types := reflect.TypeOf(query)
	values := reflect.ValueOf(query)
	for i := 0; i < types.NumField(); i++ {
		// 获取每个成员的结构体字段类型
		fieldType := types.Field(i)
		if fieldType.Type.String() =="string" {
			v := values.Field(i).String()
			if v!=""{
				table = table.Filter(fieldType.Name,v)
			}
		}
		// 获取interface{}类型的值, 通过类型断言转换
		//fmt.Printf("name: %v  tag: '%v'  %v  %v\n", fieldType.Name, fieldType.Tag,fieldType.Type,)
	}
	table.Limit(pageSize, (pageNum-1)*pageSize).All(user)
	TotalCount, _ :=table.Count()

	pager.Total = int(TotalCount)
	pager.PageSize = pageSize
	pager.List = user
	fmt.Println(user)
	return pager

}

func (mgr *SysUserManager) GetList() interface{} {
	o := orm.NewOrm()
	user := new([]entity.SysUser)
	o.QueryTable("sys_user").All(user)
	return user
}

func (mgr *SysUserManager) DeleteById(id int) bool {
	o := orm.NewOrm()
	i, err := o.Delete(&entity.SysUser{Id: id})
	if err == nil {
		logs.Debug(i)
	}
	return i > 0
}

func (mgr *SysUserManager) SelectById(id int) (model entity.SysUser, err error) {
	o := orm.NewOrm()
	ob := entity.SysUser{Id: id}
	err = o.Read(&ob)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
		return
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
		return
	}
	return ob, err
	//Read 默认通过主键查询，可以使用指定的字段进行查询：
	/*user := User{Name: "slene"}
	err = o.Read(&user, "Name")*/
}

func (mgr *SysUserManager) Update(sysUser entity.SysUser) bool {
	orm.Debug = true
	var o orm.Ormer
	o = orm.NewOrm()
	if i, err := o.Update(&sysUser); err == nil {
		logs.Debug(i)
		return true
	}
	return false
}

func (mgr *SysUserManager) Insert(sysUser entity.SysUser) (model entity.SysUser) {
	var o orm.Ormer
	o = orm.NewOrm()
	_, err := o.Insert(&sysUser)
	if err == nil {
		model = sysUser
	}
	return model
}

// test

/**
根据某些字段来read 1：采用queryTable方法来查询 2：采用Raw执行sql语句
*/
func (mgr *SysUserManager) GetUsersByIdWithQueryTable(id string) (*[]entity.SysUser, error) {
	orm.Debug = true
	o := orm.NewOrm()
	user := new([]entity.SysUser)
	_, err := o.QueryTable("sys_user").Filter("Id", id).All(user)
	//err := o.QueryTable("user").Filter("id",key).One(user)
	//err:=o.Raw("select * from user where Id = ?",id).QueryRow(user)//使用sql语句进行查询
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return user, nil
}

//test vo

func testVO() {
	var users []vo.SysUserVO
	// 获取 QueryBuilder 对象. 需要指定数据库驱动参数。
	// 第二个返回值是错误对象，在这里略过
	qb, _ := orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.Select("user.name",
		"profile.age").
		From("user").
		InnerJoin("profile").On("user.id_user = profile.fk_user").
		Where("age > ?").
		OrderBy("name").Desc().
		Limit(10).Offset(0)
	// 导出 SQL 语句
	sql := qb.String()
	// 执行 SQL 语句
	o := orm.NewOrm()
	o.Raw(sql, 20).QueryRows(&users)
}

//事务方法
func testTrancation() {
	o := orm.NewOrm()
	_ = o.Begin()

	// 此过程中的所有使用 o Ormer 对象的查询都在事务处理范围内
	if true {
		_ = o.Rollback()
	} else {
		_ = o.Commit()
	}
}
