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
	go test ./test/... -p=1 -count=1

feature_test:
	go test ./test/feature/... -p=1 -count=1

unit_test:
	go test ./test/unit/... -count=1

gotest_cover:
	mkdir -p coverage
	go test -coverpkg=./internal/... -coverprofile=coverage/coverage.out ./test/...
	go tool cover -html=coverage/coverage.out -o ./coverage/coverage.html
