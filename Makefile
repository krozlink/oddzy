.PHONY: build deploy

cup:
	OD_DEPLOY=local docker-compose -f ./build/docker-compose.core.yml up -d --build

cdown:
	OD_DEPLOY=local docker-compose -f ./build/docker-compose.core.yml down

up:
	OD_DEPLOY=local docker-compose -f ./build/docker-compose.app.yml up -d --build

down:
	OD_DEPLOY=local docker-compose -f ./build/docker-compose.app.yml down

deploy: build
	@eval $$(aws ecr get-login --no-include-email --region ap-southeast-2)
	OD_DEPLOY=remote docker-compose -f ./build/docker-compose.app.yml -f ./build/docker-compose.core.yml push

build:
	OD_DEPLOY=remote docker-compose -f ./build/docker-compose.core.yml -f ./build/docker-compose.app.yml build

init:
	docker network create oddzy

apply:
	cd ./deploy/terraform; terraform apply -auto-approve

plan:
	cd ./deploy/terraform; terraform plan

destroy:
	cd ./deploy/terraform; terraform destroy -auto-approve