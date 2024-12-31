package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stanwar/sky-bot/Controllers"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/sky-bot")
	{
		grp1.GET("task", Controllers.GetTasks)
		grp1.POST("task", Controllers.InsertJob)
		grp1.POST("/webhook", Controllers.WebhookHandler)
	}
	return r
}
