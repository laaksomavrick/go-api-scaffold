run:
	GO_ENV=development DB_NAME=goals_development go run src/main.go

install:
	dep ensure

test:
	GO_ENV=testing DB_NAME=goals_testing go test -v ./...

up:
	docker-compose up

down:
	docker-compose down

psql:
	psql -U postgres -h 127.0.0.1 -p 5432

migrate:
	ruby db/migrater.rb