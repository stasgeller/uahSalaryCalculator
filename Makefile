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