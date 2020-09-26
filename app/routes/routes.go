package routes

import (
	controllers "emailer_service/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine) {
	r.POST("/send", controllers.Produce)
	r.GET("/recv", controllers.Receive)
}
