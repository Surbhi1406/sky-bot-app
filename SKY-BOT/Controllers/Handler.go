package Controllers

import (
	"fmt"
	"net/http"

	"github.com/stanwar/sky-bot/Database"
	"github.com/stanwar/sky-bot/Models"

	"github.com/gin-gonic/gin"
)

var skyBotEngnURL = "http://localhost:8082/execute-task"
var webhookURL = "http://localhost:8080/sky-bot/webhook"

// GetBooks ... Get all books
func GetTasks(c *gin.Context) {
	var task []Models.Task
	err := Database.FetchAllTasks(&task)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, task)
	}
}

// CreateBook ... Create Book
func InsertJob(c *gin.Context) {
	fmt.Printf("End point - InsertJob")
	var task Models.Task
	c.BindJSON(&task)

	err := Database.InsertTaskIntoTable(Database.DB, "queued_jobs", task)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert task"})
	} else {
		fmt.Printf("Data inserted successfully!!")
		fmt.Printf("task = %s", task)
		sendRequestToSkybot(skyBotEngnURL, webhookURL, task)
		c.JSON(http.StatusOK, task)
	}

}
