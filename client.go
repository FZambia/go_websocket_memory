package main

import (
	"flag"
	"log"

	"github.com/gorilla/websocket"
)

var (
	nConns = flag.Int("conns", 10000, "number of websocket connections")
)

func main() {
	flag.Parse()

	var conns []*websocket.Conn
	for i := 0; i < *nConns; i++ {
		c, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8000", nil)
		if err != nil {
			log.Fatalf("Failed to connect %d: %v", i, err)
		}
		go func(conn *websocket.Conn) {
			for {
				for {
					_, _, err := conn.ReadMessage()
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}(c)
		conns = append(conns, c)
		log.Printf("conn %d established", i)
	}

	log.Println("all connections established")
	select {}
}
