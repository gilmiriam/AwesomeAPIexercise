.PHONY: build up down run_test

build:
	docker-compose build
up:
	docker-compose up -d
down:
	docker-compose down -v
run_test:
	docker-compose exec app bash -lc 'go test -v ./...'