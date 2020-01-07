package main

import (
	"io"
	"log"
	"net"
	"os"
	"ssnbee/utils"
)

func StartClient() {

	conn, e := net.Dial("tcp", "127.0.0.1:8787")
	utils.AssertErr(e)
	defer conn.Close()

	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn)
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}



