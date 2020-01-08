package main

import (
	"fmt"
	"net"
	"ssnbee/utils"
	"time"
)

type UserInfo struct {
	name    string
	C       chan []byte
	NewUser chan []byte
}

var (
	message     = make(chan []byte)
	onlineUsers = make(map[string]UserInfo)
)

func startServer() {

	listener, e := net.Listen("tcp", "127.0.0.1:9091")
	utils.AssertErr(e)
	defer listener.Close()
	fmt.Println("够浪聊天室-服务器已启动")

	fmt.Println("正在监听客户端连接请求……")
	go notifyAll()

	for {
		conn, e := listener.Accept()
		utils.AssertErr(e)
		fmt.Printf("地址为[%v]的客户端已连接成功\n", conn.RemoteAddr())
		go handleConnect(conn)
	}

}

func handleConnect(conn net.Conn) {
	defer conn.Close()
	// 管道overTime用于处理超时
	overTime := make(chan bool)

	buf1 := make([]byte, 4096)
	n, err := conn.Read(buf1)
	utils.AssertErr(err)
	userName := string(buf1[:n])
	perC := make(chan []byte)
	perNewUser := make(chan []byte)

	userInfo := UserInfo{name: userName, C: perC, NewUser: perNewUser}
	onlineUsers[conn.RemoteAddr().String()] = userInfo
	_, _ = conn.Write([]byte("💟💓💖💞💛你好," + userName + ",欢迎来到『够浪』™聊天室,请畅所欲言！💝💘💗💕💗"))
	utils.AssertErr(err)

	go func() {
		for _, v := range onlineUsers {
			v.NewUser <- []byte("🤵用户[" + userName + "]已加入当前聊天室\n")
		}
	}()
	//监听每位用户自己的channel
	go func() {
		for {
			select {
			case m1 := <-userInfo.C:
				conn.Write(m1)
			case m2 := <-userInfo.NewUser:
				conn.Write(m2)
			}
		}
	}()
	//循环读取客户端发来的消息
	go func() {
		buf2 := make([]byte, 4096)
		for {
			i, e := conn.Read(buf2)
			utils.AssertErr(e)
			thisUser := onlineUsers[conn.RemoteAddr().String()].name
			switch {
			case i == 0:
				fmt.Println(conn.RemoteAddr(), "已断开连接")
				for _, v := range onlineUsers {
					v.NewUser <- []byte("💨用户[" + thisUser + "]已退出当前聊天室\n")
				}
				delete(onlineUsers, conn.RemoteAddr().String())
			case string(buf2[:i]) == "who\n":
				conn.Write([]byte("当前在线用户:\n"))
				for _, v := range onlineUsers {
					conn.Write([]byte("🟢" + v.name + "\n"))
				}
			case len(string(buf2[:i])) > 7 && string(buf2[:i])[:7] == "rename|":
				onlineUsers[conn.RemoteAddr().String()] = UserInfo{name: string(buf2[:i-1])[7:], C: perC, NewUser: perNewUser}
				_, _ = conn.Write([]byte("您已成功修改用户名！\n"))
			}

			var msgByte []byte
			if buf2[0] != 10 && string(buf2[:i]) != "who\n" {
				if len(string(buf2[:i])) <= 7 || string(buf2[:i])[:7] != "rename|" {
					msgByte = append([]byte("💬["+thisUser+"]对大家说:"), buf2[:i]...)
				}
			} else {
				msgByte = nil
			}
			overTime <- true
			message <- msgByte
		}
	}()
	for {
		select {
		case <-overTime:
		case <-time.After(time.Second * 60):
			conn.Write([]byte("抱歉，由于长时间未发送聊天内容，您已被系统踢出"))
			thisUser := onlineUsers[conn.RemoteAddr().String()].name
			for _, v := range onlineUsers {
				if thisUser != "" {
					v.NewUser <- []byte("💨用户[" + thisUser + "]由于长时间未发送消息已被踢出当前聊天室\n")
				}
			}
			delete(onlineUsers, conn.RemoteAddr().String())
			return

		}
	}

}

func notifyAll() {
	for {
		select {
		case msg := <-message:
			for _, v := range onlineUsers {
				v.C <- msg
			}
		}
	}
}
