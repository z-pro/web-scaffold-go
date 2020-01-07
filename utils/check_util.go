package utils

import (
	"log"
	"os"
)

func AssertErr(err error)  {
	if err!=nil{
		//初始化日志服务
		logger := log.New(os.Stdout, "[ErrorLog]", log.Lshortfile | log.Ldate | log.Ltime)
		// 启动web服务，监听1010端口
		go func() {
			logger.Fatal("ListenAndServe: ", err)
		}()
		panic(err)
	}
}

func HasErr(err error) bool {
	return err!=nil
}