package main

import (
    "fmt"
    "bytes"
    "net"
    "strconv"
    "io/ioutil"
)


type Server struct {
    port int `json:"port"`
}

func main() {
    server := Server{1337}
    listener, _ := net.Listen("tcp", ":" + strconv.Itoa(server.port))
    fmt.Println("Server running on", server.port)
    for {
        c, _ := listener.Accept()
        go handleConnection(c) // goroutine :D!
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
    fmt.Println(message.packetType, "received from", conn.RemoteAddr().String())
}

