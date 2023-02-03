GOOS=linux
GOARCH=amd64

build_all:
	rm -rf ./dist
	$(MAKE) build_api
	$(MAKE) build_cli

_build_base:
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -a -tags=netgo,$(ENV) -installsuffix netgo \
		-o $(OUTPUT_FILE) ./cmd/$(CMD)/.

build_api:
	CMD=api OUTPUT_FILE=dist/api ENV=production $(MAKE) _build_base

build_cli:
	CMD=cli OUTPUT_FILE=dist/cli $(MAKE) _build_base

lint:
	golangci-lint run ./... --timeout=2m

gotest:
	$(MAKE) feature_test
	$(MAKE) unit_test

feature_test:
	go test ./test/feature/... -tags=feature_test -p=5 -count=1 ${ARG} -run=${RUN}

unit_test:
	go test ./test/unit/... -count=1 ${ARG} -run=${RUN}

gotest_cover:
	mkdir -p coverage
	go test -coverpkg=./internal/... -coverprofile=coverage/coverage.out ./test/...
	go tool cover -html=coverage/coverage.out -o ./coverage/coverage.html

migrate:
	migrate \
		-database "mysql://$(DB_USER_NAME):$(DB_PASSWORD)@tcp($(DB_HOST_NAME):3306)/$(MYSQL_DATABASE)?parseTime=true&loc=Asia%2FTokyo" \
		-path ./build/migrate \
		up
migrate_down:
	migrate \
		-database "mysql://$(DB_USER_NAME):$(DB_PASSWORD)@tcp($(DB_HOST_NAME):3306)/$(MYSQL_DATABASE)?parseTime=true&loc=Asia%2FTokyo" \
		-path ./build/migrate \
		down
