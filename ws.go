package main 

import (
	"net/http"
	"github.com/gorilla/websocket"
)

var (
	//跨域
	upgrader = websocket.Upgrader{
		CheckOrigin:func (r *http.Request) bool {
			return true
		},
	}
)

// func wsHandle(w http.ResponseWriter, r *http.Request) {
// 	
// 	conn, err := upgrader.Upgrade(w,r, nil)
// 	if err != nil{
// 		return
// 	}

// 	for {
// 		_, data, err := conn.ReadMessage()
// 		if err != nil {
// 			goto ERR
// 		}
// 		//text ,binary
// 		err := conn.WriteMessage(websocket.TextMessage, data)
// 		if err != nil {
// 			goto ERR
// 		}
// 	}
// 	ERR:
// 	conn.Close()
// }

func wsHandle(w http.ResponseWriter, r *http.Request) {

	var(
		conn *websocket.Conn
		err error
		data []byte
	)

	
	if conn, err = upgrader.Upgrade(w,r, nil); err != nil{
		return
	}

	if _,data, err = conn.ReadMessage(); err != nil{
		goto ERR
	}
	if err = conn.WriteMessage(websocket.TextMessage, data); err != nil{
		goto ERR
	}

	ERR:
		conn.Close()
}

func main() {
	
	http.HandleFunc("/ws", wsHandle)

	http.ListenAndServe("0.0.0.0:8888", nil)
}
