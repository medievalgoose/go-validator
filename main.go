package main

import (
	"medievalgoose/cc-validator/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/validatecc/:number", getValidateCreditCard)

	router.Run("localhost:8080")
}

func getValidateCreditCard(c *gin.Context) {
	ccNumber := c.Param("number")
	validCC := util.ValidateCreditCard(ccNumber)

	if validCC {
		c.JSON(http.StatusOK, gin.H{"Verdict": "VALID"})
	} else {
		c.JSON(http.StatusOK, gin.H{"Verdict": "INVALID"})
	}
}
