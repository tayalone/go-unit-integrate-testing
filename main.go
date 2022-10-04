package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tayalone/go-unit-integrate-test/myoperand"
)

func main() {

	r := gin.Default()
	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	// Binding from JSON
	type Operand struct {
		A int `json:"a"  binding:"required"`
		B int `json:"b"  binding:"required"`
		C int `json:"c"  binding:"required"`
		D int `json:"d"  binding:"required"`
	}

	r.POST("/operand", func(c *gin.Context) {
		var json Operand
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := myoperand.MyOperand(json.A, json.B, json.C, json.D)

		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	})

	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
