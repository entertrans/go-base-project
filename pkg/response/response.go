package response

import "github.com/gin-gonic/gin"

type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SendResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, APIResponse{
		Code:    statusCode,
		Message: message,
		Data:    data,
	})
}

func SendErrorResponse(c *gin.Context, statusCode int, err string) {
	SendResponse(c, statusCode, err, nil)
}
