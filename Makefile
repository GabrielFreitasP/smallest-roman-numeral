.PHONY: install coverage

install:
	go get ./...
	go mod tidy
	go mod vendor

build:
	go build ./cmd/api/main.go

run:
	go run ./cmd/api/main.go

test:
	go test -v -cover ./...

lint:
	golangci-lint run ./...

coverage:
	mkdir -p ./coverage
	go test ./... -coverprofile ./coverage/cover.out.tmp
	cat ./coverage/cover.out.tmp > ./coverage/cover.out
	rm ./coverage/cover.out.tmp
	go tool cover -html=./coverage/cover.out
	go tool cover -html=./coverage/cover.out -o ./coverage/cover.html

clean:
	rm -rf ./vendor
	rm -rf ./bin
	rm -rf ./coverage
	rm -f ./go.sum

mocks:
	~/go/bin/mockgen \
	-source=internal/romanNumeral/usecase.go \
	-destination=internal/romanNumeral/mock/usecase.go \
	-package=mock