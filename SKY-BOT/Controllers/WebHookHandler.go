package Controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/stanwar/sky-bot/Database"
	"github.com/stanwar/sky-bot/Models"
)

// Webhook handler to process notifications
func WebhookHandler(c *gin.Context) {
	var webhookResponse Models.WebhookResponse
	if err := c.ShouldBindJSON(&webhookResponse); err != nil {
		c.JSON(400, gin.H{"error": "Failed to parse JSON"})
		return
	}

	fmt.Println("Received webhook notification:", webhookResponse)

	switch {
	case webhookResponse.Status == "started":
		fmt.Println("started block!!")
		Database.MoveTaskBetweenTables(Database.DB, "queued_jobs", "inprogress_jobs", webhookResponse.JobID)
	case webhookResponse.Status == "wait":
	case webhookResponse.Status == "completed":
		{
			fmt.Println("completed block!!")
			Database.MoveTaskBetweenTables(Database.DB, "inprogress_jobs", "completed_jobs", webhookResponse.JobID)
			task, error := Database.GetFirstQueuedJob(Database.DB, "queued_jobs")
			if error != nil {
				fmt.Println(" error error error!!!")
				fmt.Println(error)
				return
			}
			fmt.Println("+++++++++++++++ ", task)
			Database.MoveTaskBetweenTables(Database.DB, "queued_jobs", "inprogress_jobs", webhookResponse.JobID)
			sendRequestToSkybot(skyBotEngnURL, webhookURL, task)

			fmt.Println("completed block ends!!")
		}
	case webhookResponse.Status == "error":
	default:
		c.JSON(500, gin.H{"message": "Invalid state recived from skybot"})
	}
	c.JSON(200, gin.H{"message": "Webhook received"})
}
