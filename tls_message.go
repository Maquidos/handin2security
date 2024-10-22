package main

import (
	"crypto/tls"
	"log"
	"fmt"
	"sync"
)

// TLS client to send the local sum to the hospital
func sendResultToHospital(certFile, address string, localSum int) {
	config := &tls.Config{
		InsecureSkipVerify: true, // Skip verification for testing purposes
	}
	conn, err := tls.Dial("tcp", address, config)
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer conn.Close()

	message := fmt.Sprintf("%d", localSum) // Send the local sum as a string
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatal("Write error:", err)
	}
}

func sendShareToOtherPatient(certFile, address string, share int, wg *sync.WaitGroup){
	defer wg.Done()
	config := &tls.Config{
		InsecureSkipVerify: true, // Skip verification for testing purposes
	}
	conn, err := tls.Dial("tcp", address, config)
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer conn.Close()

	message := fmt.Sprintf("%d", share) // Send the share as a string
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatal("Write error:", err)
	}
}
