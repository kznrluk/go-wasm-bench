# go-wasm-bench

TODO: JSで同じもの書いて試す

```go
func main() {
	start := time.Now().UnixNano()
	for i := 1; i < 1000000; i++ {
		_ = collatz(i)
	}
	stop := time.Now().UnixNano()
	println(stop - start)
	// M1 Mac native go 189659000
	// Firefox wasm     1665000000
	// Chrome wasm      783000000
}
```
