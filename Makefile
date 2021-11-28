# loading env if file exists
ifeq ($(shell test -s .env && echo -n yes),yes)
	include .env
	export $(shell sed 's/=.*//' .env)
endif


default: build

build:
	go build ./...

test:
	go test ./...

swagger_gen:
	swag init

docker_build:
	docker build -t go-rest-fizzbuzz .

docker_run:
	docker run -it -p 5000:5000 go-rest-fizzbuzz

