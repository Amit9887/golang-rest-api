package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type sample struct {
	Id        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var samples = []sample{
	{Id: "1", Item: "laptop", Completed: true},
	{Id: "2", Item: "mobile", Completed: false},
	{Id: "3", Item: "mouse", Completed: true},
	{Id: "4", Item: "charger", Completed: false},
}

func getData(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, samples)
}

func addData(context *gin.Context) {
	var newSample sample

	if err := context.BindJSON(&newSample); err != nil {
		return
	}

	samples = append(samples, newSample)
	context.IndentedJSON(http.StatusCreated, newSample)
}

func getDatas(context *gin.Context) {
	id := context.Param("id")
	sample, err := getDataById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Data Not Found"})
		return
	}

	context.IndentedJSON(http.StatusOK, sample)

}

func getDataById(id string) (*sample, error) {
	for i, d := range samples {
		if d.Id == id {
			return &samples[i], nil
		}
	}
	return nil, errors.New("Data Not Found")

}

func updateData(context *gin.Context) {
	id := context.Param("id")
	sample, err := getDataById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Data Not Found"})
		return
	}

	sample.Completed = !sample.Completed
	context.IndentedJSON(http.StatusOK, sample)
}

func main() {
	router := gin.Default()
	router.GET("/get", getData)
	router.POST("/post", addData)
	router.GET("/get/:id", getDatas)
	router.PATCH("/patch/:id", updateData)
	router.Run("localhost:9090")
}
