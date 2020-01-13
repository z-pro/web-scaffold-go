package produce

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/nsqio/go-nsq"
	"time"
)

func initxxx() {
	fmt.Println("p")
	for i := 0; i < 10; i++ {
		SendMessage()
	}
	time.Sleep(time.Second * 10)
}

func SendMessage() {
	url := beego.AppConfig.String("mq_host")
	producer, err := nsq.NewProducer(url, nsq.NewConfig())
	if err != nil {
		panic(err)
	}
	err = producer.Publish("test", []byte("hello world"))
	if err != nil {
		panic(err)
	}
	producer.Stop()
}
