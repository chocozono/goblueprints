package main

type room struct {
	//forward represents the channel which has some messages for transporting to another clients
	foward chan []byte
}
