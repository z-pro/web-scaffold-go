package enums

type SysUserTypeEnum int32

const (
	ADMIN  SysUserTypeEnum = 1
	NORMAL SysUserTypeEnum = 2
)

func (userTypeEnum SysUserTypeEnum) String() string {
	switch userTypeEnum {
	case ADMIN:
		return "管理员"
	case NORMAL:
		return "普通用户"
	default:
		return "UNKNOWN"
	}
}
