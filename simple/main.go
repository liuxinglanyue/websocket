package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

func echoHandler(ws *websocket.Conn) {

	msg := make([]byte, 512)
	n, err := ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive: %s\n", msg[:n])

	buf := new(bytes.Buffer)
	buf.WriteString("server: ")
	buf.WriteString(string(msg[:n]))
	m, err := ws.Write(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send: %s\n", msg[:m])
}

func main() {
	http.Handle("/websocket", websocket.Handler(echoHandler))
	http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServe(":1889", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
