package controller

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"transport-manager/m/v1/app/db"
	middlewares "transport-manager/m/v1/app/handler"
	"transport-manager/m/v1/app/model"
)
var client = db.Dbconnect()

var GetTransportEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	var transports []*model.Transport

	collection := client.Database("admin").Collection("Transport")
	fmt.Println(collection)
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}
	for cursor.Next(context.TODO()) {
		var transport model.Transport
		err := cursor.Decode(&transport)
		if err != nil {
			log.Fatal(err)
		}

		transports = append(transports, &transport)
	}
	if err := cursor.Err(); err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}
	middlewares.SuccessArrRespond(transports, response)
})