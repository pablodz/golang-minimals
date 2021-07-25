package main

import (
	"context"
	"io"
	"log"

	pb "../datafiles"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func ReceiveStream(client pb.MoneyTransactionClient, request *pb.TransactionRequest) {
	log.Println("Started listening to the server stream!")
	stream, err := client.MakeTransaction(context.Background(), request)
	if err != nil {
		log.Fatalf("%v.MakeTransaction(_) = _, %v", client, err)
	}
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.MakeTransaction(_) = _, %v", client, err)
		}
		log.Printf("Status: %v, Operation: %v", response.Status, response.Description)
	}
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewMoneyTransactionClient(conn)
	from := "1234"
	to := "9876"
	amount := float32(3214.56)

	ReceiveStream(client, &pb.TransactionRequest{
		From:   from,
		To:     to,
		Amount: amount})
}
