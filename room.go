package main

type room struct {
	//forward represents the channel which has some messages for transporting to another clients
	foward chan []byte

	//join represents the channel for the client who wanna join to chat room
	join chan *client

	//leave represents the channel for the client who wanna leave from chat room
	leave chan *client

	//clients holds *client
	clients map[*client]bool
}
