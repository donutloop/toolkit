.PHONY: all
all: build test

ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")

setup:
	go get -u github.com/golang/lint/golint

build: fmt vet lint

fmt:
	go fmt $(ALL_PACKAGES)

vet:
	go vet $(ALL_PACKAGES)

goimports:
	goimports -w $(shell find . -type f -name '*.go' -not -path "./vendor/*")

gofmts:
	gofmt -s -w  $(shell find . -type f -name '*.go' -not -path "./vendor/*")

lint:
	@for p in $(ALL_PACKAGES); do \
		echo "==> Linting $$p"; \
		golint -set_exit_status $$p; \
	done

test: fmt vet build
	ENVIRONMENT=test go test $(ALL_PACKAGES)

test-cover-html:
	@echo "mode: count" > coverage-all.out

	$(foreach pkg, $(ALL_PACKAGES),\
	ENVIRONMENT=test go test -coverprofile=coverage.out -covermode=count $(pkg);\
	tail -n +2 coverage.out >> coverage-all.out;)
	go tool cover -html=coverage-all.out -o out/coverage.html