package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gorilla/mux"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/new-request", func(w http.ResponseWriter, r *http.Request) {
		clientIp := r.RemoteAddr
		log.Printf("client ip is: %s \n", clientIp)

		ip := IP{
			Address: clientIp,
			Time:    time.Now(),
		}

		collection := client.Database("mongo").Collection("ips")
		_, err := collection.InsertOne(context.Background(), ip)
		if err != nil {
			http.Error(w, "error storing ip: "+clientIp, http.StatusInternalServerError)
			return
		}

		log.Printf("the client ip  %s stored in the database succesfully\n", clientIp)
	})

	// Start the HTTP server on port 8080
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
