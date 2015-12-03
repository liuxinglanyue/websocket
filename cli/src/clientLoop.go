package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"strconv"
	"time"
)

var origin = "http://192.168.1.1/"
var url = "ws://localhost:8080/echo"

func main() {
	for i := 0; i < 5; i++ {
		go loop(strconv.Itoa(i))
	}
	time.Sleep(time.Duration(5) * time.Second)
}

func loop(i string) {
	ws, err := websocket.Dial(url, "", origin)
	defer ws.Close()
	if err != nil {
		log.Fatal(err)
	}

	message := []byte("hello, world! -- " + i)
	_, err = ws.Write(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(i+"-- Send: %s\n", message)

	var msg = make([]byte, 512)
	_, err = ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(i+"-- Receive: %s\n", msg)
}
