package main

import (
	"fmt"
	"net"

	"github.com/charmbracelet/log"
)

func main() {
	con, err := net.Dial("tcp", "127.0.0.1:1500")
	handleErr("fatal", "couldn't connect to remote server", err)

	con.Write([]byte("Hello from the client"))
}

func handleErr(level, msg string, err error) {
	if err != nil {
		err = fmt.Errorf("%s : %w", msg, err)
		switch level {
		case "fatal":
			log.Fatal(err)
		case "error":
			log.Error(err)
		}
	}
}
