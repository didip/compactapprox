package main

import (
	"fmt"

	"github.com/didip/compactapprox/approx"
)

func main() {
	approxTracker, _ := approx.NewCompactApproximator(2048, 5, approx.SHA256)

	keys := [][]byte{
		[]byte("temperature:room1"),
		[]byte("temperature:room2"),
		[]byte("humidity:room1"),
	}

	values := []uint64{22, 24, 60}

	for i := range keys {
		approxTracker.Insert(keys[i], values[i])
	}

	for _, key := range keys {
		value := approxTracker.Get(key)
		fmt.Printf("Approximate value for %s: %d\n", key, value)
	}
}
