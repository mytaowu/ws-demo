package ws

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var ug = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	}, // 用于测试，所有链接
}

// EchoMsg 初始化处理函数
func EchoMsg(w http.ResponseWriter, r *http.Request) {
	conn, err := ug.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("ug upgrade error: +%v\n", err)
		return
	}
	defer conn.Close()

	for {
		// 读取客户端消息
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("conn read message error: %+v\n", err)
			return
		}

		// 把消息打印到标准输出
		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		// 把消息写回客户端，完成回音
		if err = conn.WriteMessage(msgType, msg); err != nil {
			fmt.Printf("conn writer message error: %+v\n", err)
			return
		}
	}
}
