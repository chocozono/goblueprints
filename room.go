package main

type room struct {
	// forward represents the channel which has some messages for transporting to another clients
	foward chan []byte

	// join represents the channel for the client who wanna join to chat room
	join chan *client

	// leave represents the channel for the client who wanna leave from chat room
	leave chan *client

	// clients holds *client
	// treating this map exepts using channel is to be deprecated because
	// there is a possibility that some goroutines change at the same time
	clients map[*client]bool
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			// join
			r.clients[client] = true
		case client := <-r.leave:
			// leave
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.foward:
			// send messages to all clients
			for client := range r.clients {
				select {
				case client.send <- msg:
					// send a message
				default:
					// if failed sending message
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}
