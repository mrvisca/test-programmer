package helpers

import "github.com/gin-gonic/gin"

func DataResponse(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"data": data,
	})
}
