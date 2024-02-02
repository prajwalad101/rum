build-local:
	go build -o ./build

install: build-local
	go install

