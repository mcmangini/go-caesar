BINARY_NAME = caesar

all: build

build:
	go build -o ${BINARY_NAME} caesar.go

clean:
	go clean

install:
	mkdir -p /usr/local/bin
	install -m 0755 ${BINARY_NAME} /usr/local/bin

uninstall:
	rm /usr/local/bin/${BINARY_NAME}
