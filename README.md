## My experiments with WebAssembly

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