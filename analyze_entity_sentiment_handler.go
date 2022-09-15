package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AnalyzeEntitySentimentHandler ...
type AnalyzeEntitySentimentHandler struct {
	c *Client
}

// AnalyzeEntitySentimentHandlerResponse ...
type AnalyzeEntitySentimentHandlerResponse struct {
	Entities Entities `json:"entities"`
}

// NewAnalyzeEntitySentimentHandler ...
func NewAnalyzeEntitySentimentHandler(c *Client) *AnalyzeEntitySentimentHandler {
	return &AnalyzeEntitySentimentHandler{
		c: c,
	}
}

// Handle ...
func (a *AnalyzeEntitySentimentHandler) Handle(c *gin.Context) {
	ctx := c.Request.Context()
	text := c.Param("text")
	ret, err := a.c.AnalyzeEntitySentiment(ctx, text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	response := &AnalyzeEntitySentimentHandlerResponse{
		Entities: ret.Entities,
	}
	c.JSON(http.StatusOK, response)
}
