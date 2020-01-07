package main

import (
	"fmt"
	"net"
	"os"
	"ssnbee/utils"
)

func startClient()  {

	conn, e := net.Dial("tcp", "127.0.0.1:9091")
	utils.AssertErr(e)
	defer conn.Close()
	fmt.Println("连接服务器成功")
	fmt.Println("先起一个名字吧：")
	var userName string
	//使用Scan输入，不允许出现空格
	_, _ = fmt.Scan(&userName)
	_, _ = conn.Write([]byte(userName))
	buf2:=make([]byte,4096)
	n, err := conn.Read(buf2)
	utils.AssertErr(err)

	// 客户端收到“你好，***，欢迎来到够浪聊天室，请畅所欲言！”
	fmt.Println(string(buf2[:n]))
	fmt.Println("⚠提示:长时间没有发送消息会被系统强制踢出")


	go func() {
		for {
			buffer1:=make([]byte,4096)
			//这里使用Stdin标准输入，因为scanf无法识别空格
			n,err:=os.Stdin.Read(buffer1)
			utils.AssertErr(err)
			conn.Write(buffer1[:n])
		}

	}()

	for  {
		buffer2:=make([]byte,4096)
		i, e := conn.Read(buffer2)
		utils.AssertErr(e)
		if i==0{
			fmt.Println("服务器已关闭当前连接，正在退出……")
			return
		}
		fmt.Print(string(buffer2[:n]))

	}
}

