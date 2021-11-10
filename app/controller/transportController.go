package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/umangraval/Go-Mongodb-REST-boilerplate/models"
	"github.com/umangraval/Go-Mongodb-REST-boilerplate/validators"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"strings"
	"transport-manager/m/v1/app/db"
	middlewares "transport-manager/m/v1/app/handler"
	"transport-manager/m/v1/app/model"
)

var client = db.Dbconnect()

var GetTransportsEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
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

var CreateTransportEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	var transport model.Transport
	err := json.NewDecoder(request.Body).Decode(&transport)
	if err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}
	if ok, errors := validators.ValidateInputs(transport); !ok {
		middlewares.ValidationResponse(errors, response)
		return
	}
	collection := client.Database("admin").Collection("Transport")
	result, err := collection.InsertOne(context.TODO(), transport)
	if err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}
	res, _ := json.Marshal(result.InsertedID)
	middlewares.SuccessResponse(`Inserted at `+strings.Replace(string(res), `"`, ``, 2), response)
})

// DeletePersonEndpoint -> delete transport by id
var DeletePersonEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var person models.Person

	collection := client.Database("admin").Collection("Transport")
	err := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&person)
	if err != nil {
		middlewares.ErrorResponse("Transport does not exist", response)
		return
	}
	_, derr := collection.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}})
	if derr != nil {
		middlewares.ServerErrResponse(derr.Error(), response)
		return
	}
	middlewares.SuccessResponse("Deleted", response)
})

// GetTransportEndpoint -> get transport by id
var GetTransportEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var transport model.Transport

	collection := client.Database("admin").Collection("Transport")
	err := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&transport)
	if err != nil {
		middlewares.ErrorResponse("Transport does not exist", response)
		return
	}
	middlewares.SuccessRespond(transport, response)
})
