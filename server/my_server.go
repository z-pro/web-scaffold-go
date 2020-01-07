package main

import (
	"bufio"
	"fmt"
	"net"
	"ssnbee/utils"
)

func StartServer() {
	listener, err := net.Listen("tcp", "127.0.0.1:8787")
	utils.AssertErr(err)
	defer listener.Close()

	go broadcaster()
	for {
		conn, err := listener.Accept()
		utils.AssertErr(err)
		go processMsg(conn)
	}
}

type client chan<- string
var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				fmt.Println("1.1")
				//who := strings.Split(cli., "")[0]
				cli <- msg
			}
			fmt.Println("1")
		case cliEnter := <-entering:
			clients[cliEnter] = true
			fmt.Println("2")
		case cliExit := <-leaving:
			delete(clients, cliExit)
			close(cliExit)
			fmt.Println("3")
		}

	}

}

func processMsg(conn net.Conn) {
	defer conn.Close()
	ch := make(chan string) // 对外发送客户消息的通道
	go clientWriter(conn, ch)
	who := conn.RemoteAddr().String()
	ch <- "欢迎 " + who
	messages <- who + "上线"
	entering <- ch

	scanner := bufio.NewScanner(conn)

	for scanner.Scan(){
		messages<-who+":"+scanner.Text()
	}

	leaving<-ch
	messages<-who+"下线"
}

func clientWriter(conn net.Conn, ch <-chan string) {

	for msg := range ch {
		fmt.Println("qipa")
		fmt.Fprintln(conn, msg+" over") // 注意：忽略网络层面的错误
	}

}
