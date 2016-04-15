package main

import "github.com/gorilla/websocket"

//client represents a user who using the chat
type client struct {
	//socket is the websocket for this client
	socket *websocket.Conn

	//send is the channel that is recived messages
	send chan []byte

	//room is the chatroom that is joined in this client
	room *room
}

func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err != nil {
			c.room.foward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}
