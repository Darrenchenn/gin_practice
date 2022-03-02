package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandle() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
		for _, err := range context.Errors {
			context.JSON(http.StatusOK, err.JSON())
			return
		}
	}
}
