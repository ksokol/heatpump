package heatpump

import (
	"context"
	"log"
	"net"
	"time"
)

type SocketConnection struct {
	connection net.Conn
}

func (so *SocketConnection) Write(values ...int32) (int, error) {
	p := make([]byte, 0)
	for _, value := range values {
		p = append(p, toBytes(value)...)
	}
	return so.connection.Write(p)
}

func (so *SocketConnection) Read() (int32, error) {
	p := make([]byte, 4)
	_, err := so.connection.Read(p)
	var v int32
	if err == nil {
		v = byteToInt32(p)
	}
	return v, err
}

func (so *SocketConnection) Close() {
	address := so.connection.RemoteAddr()
	_ = so.connection.Close()
	log.Printf("connection to %v closed", address)
}

func NewSocketConnection(address string) (SocketConnection, error) {
	var dialer net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	connection, err := dialer.DialContext(ctx, "tcp", address)

	if err == nil {
		log.Printf("connection to %v open", connection.RemoteAddr())
		err = connection.SetDeadline(time.Now().Add(1 * time.Second))
	}
	return SocketConnection{connection}, err
}

func toBytes(v int32) []byte {
	return []byte{
		byte(v >> 24 & 255),
		byte(v >> 16 & 255),
		byte(v >> 8 & 255),
		byte(v >> 0 & 255),
	}
}

func byteToInt32(b []byte) int32 {
	return (int32(b[0]) << 24) + (int32(b[1]) << 16) + (int32(b[2]) << 8) + (int32(b[3]) << 0)
}
