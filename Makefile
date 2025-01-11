weather: weather.go
	go build -o weather

weather.wasm: weather.go
	GOOS=js GOARCH=wasm go build -o weather.wasm
	cp `go env GOROOT`/misc/wasm/wasm_exec.js .

clean:
	rm -f weather weather.wasm wasm_exec.js
