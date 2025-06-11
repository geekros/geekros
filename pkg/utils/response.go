package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func requestLog(c *gin.Context, status string, response string) {

}

func Success(c *gin.Context, data interface{}) {

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    data,
	})

	logJson, _ := json.Marshal(gin.H{"code": 0, "message": "success", "data": data})

	requestLog(c, "success", string(logJson))
}

func Warning(c *gin.Context, code int, message string, data interface{}) {

	if message == "" {
		message = "warning"
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})

	logJson, _ := json.Marshal(gin.H{"code": code, "message": message, "data": data})

	requestLog(c, "warning", string(logJson))
}

func Error(c *gin.Context, data interface{}) {

	c.JSON(http.StatusOK, gin.H{
		"code":    10000,
		"message": "error",
		"data":    data,
	})

	logJson, _ := json.Marshal(gin.H{"code": 10000, "message": "error", "data": data})

	requestLog(c, "error", string(logJson))
}
