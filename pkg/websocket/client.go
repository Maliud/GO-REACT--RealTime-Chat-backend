package websocket

import (
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct{
	ID string
	Conn *websocket.Conn
	Pool *Pool
	mu sync.Mutex

}

type Message struct{
	Type int `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read(){
	defer func(){
		c.Pool.UnRegister <- c
		c.Conn.Close()
	}()

	for{
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil{
			log.Println(err)
			return
		}
		message := Message{Type: messageType, Body: string(p)}
		c.Pool.Broadcast <- message
		fmt.Printf("Mesaj Alındı:%v\n", message)
	}
}