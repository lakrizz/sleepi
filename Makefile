.PHONY: dc-up dc-down run

dc-up:
	docker-compose up -d 

dc-down:
	docker-compose down -d 

run:
	go run cmd/sleepi/main.go

prepare-mpd:
	mkdir -p ./run/data/
	mkdir -p ./run/playlists/
	mkdir -p ./run/music/
	chmod 777 ./run/data/
	chmod 777 ./run/playlists/
	chmod 777 ./run/music/
