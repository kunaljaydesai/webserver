package main

import (
	"log"
	"net"
)

func handleConnection(conn net.Conn) error {
	defer conn.Close()
	var request = make([]byte, 100)
	_, err := conn.Read(request)
	if err != nil {
		log.Println("failed to read request contents")
		return err
	}
	log.Println(string(request))
	_, err = conn.Write(request)
	if err != nil {
		log.Println("failed to write response contents")
		return err
	}
	return nil
}

func main() {
	ln, err := net.Listen("tcp", "0.0.0.0:8081")
	if err != nil {
		panic("error listening on port 8080")
	}
	for {
		conn, err := ln.Accept()
		log.Println("received connection")
		if err != nil {
			panic("failed to accept connection")
		}
		handleConnection(conn)
	}
}
