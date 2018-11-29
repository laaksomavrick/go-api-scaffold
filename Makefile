run:
	go run src/main.go

install:
	dep ensure

test:
	go test -v ./...

up:
	docker-compose up

down:
	docker-compose down

psql:
	psql -U postgres -h 127.0.0.1 -p 5432