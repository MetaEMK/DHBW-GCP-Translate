package api

import "github.com/gin-gonic/gin"


func sendApiCall(c *gin.Context) {
    println("Send API Call")
    if len(c.Errors) > 0 {
        c.JSON(400, gin.H{"error": c.Errors})
    } else {
        c.JSON(200, gin.H{"message": "success"})
    }
}
