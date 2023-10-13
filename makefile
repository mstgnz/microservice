include .env

.PHONY: up up_build down clean build
.DEFAULT_GOAL:= up

up:
	docker-compose up -d

up_build:
	docker-compose down
	docker-compose up --build -d

down:
	docker-compose down

clean:
	docker rmi $(docker images -f "dangling=true" -q)

build: build_auth build_blog-api build_consul-register build_listener build_logger build_mail build_sms

build_auth:
	cd ../auth && docker build -t $(PROJECT_NAME)-auth .

build_blog-api:
	cd ../blog-api && docker build -t $(PROJECT_NAME)-blog-api .

build_consul-register:
	cd ../consul && docker build -t $(PROJECT_NAME)-consul-register .

build_listener:
	cd ../listener && docker build -t $(PROJECT_NAME)-listener .

build_logger:
	cd ../logger && docker build -t $(PROJECT_NAME)-logger .

build_mail:
	cd ../mail && docker build -t $(PROJECT_NAME)-mail .

build_sms:
	cd ../sms && docker build -t $(PROJECT_NAME)-sms .