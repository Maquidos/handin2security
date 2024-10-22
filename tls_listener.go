package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"sync"
	"io"
)

func handleConnection(conn net.Conn, sharesChannel chan int) {
	defer conn.Close()
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		if err != io.EOF {
			log.Println("Read error: ", err)
		}
		return
	}
	var receivedShare int
	fmt.Sscanf(string(buf[:n]), "%d", &receivedShare)

	//Put share into the share channel 
	sharesChannel <- receivedShare
}

func listenForShares(certFile, keyFile, address string, sharesChannel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Make tls certificate
	cer, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatal("Error loading certificates: ", err)
	}
	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	// Create a listener using tcp
	listener, err := tls.Listen("tcp", address, config)
	if err != nil {
		log.Fatal("Error creating TLS listener: ", err)
	}
	defer listener.Close()

	// Listen and handle received messages
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn, sharesChannel)
	}
}
