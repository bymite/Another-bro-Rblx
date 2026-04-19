package main

import (
	"encoding/binary"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" { port = "1080" }
	l, err := net.Listen("tcp", ":"+port)
	if err != nil { log.Fatal(err) }
	log.Printf("UDP-over-TCP Server active on :%s\n", port)

	for {
		conn, err := l.Accept()
		if err != nil { continue }
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 10)
	if _, err := io.ReadFull(conn, buf[:2]); err != nil { return }
	conn.Write([]byte{0x05, 0x00}) // No Auth

	if _, err := io.ReadFull(conn, buf[:4]); err != nil { return }
	if buf[1] == 0x03 { // UDP ASSOCIATE
		conn.Write([]byte{0x05, 0x00, 0x00, 0x01, 0, 0, 0, 0, 0, 0})
		runTunnel(conn)
	}
}

func runTunnel(conn net.Conn) {
	udpConn, _ := net.ListenPacket("udp", ":0")
	defer udpConn.Close()
	
	// Relay logic: TCP Frame <-> UDP Packet
	go func() {
		for {
			var length uint16
			if err := binary.Read(conn, binary.BigEndian, &length); err != nil { return }
			data := make([]byte, length)
			if _, err := io.ReadFull(conn, data); err != nil { return }
			// Logic to parse SOCKS5 header and sendto() would go here
		}
	}()
    select {} 
}
