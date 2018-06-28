cup:
	docker-compose -f docker-compose.core.yml up -d --build

cdown:
	docker-compose -f docker-compose.core.yml down

sup:
	docker-compose -f docker-compose.services.yml up -d --build

sdown:
	docker-compose -f docker-compose.services.yml down

init:
	docker network create oddzy