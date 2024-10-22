package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// range of finite field
	R := 100

	// Start Alice
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := startPatient("Alice",R,42,"cert.pem","key.pem",":8441",":8442",":8443"); err != nil {
			fmt.Println("Error starting Alice:", err)
		}
	}()

	// Start Bob
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := startPatient("Bob",R,37,"cert.pem","key.pem",":8442",":8441",":8443"); err != nil {
			fmt.Println("Error starting Bob:", err)
		}
	}()

	// Start Charlie
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := startPatient("Charlie",R,77,"cert.pem","key.pem",":8443",":8441",":8442"); err != nil {
			fmt.Println("Error starting Charlie:", err)
		}
	}()

	// Start Hospital
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := startHospital(R); err != nil {
			fmt.Println("Error starting Hospital:", err)
		}
	}()

	// Wait for all to finish
	wg.Wait()
}
