package main

import (
	"fmt"

	"github.com/didip/compactapprox/approx"
)

func main() {
	approx, err := approx.NewCompactApproximator(1024, 4, approx.FNV)
	if err != nil {
		panic(err)
	}

	key := []byte("user:567")
	approx.Insert(key, 123)
	approx.Insert(key, 110)
	value := approx.Get(key)
	fmt.Printf("Approximate value for key %s: %d\n", key, value)
}
