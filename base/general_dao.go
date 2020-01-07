package base

import (
	"ssnbee/utils"
)

type GeneralDao interface {
	GetPagedList(pageNum int, pageSize int) (pager utils.Pager)
	GetList() interface{}
	DeleteById(id int) bool
	Update(sysUser interface{}) bool
	Insert(sysUser interface{}) (model interface{})
}
