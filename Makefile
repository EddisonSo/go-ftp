all: server client

server:
	mkdir -p build
	go build  -o ./build ./cmd/client/

client:
	mkdir -p build
	go build  -o ./build ./cmd/server/

clean:
	rm -rf build