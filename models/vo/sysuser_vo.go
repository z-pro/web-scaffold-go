package vo

import "time"

type SysUserVO struct {
	Id   int
	RealName string
	UserName string
	CreateDate time.Time
}