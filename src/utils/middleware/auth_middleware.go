package middleware

import (
	"github.com/ashishkhuraishy/blogge/src/services"
	"github.com/ashishkhuraishy/blogge/src/utils/errors/resterror"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware checks for a token validates it then redirects
// the request to secured endpoint
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("authorization")
		if authHeader == "" || len(authHeader) < len("Token")+1 {
			restErr := resterror.NewUnAuthorizedError()
			c.JSON(restErr.StatusCode, restErr)
			c.Abort()
			return
		}

		token := authHeader[len("Token "):]

		authService := services.JWTAuthService()
		result, err := authService.ValidateToken(token)
		if err != nil || !result.Valid {
			restErr := resterror.NewUnAuthorizedError()
			c.JSON(restErr.StatusCode, restErr)
			c.Abort()
			return
		}

		claims := result.Claims.(jwt.MapClaims)
		c.Set("user_id", claims["user_id"])
		c.Set("is_admin", claims["is_admin"])

		c.Next()
	}
}
