package middlewares

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"Assignment_3/src/response"
	"Assignment_3/src/services"
)

// AuthorizeJWT validate the token user given
func AuthorizeJWT(service services.JWTService) gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization") // Authorization : token
		if authHeader == "" {
			res := response.BuildErrResponse("Failed", "No Token Found", nil)
			context.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		token, err := service.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]", claims["user_id"])
			log.Println("Claim[issure]", claims["issure"])
		} else {
			log.Println(err)
			res := response.BuildErrResponse("Token is not valid", err.Error(), nil)
			context.AbortWithStatusJSON(http.StatusUnauthorized, res)
		}
	}
}
