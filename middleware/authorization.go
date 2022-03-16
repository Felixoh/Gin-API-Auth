package middleware

import (
	"net/http"

	"github.com/Felixoh/Tokens"
	"github.com/gin-gonic/gin"
)

func ValidateToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.Request.Header.Get("apikey")
		referer := context.Request.Header.Get("Referer")
		valid, claims := Tokens.VerifyToken(tokenString, referer)

		if !valid {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Status": http.StatusUnauthorized, "message": "Unauthorised access, you are not authorised to access this path"})

		}

		if len(context.Keys) == 0 {
			context.Keys = make(map[string]interface{})
		}
		context.Keys["CompanyId"] = claims.CompanyId
		context.Keys["Username"] = claims.Username
		context.Keys["Roles"] = claims.Roles
	}
}

func Authorization(validRoles []int) gin.HandlerFunc {
	return func(context *gin.Context) {
		if len(context.Keys) == 0 {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Unaothorised access "})
		}
		rolesValue := context.Keys["Roles"]
		if rolesValue == nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Unaothorised access "})
		}
		roles := rolesValue.([]int)

		validation := make(map[int]int)
		for _, val := range roles {
			validation[val] = 0
		}

		for _, val := range validRoles {
			if _, ok := validation[val]; !ok {
				context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Unaothorised access "})
			}

		}

	}
}
