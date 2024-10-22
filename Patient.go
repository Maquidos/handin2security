package main

import (
	"fmt"
	"sync"
	"time"
)

func startPatient(name string, R int , secret int, cert string, key string, ownAddress string, p1Address string, p2Address string ) error {

	// Generate shares
	
	shares := generateShares(secret, R)
	//fmt.Println(name, "Shares:",shares.Share1, shares.Share2, shares.Share3)

	// Channel to receive shares
	sharesChannel := make(chan int, 2) // Expecting 2 shares from other participants
	var wg sync.WaitGroup

	// Start TLS listener
	wg.Add(1)
	go listenForShares(cert, key, ownAddress, sharesChannel, &wg)

	time.Sleep(2 * time.Second)

	// Send shares to patient 1 and patient 2
	wg.Add(2)
	go sendShareToOtherPatient("cert.pem", p1Address, shares.Share2, &wg)
	go sendShareToOtherPatient("cert.pem", p2Address, shares.Share3, &wg)

	// Wait for shares from patient 1 and patient 2
	shareFromP1 := <-sharesChannel
	shareFromP2 := <-sharesChannel
	//fmt.Println(name, "has received both shares:", shareFromP1, shareFromP2)

	// Local computation: sum the shares with patients own share
	ownShare := shares.Share1
	localSum := (ownShare + shareFromP1 + shareFromP2) % (R+1)
	fmt.Println(name, "local sum is:", localSum)

	// Send the result to the hospital
	hospitalAddress := "localhost:8444" // Hospital's address and port
	sendResultToHospital("cert.pem", hospitalAddress, localSum)

	// Wait for the listener goroutine to finish
	wg.Wait()

	return nil
}
