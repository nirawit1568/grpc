package handlers

import (
	"client/services"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Hello(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")

	// connect grpc
	creds := insecure.NewCredentials()
	cc, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	// init client
	calculatorClient := services.NewCalculatorClient(cc)
	req := services.HelloRequest{
		Name: name,
	}

	// call grpc
	resp, err := calculatorClient.Hello(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}

	// response api
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Result)
}
