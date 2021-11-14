package db

import (
	"context"
	"github.com/OnlyM1ss/transport-manager/v2/app/config"
	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

// Dbconnect -> connects mongo
func Dbconnect() *mongo.Client {
	conf := config.New()

	clientOptions := options.Client().ApplyURI(conf.MongoDb.Url)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("⛒ Connection Failed to Database")
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("⛒ Connection Failed to Database")
		log.Fatal(err)
	}
	color.Green("⛁ Connected to Database")
	return client
}
