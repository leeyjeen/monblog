APP_NAME = monblog_api

test:
	@go test -v -cover

build:
	@go build -o ${APP_NAME}