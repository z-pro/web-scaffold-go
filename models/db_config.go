package models


/**
*数据相关配置
 */
type DBConfig struct {
	Host         string
	Port         string
	Database     string
	Username     string
	Password     string
	MaxIdleConns int   //最大空闲连接
	MaxOpenConns int   //最大连接数
}
