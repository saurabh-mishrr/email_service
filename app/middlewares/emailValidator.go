package middlewares

import (
	configs "emailer_service/configs"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Validate(c *gin.Context) {
	requestPayloads := &configs.EmailRequestPayLoad{}
	c.BindJSON(requestPayloads)
	fmt.Println("Got the name ", requestPayloads.Name)
	c.Next()
}
