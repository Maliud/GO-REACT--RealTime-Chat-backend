package websocket

import "fmt"

type Pool struct {
	Register   chan *Client
	UnRegister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		UnRegister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func(pool *Pool) Start(){
	for{
		select{
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("bağlantı havuzunun boyutu:", len(pool.Clients))
			for client, _ := range pool.Clients{
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: 1, Body: "Yeni Kullanıcı Bağlandı..."})
			}
			break
		case client := <-pool.UnRegister:
			delete(pool.Clients, client)
			fmt.Println("Bağlantı Havuzunun boyutu:", len(pool.Clients))
			for client, _ := range pool.Clients{
				client.Conn.WriteJSON(Message{Type: 1, Body: "Kullanıcı Ayrıldı..."})
			}
			break
		case message := <-pool.Broadcast:
			fmt.Println("Havuzdaki tüm istemcilere mesaj gönderiliyor...")
			for client, _ := range pool.Clients{
				if err := client.Conn.WriteJSON(message); err != nil{
					fmt.Println(err)
					return
				}
			}
		}
	}
}