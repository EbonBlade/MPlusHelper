all: build-tools lint test vet fmt build

test:
	go test -v -cover ./...

vet:
	go vet ./...

fmt:
	go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l
	test -z $$(go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l)

lint:
	go list ./... | grep -v /vendor/ | xargs -L1 golint -set_exit_status

build-tools:
	go install golang.org/x/lint/golint

build:
	go build -o bin/service ./cmd/service

run:
	./bin/service
