package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Connection struct {
	// The websocket Connection.
	ws *websocket.Conn
	// channel for objs to jsonify and send
	send chan interface{}
}

func (c *Connection) WritePump() {
	defer c.ws.Close()
	for msg := range c.send {
		err := c.ws.WriteJSON(msg)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func (c *Connection) ReadPump() {
	defer c.ws.Close()
	var obj interface{}
	for {
		err := c.ws.ReadJSON(&obj)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(obj)
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving WebSocket connection")
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}
	c := &Connection{send: make(chan interface{}), ws: ws}
	go c.WritePump()
	c.ReadPump()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "home.html")
	})
	http.HandleFunc("/ws", serveWs)
	http.ListenAndServe(":8080", nil)
}
