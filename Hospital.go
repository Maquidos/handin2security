package main

import (
	"fmt"
	"sync"
)

func startHospital(R int) error {

	outChannel := make(chan int, 3) // Expecting 3 output from the three patients
	var wg sync.WaitGroup

	// Start TLS listener 
	address := ":8444" // Unique port for hospital
	wg.Add(1)
	go listenForShares("cert.pem", "key.pem", address, outChannel, &wg)

	out1 := <-outChannel
	out2 := <-outChannel
	out3 := <-outChannel

	
	out := (out1 + out2 + out3) % (R+1)
	fmt.Println("Total out:", out)

	wg.Wait()
	return nil
}