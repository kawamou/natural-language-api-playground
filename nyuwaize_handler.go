package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NyuwaizeHandler ...
type NyuwaizeHandler struct {
	c *Client
}

// NyuwaizeHandlerResponse ...
type NyuwaizeHandlerResponse struct {
	Message string `json:"message"`
}

// NewNyuwaizeHandler ...
func NewNyuwaizeHandler(c *Client) *NyuwaizeHandler {
	return &NyuwaizeHandler{
		c: c,
	}
}

// Handle ...
func (n *NyuwaizeHandler) Handle(c *gin.Context) {
	ctx := c.Request.Context()
	text := c.Param("text")
	ret, err := n.c.Nyuwaize(ctx, text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	response := &NyuwaizeHandlerResponse{
		Message: ret,
	}
	c.JSON(http.StatusOK, response)
	return
}
