package main

import (
	"fmt"
	"net/http"

	"github.com/Maliud/GO-REACT--RealTime-Chat/pkg/websocket"
)




func serveWS(pool *websocket.Pool, w http.ResponseWriter, r *http.Request){
	fmt.Println("websocket uç noktasına ulaşıldı")
	conn, err := websocket.Upgrade(w,r)
	if err != nil{
		fmt.Fprintf(w, "%+v\n", err)
	}
	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	client.Read()
}


func setupRoutes(){
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request){
		serveWS(pool, w,r)
	})
}

func main(){
	fmt.Println("MuhammedAli Ud Full Stack Project")
	setupRoutes()
	http.ListenAndServe(":9000", nil)
}