package main

import (
	"log"
	"os"

	"github.com/leetcode-golang-classroom/golang-protobuf-sample/protos"
	"google.golang.org/protobuf/proto"
)

func main() {
	person := &protos.Person{
		Name:  "John Wick",
		Id:    1234,
		Email: "wick@codetest.com",
		Phones: []*protos.PhoneNumber{
			{
				Number: "1234-111-2222",
				Type:   protos.PhoneType_MOBILE,
			},
		},
	}
	data, err := proto.Marshal(person)
	if err != nil {
		log.Fatalf("Failed to marshal: %v", err)
	}
	os.WriteFile("tmp/person.bin", data, 0644)
}
