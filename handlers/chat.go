package handlers

import (
	"fmt"

	"github.com/atsman/interviewr-go/models"
	"github.com/gin-gonic/gin"
	"github.com/googollee/go-socket.io"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("chat")

type CodeSharingState struct {
	Lang   string `json:"lang"`
	Code   string `json:"code"`
	RoomId string `json:"roomId"`
	Cursor Cursor `json:"cursor"`
}

type Cursor struct {
	Column int `json:"column"`
	Row    int `json:"row"`
}

// GetSocketHandler - Returns socket handler
func GetSocketHandler(socketServer *socketio.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		socketServer.On("connection", func(so socketio.Socket) {
			fmt.Println("on connection")

			so.On("joinRoom", func(roomId string) {
				log.Debugf("joinRoom, roomId: %s", roomId)
				so.Join(roomId)
			})

			so.On("sendMessage", func(message models.Message) {
				log.Debug("sendMessage, message: ", message)
				so.BroadcastTo(message.RoomID, "newMessage", message)
			})

			so.On("sendCode", func(code CodeSharingState) {
				log.Debug("sendCode, codeState:", code)
				so.BroadcastTo(code.RoomId, "receiveCodeChange", code)
			})

			so.On("disconnection", func() {
				fmt.Println("on disconnect")
			})
		})

		socketServer.On("error", func(so socketio.Socket, err error) {
			fmt.Printf("[ WebSocket ] Error : %v", err.Error())
		})

		socketServer.ServeHTTP(c.Writer, c.Request)
	}
}
