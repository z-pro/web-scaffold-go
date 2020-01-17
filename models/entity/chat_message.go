package entity

import (
	"fmt"
	"time"
)

type JsonTime time.Time

//实现它的json序列化方法
func (this JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

type ChatMessage struct {
	Username string   `json:"username""`
	Message  string   `json:"message""`
	CreateOn JsonTime `json:"createon""`
}
