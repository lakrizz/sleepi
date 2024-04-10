.PHONY: dc-up dc-down run

dc-up:
	docker-compose up -d 

dc-down:
	docker-compose down -d 

run:
	go run cmd/sleepi/main.go
