package middleware

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/hadyanadam/go-gorm-gin/helper"
	"github.com/hadyanadam/go-gorm-gin/service"
)


func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader, errAuth := helper.SplitHeader(c.GetHeader("Authorization"), "Bearer")
		if errAuth != nil {
			response := helper.BuildErrorResponse("Authorization header invalid", "Must include Bearer before token", helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid{
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id] : ", claims["user_id"])
			log.Println("Claim[issuer] : ", claims["issuer"])
		} else {
			log.Println(err)
			response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}