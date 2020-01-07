package query

import "ssnbee/base"

type SysUserQuery struct {
	base.BasePageQuery

	UserName string
	RealName string
	Phone string
}
