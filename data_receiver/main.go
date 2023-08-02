package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"tolling/types"
)

func main() {
	recv := NewDataReceiver()
	http.HandleFunc("/ws", recv.handleWs)
	if err := http.ListenAndServe(":30000", nil); err != nil {
		log.Fatal(err)
	}

}

type DataReceiver struct {
	msgChan chan types.OBUData
	conn    *websocket.Conn
}

func NewDataReceiver() *DataReceiver {
	return &DataReceiver{
		msgChan: make(chan types.OBUData, 128),
	}
}

func (dr *DataReceiver) handleWs(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{
		WriteBufferSize: 1028,
		ReadBufferSize:  1028,
	}
	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	dr.conn = conn
	go dr.WsReceiveLoop()
}
func (dr *DataReceiver) WsReceiveLoop() {
	fmt.Println("New OBU Connected!")
	for {
		var data types.OBUData
		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Println("Receive Error ", err)
			continue
		}
		fmt.Printf("%+v\n", data)
		dr.msgChan <- data
	}
}
