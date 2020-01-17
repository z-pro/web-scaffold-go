package main

import (
	"encoding/json"
	"log"
	"net/http"
	"ssnbee/models/entity"
	"time"
)

var (
	hub = NewHub(nil) //新建一个用户
)

func init() {

	go hub.Run() //开始获取用户中传送的数据

	http.HandleFunc("/echo", func(res http.ResponseWriter, r *http.Request) {
		ServeWs(hub, res, r)
	})
	SendChat()
}

func main() {
	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		log.Panic(err)
	}
}

func SendChat() {
	hub.Broadcast <- []byte(string("this is return message"))
	fn := func(message []byte, hub *Hub) error {
		msg := new(entity.ChatMessage)
		json.Unmarshal(message, &msg)
		msg.CreateOn = entity.JsonTime(time.Now())
		bytes, _ := json.Marshal(msg)
		for cli := range hub.clients {
			cli.send <- bytes
			/*err := websocket.Message.Send(key, data)
			if err != nil{
				// 移除出错的连接
				delete(users, key)
				fmt.Println("发送出错: " + err.Error())
				break
			}*/
		}
		log.Println("message:", string(bytes))
		return nil
	}
	hub.OnMessage = fn
}
