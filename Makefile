.PHONY: build deploy web

# Create/Destroy core images
cup:
	OD_DEPLOY=local docker-compose -f ./build/docker-compose.core.yml up -d --build

cdown:
	OD_DEPLOY=local docker-compose -f ./build/docker-compose.core.yml down



# Create/Destroy app images
up:
	OD_DEPLOY=local docker-compose -f ./build/docker-compose.app.yml up -d --build

down:
	OD_DEPLOY=local docker-compose -f ./build/docker-compose.app.yml down



# Create/Destroy web images
wup:
	OD_DEPLOY=local docker-compose -f ./build/docker-compose.web.yml up -d --build

wdown:
	OD_DEPLOY=local docker-compose -f ./build/docker-compose.web.yml down


# build and deploy all images
deploy: build
	@eval $$(aws ecr get-login --no-include-email --region ap-southeast-2)
	OD_DEPLOY=remote docker-compose -f ./build/docker-compose.app.yml -f ./build/docker-compose.core.yml -f ./build/docker-compose.web.yml push

# build all images
build:
	OD_DEPLOY=remote docker-compose -f ./build/docker-compose.core.yml -f ./build/docker-compose.app.yml -f ./build/docker-compose.web.yml build

# Must be run the first time doing local dev on a machine.
init:
	docker network create oddzy

# Updates remote deployment
apply:
	cd ./deploy/terraform; terraform apply -auto-approve

# Determine changes to remote deployment
plan:
	cd ./deploy/terraform; terraform plan

# Destroy remote deployment
destroy:
	cd ./deploy/terraform; terraform destroy -auto-approve

