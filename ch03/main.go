package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/liuning108/buildingDistributedAplicationsInGinForMe/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

var ctx context.Context
var err error
var clinet *mongo.Client
var collection *mongo.Collection
var recipesHandler *handlers.RecipesHandler

func init() {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	status := redisClient.Ping()
	fmt.Println(status)

	MONGO_URI := "mongodb://admin:ningwang@localhost:27017/test?authSource=admin"
	ctx = context.Background()
	clinet, err = mongo.Connect(ctx, options.Client().ApplyURI(MONGO_URI))
	if err = clinet.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connnected to MongoDB")

	MONGO_DATABASE := "test"
	collection = clinet.Database(MONGO_DATABASE).Collection("recipes")
	recipesHandler = handlers.NewRecipesHandler(collection, ctx, redisClient)
}

func main() {
	router := gin.Default()
	router.POST("/recipes", recipesHandler.NewRecipeHandler)
	router.GET("/recipes", recipesHandler.ListRecipesHandler)
	router.PUT("/recipes/:id", recipesHandler.UpdateRecipeHandler)

	err := router.Run()
	if err != nil {
		return
	}
}
