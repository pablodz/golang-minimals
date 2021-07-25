package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

func main() {
	p := pb.Person{
		Id:    1234,
		Name:  "Albert",
		Email: "albert@gmail.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	p1 := &pb.Person{}
	body, _ := proto.Marshal(p)
	_ = proto.Unmarshal(body, p1)
	fmt.Println("Original structu loaded from proto file: ", p, "\n")
	fmt.Println("Marshaled proto data: ", body, "\n")
	fmt.Println("Unmarshañled struct: ", p1)
}
