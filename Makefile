# Use once when install app
init:
	@echo "Install application"
	@cp ./.env.example ./.env
	command make up

up:
	@docker-compose build
	@docker-compose up

down:
	@docker-compose down

migrate:
	@migrate -path=./db/migrations/ -database sqlite3://db/salary_bot up

create-migration:
	@migrate create -ext sql -dir ./db/migrations -seq `date +"%s"` ' * 1000000)/1'