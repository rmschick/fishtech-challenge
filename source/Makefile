GOVERSION = $(shell go version | awk '{print $$3;}')
SOURCE_FILES?=./...

export PATH := ./bin:$(PATH)
export GOPRIVATE := github.com/rmschick99
export CGO_ENABLED := 0

clean:
	rm -rf ./dist && rm -rf ./vendor
.PHONY: clean

upgrade:
	go get -t -u ./...
.PHONY: upgrade

vendor:
	go mod vendor
.PHONY: vendor

tidy:
	go mod tidy
.PHONY: tidy