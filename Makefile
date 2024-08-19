.PHONY: compose-build compose-up compose-down

compose-build:
	docker-compose build

compose-up:
	docker-compose up

compose-down:
	docker-compose down
