dev1:
	pnpm run dev

dev2:
	go run main.go

build:
	go build

migrate:
	goose up
