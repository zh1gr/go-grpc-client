package main

import (
	"context"
	"github.com/zh1gr/go-grpc-domain/invoice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("could not close connection: %v", err)
		}
	}(conn)

	c := invoice.NewInvoiceClient(conn)

	message := invoice.CreateRequest{
		Amount: &invoice.Amount{
			Amount:   425,
			Currency: "USDT",
		},
		From: "aliquip eiusmod officia in in",
		To:   "amet",
	}

	res, err := c.Create(context.Background(), &message)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
		return
	}

	log.Fatalf("response was: %v", res)
}
