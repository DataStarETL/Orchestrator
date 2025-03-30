package rest

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func createResponseMetadata(status bool, errorCode string, errorMessage string) map[string]string {
	var metaData = make(map[string]string)
	if status {
		metaData["status"] = "success"
	} else {
		metaData["status"] = "failed"
	}
	if errorCode != "" {
		metaData["error_code"] = errorCode
	}

	if errorMessage != "" {
		metaData["error_message"] = errorMessage
	}

	return metaData
}

func createErrorResponse(c *gin.Context, statusCode int, errorMsg string) {
	metaData := createResponseMetadata(false, strconv.Itoa(statusCode), errorMsg)
	c.JSON(statusCode, gin.H{
		"meta": metaData,
		"data": nil,
	})
}

func createDataResponse(c *gin.Context, statusCode int, data any) {
	var metaData = createResponseMetadata(true, "", "")
	c.JSON(statusCode, gin.H{
		"meta": metaData,
		"data": data,
	})
}
