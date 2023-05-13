export

build: ### run movements_producer in local environment
	go build -o morty_grab.exe ./cmd/main.go
	./morty_grab.exe
.PHONY: morty_grab.exe