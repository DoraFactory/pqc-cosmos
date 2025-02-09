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

	// Extract the corresponding public key
	pubKey := privKey.PubKey()

	// Create a dummy message to sign
	message := []byte("Benchmarking signature verification speed")

	// Generate a signature for the message
	signature, err := privKey.Sign(message)
	if err != nil {
		fmt.Println("Error generating signature:", err)
		return
	}

	// Start the timer
	start := time.Now()

	// Verify the signature multiple times
	for i := 0; i < iterations; i++ {
		if !pubKey.VerifySignature(message, signature) {
			fmt.Println("Signature verification failed!")
			return
		}
	}

	// Calculate the total time taken
	duration := time.Since(start)

	// Calculate operations per second (ops/sec)
	opsPerSecond := float64(iterations) / duration.Seconds()

	// Print the results
	fmt.Printf("Verified %d signatures in %v\n", iterations, duration)
	fmt.Printf("Signature verification speed: %.2f ops/second\n", opsPerSecond)
}
