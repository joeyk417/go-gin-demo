package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const (
	authorizationHeaderKey  = "Authorization"
	authorizationTypeBearer = "bearer"
	testToken               = "fJCoxhq8uR9GiUIgaIGfMgw7zCqxwDhQ"
)

// handle authentication and handle an invalid token being passed in
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get Authorization Header
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header is not provided"})
			log.Error("authorization header is not provided")
			return
		}

		// check Authorization Header format
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			log.Error("invalid authorization header format")
			return
		}

		// check Authorization Type
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unsupported authorization type"})
			log.Error("unsupported authorization type")
			return
		}

		// verify token
		token := fields[1]
		if token != testToken {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			log.Error("invalid token")
			return
		}

		ctx.Next()
	}
}
