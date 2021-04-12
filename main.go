package main

import (
	"GolangProject/controllers"
	"GolangProject/repositories"
	"GolangProject/services"
	"context"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Database   = "goMongo"
	Collection = "accounts"
	MongoDbUrl = "mongodb://mongodb:27017/"
)

var collection *mongo.Collection
var ctx = context.TODO()

func main() {
	log.Info("Hi, Welcome !")

	initDatabase()

	repo := repositories.NewAccountRepositories(collection, ctx)
	service := services.NewService(repo)
	controller := controllers.NewAccountsController(service)

	r := gin.Default()

	registerHandlers(r, controller)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}

func initDatabase() {

	log.Info("Connecting to datastore")
	clientOptions := options.Client().ApplyURI(MongoDbUrl)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database(Database).Collection(Collection)
}

func registerHandlers(r *gin.Engine, controller controllers.AccountController) {
	r.GET("/account", controller.GetAccounts())
	r.GET("/account/:id", controller.GetAccount())
	r.POST("/account", controller.CreateAccount())
	r.PATCH("/account/:id", controller.UpdateAccount())
	r.DELETE("/account/:id", controller.DeleteAccount())
}
