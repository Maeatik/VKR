export

build: ### run movements_producer in local environment
	go build -o morty_grab ./cmd/main.go
	./morty_grab
.PHONY: morty_grab