package main

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	name := c.Params.ByName("name")
	c.XML(200, Person{
		FirstName: name,
		LastName:  "Labouardy",
	})

}

type Person struct {
	XMLName   xml.Name `xml:"person"`
	FirstName string   `xml:"firstName,attr"`
	LastName  string   `xml:"lastName,attr"`
}

func main() {
	router := gin.Default()
	router.GET("/:name", IndexHandler)
	err := router.Run(":5000")

	if err != nil {
		return
	}
}
