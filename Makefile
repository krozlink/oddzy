cup:
	docker-compose -f docker-compose.core.yml up -d --build

cdown:
	docker-compose -f docker-compose.core.yml down

up:
	docker-compose -f docker-compose.services.yml up -d --build

down:
	docker-compose -f docker-compose.services.yml down

push: build
	@eval $$(aws ecr get-login --no-include-email --region ap-southeast-2)
	docker-compose -f docker-compose.services.yml push

build:
	docker-compose -f docker-compose.services.yml build

init:
	docker network create oddzy