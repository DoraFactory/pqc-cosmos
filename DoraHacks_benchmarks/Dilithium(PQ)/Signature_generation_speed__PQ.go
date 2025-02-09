package main

import (
	"fmt"
	"time"

	"github.com/tendermint/tendermint/crypto/dilithium"
)

func main() {
	// Number of iterations to benchmark
	const iterations = 1000

	// Generate a new Dilithium private key
	privKey := dilithium.GenPrivKey()

	// Create a dummy message to sign
	message := []byte("Benchmarking signature generation speed")

	// Start the timer
	start := time.Now()

	// Sign the message multiple times
	for i := 0; i < iterations; i++ {
		_, err := privKey.Sign(message)
		if err != nil {
			fmt.Println("Error signing message:", err)
			return
		}
	}

	// Calculate the total time taken
	duration := time.Since(start)

	// Calculate operations per second (ops/sec)
	opsPerSecond := float64(iterations) / duration.Seconds()

	// Print the results
	fmt.Printf("Generated %d signatures in %v\n", iterations, duration)
	fmt.Printf("Signature generation speed: %.2f ops/second\n", opsPerSecond)
}
