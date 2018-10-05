clean:
	@rm -rf dist
	@mkdir -p dist

BUILD_DIRS = users

build:
	@for ent in $(BUILD_DIRS); do \
		for fun in `ls src/$$ent`; do \
			GOOS=linux go build -o bin/$$ent-$$fun ./src/$$ent/$$fun; \
			$(GOPATH)/bin/build-lambda-zip -o ../../tmp/$$ent-$$fun.zip bin/$$ent-$$fun; \
		done \
	done
