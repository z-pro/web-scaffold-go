package conf

import (
	"os"
)

func initProdConf() {
	LogicConf = logicConf{
		MySQL:                  "root:123456@tcp(localhost:3306)/blog?charset=utf8&parseTime=true",
		NSQIP:                  "127.0.0.1:4150",
		RedisIP:                "127.0.0.1:6379",
		RPCIntListenAddr:       ":50000",
		ClientRPCExtListenAddr: ":50001",
		ServerRPCExtListenAddr: ":50002",
		ConnRPCAddrs:           "addrs:///127.0.0.1:60000,127.0.0.1:60001",
	}

	ConnConf = connConf{
		TCPListenAddr: ":8080",
		RPCListenAddr: ":60000",
		LocalAddr:     "127.0.0.1:60000",
		LogicRPCAddrs: "addrs:///127.0.0.1:50000",
	}

	WSConf = wsConf{
		WSListenAddr:  ":8081",
		RPCListenAddr: ":60001",
		LocalAddr:     "127.0.0.1:60001",
		LogicRPCAddrs: "addrs:///127.0.0.1:50000",
	}

	/*logger.Leavel = zap.InfoLevel
	logger.Target = logger.File*/
	_, _ = os.OpenFile("e:/beeapp.log", os.O_RDWR|os.O_APPEND, 0644)
	//logger := log.New(os.Stdout, "[ErrorLog]", log.Lshortfile | log.Ldate | log.Ltime)

}
