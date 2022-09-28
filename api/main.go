package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	pb "github.com/whutchinson98/hutchery-k8/api/pb/inventory"

	"google.golang.org/grpc"
)

func main() {
	// If running locally you'll want to use localhost:5000 instead
	conn, err := grpc.Dial("books-service:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	client := pb.NewInventoryClient(conn)
	routes := mux.NewRouter()
	routes.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)
		defer cancel()
		bookList, err := client.GetBookList(ctx, &pb.GetBookListRequest{})

		if err != nil {
			log.Fatalf("failed to get book list: %v", err)
			json.NewEncoder(w).Encode(fmt.Sprintf("Error occured %v", err.Error()))
		} else {
			json.NewEncoder(w).Encode(bookList)
		}
	}).Methods("GET")
	fmt.Println("Application listening on :8080")
	http.ListenAndServe(":8080", routes)
}
