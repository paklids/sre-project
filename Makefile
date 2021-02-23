all: test run
test:
	@docker-compose -f docker-compose.test.yaml build --no-cache && \
	docker-compose -f docker-compose.test.yaml up --abort-on-container-exit && \
	docker-compose -f docker-compose.test.yaml down
run:
	@docker-compose down && \
	docker-compose build && \
	docker-compose up
clean-dock:
	@docker rmi `docker images -f "dangling=true" -q --no-trunc`