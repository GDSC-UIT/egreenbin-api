package main

import (
	"context"
	"fmt"
	"log"

	"github.com/GDSC-UIT/egreenbin-api/component"
	middleware "github.com/GDSC-UIT/egreenbin-api/middlewares"
	"github.com/GDSC-UIT/egreenbin-api/modules/person/delivery"
	"github.com/GDSC-UIT/egreenbin-api/modules/person/repositories"
	"github.com/GDSC-UIT/egreenbin-api/modules/person/usecases"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// const uri = "mongodb://admin:123123123@localhost:27017/?maxPoolSize=20&w=majority"
const uri = "mongodb+srv://admin:123123123@gdsc.uytfb9v.mongodb.net/?retryWrites=true&w=majority"

func main() {
	// Set up the MongoDB connection
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")

	router := gin.Default()
	db := client.Database("egreenbin")
	appContext := component.NewAppContext(db)

	router.Use(middleware.Recover(appContext))
	personRepo := repositories.NewPersonRepository(db)
	personUseCase := usecases.NewPersonUsecase(personRepo)

	apiR := router.Group("/api")
	delivery.NewPersonHandler(apiR, appContext, personUseCase)

	log.Fatal(router.Run())
}
