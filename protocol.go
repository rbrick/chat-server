package main

import "bytes"

const (
// Connection related things
    PING int = iota // 0
    LOGIN                  // 1
    DISCONN                // 2
// Chat
    CHAT                   // 3

)

type Client struct {
    state int // The state of the client
    
}

type Message struct {
    packetType int
    payload []byte
}

// Encodes a message into a slice of bytes
func (m *Message) Encode() []byte {
    encoded := new(bytes.Buffer)
    encoded.WriteByte(byte(m.packetType))
    encoded.Write(m.payload)
    return encoded.Bytes();
}

func Decode(buffer *bytes.Buffer) *Message {
    id, _ := buffer.ReadByte()
    data := buffer.Bytes()
    return &Message{int(id), data}
}
