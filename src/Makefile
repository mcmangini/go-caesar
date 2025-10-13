BINARY_NAME = caesar

all: build test

build:
	go build -o ${BINARY_NAME} caesar.go

test:
	go test -v caesar.go

run:
	go build -o ${BINARY_NAME} caesar.go
	./${BINARY_NAME}

clean:
	go clean

install:
	go build -o ${BINARY_NAME} caesar.go
	chmod +x ${BINARY_NAME}
	mv ${BINARY_NAME} /usr/local/bin

uninstall:
	rm /usr/local/bin/${BINARY_NAME}
