package main

import (
	"fmt"
	"syscall/js" //Interacting with JavaScript using syscall/js
)

func main() {
	fmt.Println("Hello, WebAssembly from Go!")
	js.Global().Get("console").Call("log", "Hello from Go!") // Interacting with JavaScript using syscall/js
}
