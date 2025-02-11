package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/akashriva/gin_framework/helper"
)

type manger struct {
	Connection *mongo.Client
	Ctx        context.Context
	Cancel     context.CancelFunc
}

var Mgr manger

// init function to initialize MongoDB connection
func InitDdConnection() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("MONGO_URI environment variable is not set")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf(helper.DB_FAIL_CREATE_CLINT, err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf(helper.DB_CONN_FAIL, err)
		return
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
		return
	}

	Mgr = manger{Connection: client, Ctx: ctx, Cancel: cancel}
	fmt.Println(helper.DB_CONN_SUCCESSFULLY)
}

func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	//CancelFunc to cancel to context
	defer cancel()

	//client provide a methode to close
	defer func() {

		//client Disconnect method also has deadline
		//return error if any
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

}
