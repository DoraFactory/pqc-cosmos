package main

import (
	"fmt"
	"time"

	"github.com/tendermint/tendermint/crypto/ed25519"
)

func main() {

	const iterations = 1000
	var totalDuration time.Duration

	for i := 0; i < iterations; i++ {
		start := time.Now() // start time
		ed25519.GenPrivKey()

		elapsed := time.Since(start) // start time + generation time
		totalDuration += elapsed
	}

	averageDuration := totalDuration / iterations // getting the avg execution time
	fmt.Printf("Average duration over %d iterations: %v\n", iterations, averageDuration)
}
