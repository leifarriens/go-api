package helpers

import "github.com/gin-gonic/gin"

func httpError(msg string) map[string]any {
	return gin.H{"message": msg}
}
