package main

import (
	"medievalgoose/cc-validator/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/validate/cc/:number", getValidateCreditCard)
	router.GET("/validate/email/:address", getValidateEmail)

	router.Run("localhost:8080")
}

func getValidateCreditCard(ctx *gin.Context) {
	ccNumber := ctx.Param("number")

	for _, char := range ccNumber {
		_, err := strconv.Atoi(string(char))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "only numbers are allowed."})
			return
		}
	}

	validCC, classification := util.ValidateCreditCard(ccNumber)

	if validCC {
		ctx.JSON(http.StatusOK, gin.H{"Provider": classification, "Verdict": "VALID"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"Provider": classification, "Verdict": "INVALID"})
	}
}

func getValidateEmail(ctx *gin.Context) {
	emailAddress := ctx.Param("address")
	validEmail := util.ValidateEmail(emailAddress)

	if validEmail {
		ctx.JSON(http.StatusOK, gin.H{"Email": emailAddress, "Verdict": "VALID"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"Email": emailAddress, "Verdict": "INVALID"})
	}
}
