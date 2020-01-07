package entity

import "time"

type User struct {
	Id   int `form:"id" orm :"auto,pk"` //默认//Id   int `orm:"pk;auto;column(user_id)"` //表示设置为主键并且自增，列名为user_id
	RealName string `orm:"size(100)" json:"realName" form:"realName"`
	UserName string `orm:"size(100)"column(user_name)`
	Password string `orm:"size(100)"column(password)`
	Phone string `orm:"size(100)"`
	Avatar string `orm:"size(200)"`
	Sex  int `orm:"size(2);column(gender)"`
	CreateDate time.Time `orm:"type(date);column(create_date)"`
	LoginDate time.Time `orm:"type(date);column(login_date)"`
	Remark string `orm:"size(200)"column(remark)`
}
