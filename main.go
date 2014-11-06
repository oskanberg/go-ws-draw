package gowebsockets

import (
	"fmt"
	"net/http"
	"os"
)

// this makes me so sad
var filepath = os.Getenv("GOPATH") + "/src/github.com/oskanberg/go-ws-draw/home.html"

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, filepath)
}

func ServeForever(port int) {
	go h.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", serveWs)
	addr := fmt.Sprintf(":%d", port)
	http.ListenAndServe(addr, nil)
}
