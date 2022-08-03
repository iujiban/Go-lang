package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//ResponseWrite와 관련성을 받아들이는 함수를 build HTTP 응답.
//Upgrader 기능을 사용하여 HTTP 연결을 WebSocket protocol로 변환

func echo(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
		return
	}
	defer conn.Close()
	// 메시지를 수신 대기하고 읽습니다.
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
			fmt.Println(err)
			break
		}
		// testing what is receving the message
		fmt.Printf("recv: %s", message)
		err = conn.WriteMessage(mt, message)

		if err != nil {
			log.Fatal(err)
			fmt.Println(err)
			break
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	//writing, requesting, index.html
	http.ServeFile(w, r, "index.html")

}
func main() {
	//핸들러 등록
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)
	//port
	http.ListenAndServe(":8080", nil)
}
