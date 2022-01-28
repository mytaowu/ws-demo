package main

import (
	"net/http"

	"websocket_demo/ws"
)

func main() {
	http.HandleFunc("/ws/echo", ws.EchoMsg)

	http.ListenAndServe(":8080", nil)
}
