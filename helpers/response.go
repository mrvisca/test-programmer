package helpers

import "github.com/gin-gonic/gin"

func DataResponse(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"data": data,
	})
}

func SuksesWithDataResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(201, gin.H{
		"status":  "Sukses",
		"message": message,
		"data":    data,
	})
}

func SuksesResponse(c *gin.Context, message string) {
	c.JSON(200, gin.H{
		"status":  "Sukses",
		"message": message,
	})
}

func ErrorResponse(c *gin.Context, message string) {
	c.JSON(400, gin.H{
		"status":  "Error",
		"message": message,
	})
}
