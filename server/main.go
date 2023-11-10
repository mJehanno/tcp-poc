package main

import (
	"fmt"
	"io"
	"net"

	"github.com/charmbracelet/log"
)

func main() {
	l, err := net.Listen("tcp", ":1500")
	handleErr("fatal", "couldn't start tcp server", err)
	log.Infof("Server ready ! Listening on 1500")

	for {
		con, err := l.Accept()
		handleErr("fatal", "couldn't accept incoming tcp connection", err)

		log.Infof("Received tcp connection from %s", con.RemoteAddr().String())

		go handleConnection(con)
	}

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

func handleConnection(con net.Conn) {
	defer con.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := con.Read(buffer)
		if err != nil && err == io.EOF {
			log.Info("Socket was close")
			break
		}
		handleErr("error", "couldn't read message", err)
		log.Infof("Received the current message : %s", buffer[:n])
	}

}
