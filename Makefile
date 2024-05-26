local:
	cls
	go build
	./bobby

build:
	docker build -t bobby:latest .

run:
	docker run --name bobby -p 9033:9033 bobby:latest

.PHONY: local build run