package main

import (
	"medievalgoose/cc-validator/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/validate/cc/:number", getValidateCreditCard)

	router.Run("localhost:8080")
}

func getValidateCreditCard(ctx *gin.Context) {
	ccNumber := ctx.Param("number")
	validCC, classification := util.ValidateCreditCard(ccNumber)

	if validCC {
		ctx.JSON(http.StatusOK, gin.H{"Provider": classification, "Verdict": "VALID"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"Provider": classification, "Verdict": "INVALID"})
	}
}
