package main

import (
	"log"
	"net/http"
)

var (
	hub = NewHub(nil) //新建一个用户
)

func init() {

	go hub.Run() //开始获取用户中传送的数据

	http.HandleFunc("/echo", func(res http.ResponseWriter, r *http.Request) {
		ServeWs(hub, res, r)
	})
	DataTest()
}

func main() {
	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		log.Panic(err)
	}
}

func DataTest() {
	hub.Broadcast <- []byte(string("this is return message"))
	fn := func(message []byte, hub *Hub) error {
		log.Println("message:", string(message))
		return nil
	}
	hub.OnMessage = fn
}
