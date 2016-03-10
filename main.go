package main

import (
	"fmt"

	"github.com/atsman/interviewr-go/db"
	"github.com/atsman/interviewr-go/handlers"
	"github.com/gin-gonic/gin"
	"github.com/googollee/go-socket.io"
)

const (
	Port = "3000"
)

func init() {
	db.Connect()
}

var Socketio_Server *socketio.Server

func socketHandler(c *gin.Context) {
	Socketio_Server.On("connection", func(so socketio.Socket) {
		fmt.Println("on connection")

		so.Join("chat")

		so.On("chat message", func(msg string) {
			fmt.Println("emit:", so.Emit("chat message", msg))
			so.BroadcastTo("chat", "chat message", msg)
		})
		so.On("disconnection", func() {
			fmt.Println("on disconnect")
		})
	})

	Socketio_Server.On("error", func(so socketio.Socket, err error) {
		fmt.Printf("[ WebSocket ] Error : %v", err.Error())
	})

	Socketio_Server.ServeHTTP(c.Writer, c.Request)
}

func main() {
	r := handlers.NewEngine()

	var err error
	Socketio_Server, err = socketio.NewServer(nil)
	if err != nil {
		panic(err)
	}

	r.GET("/socket.io", socketHandler)
	r.POST("/socket.io", socketHandler)
	r.Handle("WS", "/socket.io", socketHandler)
	r.Handle("WSS", "/socket.io", socketHandler)

	port := Port
	r.Run(":" + port)
}
