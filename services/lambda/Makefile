clean:
	@rm -rf bin
	@mkdir -p bin

BUILD_DIRS = users

build: clean
	for fun in `ls cmd`; do \
		GOOS=linux go build -o bin/$$fun ./cmd/$$fun; \
		$(GOPATH)/bin/build-lambda-zip -o ../../tmp/$$fun.zip bin/$$fun; \
	done
