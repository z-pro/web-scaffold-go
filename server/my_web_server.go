package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	// log "github.com/Sirupsen/logrus"
	// "github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//NewHub 分配一个新的Hub
func NewHub(onEvent OnMessageFunc) *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),      //包含要想向前台传递的数据
		register:   make(chan *Client),     //有新的连接，将放入这里
		unregister: make(chan *Client),     //断开连接加入这
		clients:    make(map[*Client]bool), //包含所有的client连接信息
		OnMessage:  onEvent,
	}
}

//OnMessageFunc 接收到消息触发的事件
type OnMessageFunc func(message []byte, hub *Hub) error

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

//Hub maintains the set of active clients and broadcasts messages to the
//clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	//Broadcast 公告消息队列
	Broadcast chan []byte
	//OnMessage 当收到任意一个客户端发送到消息时触发
	OnMessage OnMessageFunc

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod) //设置定时
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send: //这里send里面的值时run里面获取的，在这里才开始实际向前台传值
			c.conn.SetWriteDeadline(time.Now().Add(writeWait)) //设置写入的死亡时间，相当于http的超时
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{}) //如果取值出错，关闭连接，设置写入状态，和对应的数据
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage) //以io的形式写入数据，参数为数据类型
			if err != nil {
				return
			}
			w.Write(message) //写入数据，这个函数才是真正的想前台传递数据

			if err := w.Close(); err != nil { //关闭写入流
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait)) //心跳包，下面ping出错就会报错退出，断开这个连接
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	log.Println("star read message")
	defer func() {
		c.hub.unregister <- c //读取完毕后注销该client
		c.conn.Close()
		log.Println("websocket Close")
	}()
	c.conn.SetReadLimit(maxMessageSize)              //设置最大读取容量
	c.conn.SetReadDeadline(time.Now().Add(pongWait)) //设置读取死亡时间
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Println("websocket.IsUnexpectedCloseError:", err)
			}
			log.Println("websocket read message have err")
			break
		}
		if c.hub.OnMessage != nil {
			if err := c.hub.OnMessage(message, c.hub); err != nil {
				log.Println(err)
				break
			}
		}
	}
}

//Run 开始消息读写队列，无限循环，应该用go func的方式调用
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register: //客户端有新的连接就加入一个
			h.clients[client] = true
		case client := <-h.unregister: //客户端断开连接，client会进入unregister中，直接在这里获取，删除一个
			if _, ok := h.clients[client]; ok { //找到对应需要删除的client
				delete(h.clients, client) //在map中根据对应value值，使用delete删除对应client
				close(client.send)        //关闭对应连接
			}
		case message := <-h.Broadcast: //将数据发给连接中的send，用来发送
			for client := range h.clients { //clients中保存了所有的客户端连接，循环所有连接给与要发送的数据
				select {
				case client.send <- message: //将需要发送的数据放入send中，在write函数中实际发送
				default:
					close(client.send)        //关闭发送通道
					delete(h.clients, client) //删除连接
				}
			}
		}
	}
}

// ServeWs handles websocket requests from the peer.
func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil) //返回一个websocket连接
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("ServeWs is star")
	//生成一个client，里面包含用户信息连接信息等信息
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client //将这个连接放入注册，在run中会加一个
	go client.writePump()         //新开一个写入，因为有一个用户连接就新开一个，相互不影响，在内部实现心跳包检测连接，详细看函数内部

	client.readPump() //读取websocket中的信息，详细看函数内部

}
