// GOOS=js GOARCH=wasm tinygo build -o main.wasm

package main

import (
	"time"
)

func collatz(n int) int {
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
	}
	return 1
}

func main() {
	start := time.Now().UnixNano()
	for i := 1; i < 1000000; i++ {
		_ = collatz(i)
	}
	stop := time.Now().UnixNano()
	println(stop - start)
	// M1 Mac native go 189659000
	// Firefox wasm     1665000000
}
