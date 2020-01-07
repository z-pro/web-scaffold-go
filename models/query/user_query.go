package query

import "ssnbee/base"

type UserQuery struct {
	base.BasePageQuery

	UserName string
	RealName string
	Phone string
}
