default: clean run

build:
	cd lib ; go build -buildmode=plugin

run: build
	go run main.go

clean:
	go clean
	find . -name "*.so" | xargs -I{} rm -f {}