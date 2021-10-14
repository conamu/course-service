run:
	@echo "Building and running test instance"
	docker build --tag "course-service:local" .
	docker-compose -f docker-compose-dev.yml up

run-prod:
	docker-compose -f docker-compose-prod.yml up
