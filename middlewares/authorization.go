package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/memclutter/gontacts/utils"
)

func Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// extract authorization header
		headerString := ctx.GetHeader("Authorization")

		if len(headerString) == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Missing authorization header",
			})
		} else if headerPairs := strings.Split(headerString, " "); len(headerPairs) != 2 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization header invalid format",
			})
		} else if strings.ToUpper(headerPairs[0]) != "JWT" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization header invalid format",
			})
		} else if claims, err := utils.CheckJwtToken(headerPairs[1]); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
		} else {
			ctx.Set("claims", claims)
			ctx.Next()
		}
	}
}
