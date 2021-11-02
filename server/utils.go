package main

import (
	"fmt"

	socketio "github.com/googollee/go-socket.io"
)

func DisconectAfk(users []User) {
	// TODO
}

func CheckUsernameAndRoom(s socketio.Conn) error {
	fmt.Printf("%v", s.Namespace())
	return nil
}
