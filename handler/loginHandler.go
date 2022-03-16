package handler

import (
	"net/http"
	"time"

	"github.com/Felixoh/Tokens"
	"github.com/Felixoh/models"
	"github.com/gin-gonic/gin"
)

func LoginHandler(context *gin.Context) {
	var loginObj models.LoginRequest

	if err := context.ShouldBindJSON(&loginObj); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error(),"satus":http.StatusBadRequest,"message":"Incorrect Arguments"})
		return
	}

	//validate the loginObj for valid credentials and add on if they are valid then save your
	//validations done here
	var claims = &models.JwtClaims{}
	claims.CompanyId = "CompanyId"
	claims.Username = loginObj.Username
	claims.Roles = []int{1, 2, 3, 4}

	//get the refereree from the header
	claims.Audience = context.Request.Header.Get("Referer")

	var tokenCreationTime = time.Now().UTC()

	//adjust the token expiration to any preferred expiry time i.e in Second,Minutes or Hours.
	var expirationTime = tokenCreationTime.Add(time.Duration(10) * time.Minute)

	tokenString, err := Tokens.GenerateToken(claims, expirationTime)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"success": "token has been created", "token": tokenString})

}
