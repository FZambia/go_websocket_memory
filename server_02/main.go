package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	_ "net/http/pprof"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  0,
	WriteBufferSize: 0,
}

// Client is a wrapper over ws connection.
type Client struct {
	conn *websocket.Conn
}

func newClient(conn *websocket.Conn) *Client {
	return &Client{conn}
}

func (c *Client) write() {
	select {}
}

func (c *Client) read() {
	for {
		_, _, err := c.conn.ReadMessage()
		if err != nil {
			return
		}
	}
}

func serveWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()
	client := newClient(conn)
	go client.write()
	client.read()
}

func main() {
	go func() {
		for {
			time.Sleep(5 * time.Second)
			printMemUsage()
		}
	}()
	http.HandleFunc("/", serveWS)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

func printMemUsage() {
	runtime.GC()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tStackInuse = %v MiB", bToMb(m.StackInuse))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
