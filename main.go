package main

import (
	"github.com/atsman/interviewr-go/db"
	"github.com/atsman/interviewr-go/handlers"
	"github.com/googollee/go-socket.io"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("main")

const (
	Port = "8123"
)

func init() {
	db.Connect()
}

var Socketio_Server *socketio.Server

func main() {
	r := handlers.NewEngine()

	var err error
	Socketio_Server, err = socketio.NewServer(nil)
	if err != nil {
		panic(err)
	}

	socketHandler := handlers.GetSocketHandler(Socketio_Server)

	r.GET("/socket.io", socketHandler)
	r.POST("/socket.io", socketHandler)
	r.Handle("WS", "/socket.io", socketHandler)
	r.Handle("WSS", "/socket.io", socketHandler)

	port := Port
	r.Run(":" + port)
}
