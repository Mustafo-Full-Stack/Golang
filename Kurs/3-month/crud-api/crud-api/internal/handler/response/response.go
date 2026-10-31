package response

import "github.com/gin-gonic/gin"

type errorBody struct {
	Error string `json:"error"`
}

func Error(c *gin.Context, status int, message string) {
	c.AbortWithStatusJSON(status, errorBody{Error: message})
}
