package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
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
var GetTransportByTypeName = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	var transports []*model.Transport
	var transportGroup *model.TransportGroup
	params := mux.Vars(request)
	typeName, _ := params["name"]
	fmt.Println(typeName)
	typeGroupCollection := client.Database("admin").Collection("TransportGroup")
	transportCollection := client.Database("admin").Collection("Transport")
	filter := bson.D{}

	// find some transport groups from request body by name
	if typeName != "" {
		filter = bson.D{{"name", typeName}}
		typeGroupCollection.FindOne(context.TODO(), filter).Decode(&transportGroup)

		for _, unitId := range transportGroup.UnitsIds {
			var transport model.Transport
			transportCollection.FindOne(context.TODO(), bson.M{"_id": unitId}).Decode(&transport)
			transports = append(transports, &transport)
		}
	}

	middlewares.SuccessArrRespond(transports, response)
})

var UpdateTransportsEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
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

	addTransportIdInGroup(&transport)

	collection := client.Database("admin").Collection("Transport")
	filter := bson.D{{"_id", transport.ID}}
	result, err := collection.ReplaceOne(context.TODO(), filter, transport)

	if err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}

	res, _ := json.Marshal(result.UpsertedID)
	middlewares.SuccessResponse(`Updated at `+strings.Replace(string(res), `"`, ``, 2), response)
})

var GetTransportTypesEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

	collection := client.Database("admin").Collection("Transport")
	var transportTypes []*model.TransportGroup
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}
	for cursor.Next(context.TODO()) {
		var transportType model.TransportGroup
		err := cursor.Decode(&transportType)
		if err != nil {
			log.Fatal(err)
		}

		transportTypes = append(transportTypes, &transportType)
	}
	if err := cursor.Err(); err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}
})

var GetTransportsEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	var transports []*model.Transport

	collection := client.Database("admin").Collection("Transport")
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
	fmt.Println(transports)
	middlewares.SuccessArrRespond(transports, response)
})

var CreateTransportEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	var transport model.Transport
	transport.ID = primitive.NewObjectID()

	err := json.NewDecoder(request.Body).Decode(&transport)
	if err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}

	if ok, errors := validators.ValidateInputs(transport); !ok {
		middlewares.ValidationResponse(errors, response)
		return
	}

	addTransportIdInGroup(&transport)

	collection := client.Database("admin").Collection("Transport")
	result, err := collection.InsertOne(context.TODO(), transport)

	if err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}

	res, _ := json.Marshal(result.InsertedID)
	middlewares.SuccessResponse(`Inserted at `+strings.Replace(string(res), `"`, ``, 2), response)
})

func addTransportIdInGroup(transport *model.Transport) {
	// filter for Transport group
	var transportGroup model.TransportGroup
	filter := bson.D{{"name", transport.Type}}
	//find existing transport group
	client.Database("admin").Collection("TransportGroup").FindOne(context.TODO(), filter).Decode(&transportGroup)
	if transportGroup.Name != "" {
		transportGroup.UnitsIds = append(transportGroup.UnitsIds, transport.ID)
		client.Database("admin").Collection("TransportGroup").ReplaceOne(context.TODO(), filter, transportGroup)
	}
}

// DeletePersonEndpoint  delete transport by id
var DeleteTransportEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var transport model.Transport

	collection := client.Database("admin").Collection("Transport")
	err := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&transport)
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
