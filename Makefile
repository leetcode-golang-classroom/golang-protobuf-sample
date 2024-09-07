.PHONY=build

build:
	@CGO_ENABLED=0 GOOS=linux go build -o bin/main cmd/main.go

run: build
	@./bin/main

coverage:
	@go test -v -cover ./tests/...

test:
	@go test -v ./tests/...

gen-marshall-binary:
	@go run helps/marshall/person_marshall.go

deserialize-marshall-binary:
	@go run helps/unmarshall/person_unmarshall.go

send-marshall-data: gen-marshall-binary
	@curl -X POST --data-binary @tmp/person.bin http://localhost:8080/person

receive-binary-data:
	@curl -X GET "http://localhost:8080/person?id=1234" --output tmp/retrieved_person.bin

gen-proto-client:
	@protoc --go_out=. --go_opt=paths=source_relative protos/person.proto