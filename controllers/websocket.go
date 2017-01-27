package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/gorilla/websocket"
	"fmt"
)

type (
	WebSocketController struct{}
)

var ws_con *websocket.Conn

func NewWebSocketController() *WebSocketController{
	return &WebSocketController{}
}

func (ws WebSocketController) Handle(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	var err error
	ws_con, err = websocket.Upgrade(w,r,w.Header(),1024,2014)
	if err!=nil {
		fmt.Println("Could not open websocket connection")
	}else{
		fmt.Println("Web socket client connected")
	}
}