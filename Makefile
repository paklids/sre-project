all: run
test:
	@docker-compose -f docker-compose.test.yaml down && docker-compose -f docker-compose.test.yaml build && docker-compose -f docker-compose.test.yaml up --abort-on-container-exit
run:
	@docker-compose down && docker-compose build --no-cache && docker-compose up