default: start

start:
	ENV_FILE=.env go run .

build:
	go build . -o target/main

serve:
	./target/main
