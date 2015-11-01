package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"strconv"
)

type Server struct {
	port int `json:"port"`
}

func main() {
	server := Server{1337}
	listener, _ := net.Listen("tcp", ":"+strconv.Itoa(server.port))
	fmt.Println("Server running on", server.port)

	// accepts connections
	c, _ := listener.Accept()

	// infite loop listening for new connecitons
	for {
		message, _ := bufio.NewReader(c).ReadString("\n")
		go handleMessage(message) // go routine swag
	}
}

func handleConnection(c net.Conn) {
	// data is payload of packet
	fmt.Println("")
	data, _ := ioutil.ReadAll(c)
	m := Decode(bytes.NewBuffer(data))
	handleMessage(m, c)
}

func handleMessage(message *Message, conn net.Conn) {
	newmessage := strings.ToUpper(message)
	c.Write([]byte(newmessage + "\n"))
}
