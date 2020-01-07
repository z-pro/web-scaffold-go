package main

import (
	"fmt"
	"net"
	"ssnbee/utils"
)

func processData(conn net.Conn) {
	bytes := make([]byte, 1024)
	for {
		n, err := conn.Read(bytes)
		utils.AssertErr(err)
		if n>0{
			fmt.Printf(string(bytes))
		}
	}
}

func mainT2() {

	/*defer func() {

		recover()

	}()
	*/
	fmt.Print("running")
	listener, e := net.Listen("tcp", "127.0.0.1:9001")
	utils.AssertErr(e)
	defer listener.Close()

	conn, e := listener.Accept()
	utils.AssertErr(e)
	defer conn.Close()

	go processData(conn)
}
