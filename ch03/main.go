package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"io/ioutil"
	"log"
	"time"
)

type Recipe struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}

var recipes []Recipe

var ctx context.Context
var err error
var clinet *mongo.Client

func init() {
	recipes = make([]Recipe, 0)
	fmt.Println("init")
	file, _ := ioutil.ReadFile("data/recipes.json")
	_ = json.Unmarshal([]byte(file), &recipes)
	fmt.Println(recipes)
	MONGO_URI := "mongodb://admin:ningwang@localhost:27017/test?authSource=admin"
	ctx = context.Background()
	clinet, err = mongo.Connect(ctx, options.Client().ApplyURI(MONGO_URI))
	if err = clinet.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connnected to MongoDB")

	MONGO_DATABASE := "test"
	collection := clinet.Database(MONGO_DATABASE).Collection("recipes")
	var listOfRecipes []interface{}
	for _, recipe := range recipes {
		listOfRecipes = append(listOfRecipes, recipe)
	}
	inserManyResult, err := collection.InsertMany(ctx, listOfRecipes)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted recipes: ", len(inserManyResult.InsertedIDs))
}

func main() {
}
