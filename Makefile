.PHONY: build deploy web

# Create/Destroy core images
core-up:
	OD_DEPLOY=local docker-compose -f ./build/docker-compose.core.yml up -d --build

core-down:
	OD_DEPLOY=local docker-compose -f ./build/docker-compose.core.yml down



# Create/Destroy app images
app-up:
	OD_DEPLOY=local docker-compose -f ./build/docker-compose.app.yml up -d --build

app-down:
	OD_DEPLOY=local docker-compose -f ./build/docker-compose.app.yml down



# Create/Destroy web images
web-up:
	OD_DEPLOY=local docker-compose -f ./build/docker-compose.web.yml up -d --build

web-down:
	OD_DEPLOY=local docker-compose -f ./build/docker-compose.web.yml down


# build and deploy all images
deploy: build
	@eval $$(aws ecr get-login --no-include-email --region ap-southeast-2)
	OD_DEPLOY=remote docker-compose -f ./build/docker-compose.app.yml -f ./build/docker-compose.core.yml -f ./build/docker-compose.web.yml push

# build all images
build:
	OD_DEPLOY=remote docker-compose -f ./build/docker-compose.core.yml -f ./build/docker-compose.app.yml -f ./build/docker-compose.web.yml build

# Updates remote deployment
apply:
	cd ./deploy/terraform; terraform apply -auto-approve

# Destroy remote deployment
destroy:
	cd ./deploy/terraform; terraform destroy -auto-approve

# Build the website, upload it to s3 and trigger the remote update script
web-update:
	cd ./web; npm run build
	aws s3 sync ./web/dist s3://oddzy/web/dist --delete
	aws ssm send-command --document-name oddzy-test-update-website --targets Key=tag:name,Values=oddzy

make web-serve:
	cd ./web; npm run serve