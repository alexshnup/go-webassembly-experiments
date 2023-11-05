## My experiments with WebAssembly

We can build an application with Go (Golang) and compile it to WebAssembly (WASM) to run in a web browser. This approach leverages the client's CPU for calculations and data processing. Here's how it works:

- Go and WebAssembly: Go has support for compiling to WebAssembly. This means we can write our application logic in Go and then compile it into a .wasm file.

- Running in Browser: The compiled WebAssembly code can be executed in modern web browsers. Browsers have built-in WebAssembly runtimes, allowing our Go code to run at near-native speed.

- JavaScript Integration: Although WebAssembly runs in the browser, it typically needs to interact with the web page's JavaScript to access the DOM or perform other web-specific tasks. We can use JavaScript to load and run the WebAssembly module and establish communication between our Go code and the web environment.

- WASI (WebAssembly System Interface): WASI is a system interface for WebAssembly that aims to provide a standard set of APIs for WebAssembly modules, making it easier to run WebAssembly outside the web context, like on servers or in other computing environments. However, its usage in the context of a web browser is limited, as WASI is primarily designed for server-side or standalone applications.

- Efficiency and Performance: By using WebAssembly, we can offload complex computations to the client's CPU, potentially improving the performance and efficiency of our web application. This is especially useful for compute-intensive tasks like image processing, simulations, or cryptographic operations.

- Cross-Browser Support: WebAssembly is supported by all major browsers, ensuring broad compatibility.

- Security Considerations: Running code in WebAssembly is subject to the same security restrictions as JavaScript in the browser. It's sandboxed and cannot directly access the local file system or other sensitive resources on the client machine.


### Build server for local testing

```bash
go build -o srv _server/main.go 
```

### Build wasm with original go compiler

```bash
cd original
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js” .
GOOS=js GOARCH=wasm go build -o main.wasm main.go
ls -lh main.wasm
```
Size main.wasm 2.1M

### Build wasm with tinygo



```bash
cd tinygo
```

Please note the file wasm_exec.js must match the version tinygo was compiled with.
for latest version of tinygo it is located in tinygo/targets/wasm_exec.js
https://github.com/tinygo-org/tinygo/blob/release/targets/wasm_exec.js

```bash
tinygo build -o main.wasm -target=wasi -no-debug main.go
ls -lh main.wasm
81K
```
Size main.wasm 13K


```bash
tinygo build -target=wasi -gc=leaking -no-debug -o main.wasm main.go
ls -lh main.wasm
```
Size main.wasm 9K

-gc=leaking: This flag sets the garbage collector mode to "leaking". In this mode, the garbage collector doesn't actually reclaim memory, effectively "leaking" memory. This mode can improve performance and reduce binary size at the cost of increasing memory usage over time. It's typically used in scenarios where memory footprint is less of a concern or in short-lived programs.


## Summary
In summary, using Go to write code for WebAssembly is a viable approach for building web applications that require significant client-side computation, leveraging the client's CPU for better performance and efficiency.