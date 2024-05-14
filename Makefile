dev1:
	pnpm run dev

dev2:
	go run main.go

builds:
	tinygo build -o ./build/app.wasm -target wasm -no-debug ./main.go
