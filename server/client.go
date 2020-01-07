package main

import (
	"net"
	"ssnbee/utils"
)

func mainT() {
	conn, e := net.Dial("tcp", "127.0.0.1:9001")
	utils.AssertErr(e)
	defer conn.Close()
	conn.Write([]byte("你好"))
}
