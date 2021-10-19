package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Recipe struct {
	Name        string    `json:"name"`
	Tags        []string  `json:"tags"`
	Ingredients []string  `json:"ingredients"`
	PublishedAt time.Time `json:"publishedAt"`
}

func main() {
	router := gin.Default()
	err := router.Run()
	if err != nil {
		return
	}
}
