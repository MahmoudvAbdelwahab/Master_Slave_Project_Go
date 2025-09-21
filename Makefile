.PHONY: all build clean run-example

all: build

build:
	mkdir -p bin
	go build -o bin/master ./master
	go build -o bin/slave ./slave

clean:
	rm -rf bin

run-example:
	./run_example.sh
