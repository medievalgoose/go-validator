package main

import (
	"medievalgoose/cc-validator/util"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Validator Endpoints
	router.GET("/validate/cc/:number", getValidateCreditCard)
	router.GET("/validate/email/:address", getValidateEmail)

	router.GET("/rot13/", getRot13Cipher)

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

func getRot13Cipher(ctx *gin.Context) {
	message := ctx.Query("message")

	if message == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Include the message in query"})
		return
	}

	var eliminatedSpaceMessage string

	for _, char := range message {
		if string(char) != " " && strings.Contains(util.Charset, strings.ToLower(string(char))) {
			eliminatedSpaceMessage += strings.ToLower(string(char))
		}
	}

	// TODO: Make sure the omit the symbols from the message.

	encodedMessage := util.Rot13Encode(eliminatedSpaceMessage)

	ctx.JSON(http.StatusOK, gin.H{"Input": message, "Output": encodedMessage})
}
