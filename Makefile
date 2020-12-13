SRC := $(shell find . -type f -name "*.go")

radar: $(SRC)
	go build -o radar cmd/radar/main.go
