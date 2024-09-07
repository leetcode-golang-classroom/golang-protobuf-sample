package main

import (
	"fmt"
	"log"
	"os"

	"github.com/leetcode-golang-classroom/golang-protobuf-sample/protos"
	"google.golang.org/protobuf/proto"
)

func main() {
	data, err := os.ReadFile("tmp/retrieved_person.bin")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	person := &protos.Person{}
	err = proto.Unmarshal(data, person)
	if err != nil {
		log.Fatalf("Failed to unmarshall:%v", err)
	}
	fmt.Println("Deserialized Person:", person)
}
