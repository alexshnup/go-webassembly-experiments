package main

//go:wasm-module yourmodulename
//export add
func add(x, y uint32) uint32 {
	return x + y
}

//go:wasm-module yourmodulename
//export multiply
func multiply(x, y uint32) uint32 {
	return x * y
}

// main is required for the `wasi` target, even if it isn't used.
func main() {}
