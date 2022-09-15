package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AnalyzeSentimentHandler ...
type AnalyzeSentimentHandler struct {
	c *Client
}

// AnalyzeSentimentHandlerResponse ...
type AnalyzeSentimentHandlerResponse struct {
	PositiveNegative string `json:"positive_nagative"`
}

// NewAnalyzeSentimentHandler ...
func NewAnalyzeSentimentHandler(c *Client) *AnalyzeSentimentHandler {
	return &AnalyzeSentimentHandler{
		c: c,
	}
}

// Handle ...
func (a *AnalyzeSentimentHandler) Handle(c *gin.Context) {
	ctx := c.Request.Context()
	text := c.Param("text")
	ret, err := a.c.AnalyzeSentiment(ctx, text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	response := &AnalyzeSentimentHandlerResponse{
		PositiveNegative: ScoreToPositiveNegative(ret.Sentiment.Score),
	}
	c.JSON(http.StatusOK, response)
}

func ScoreToPositiveNegative(score float32) string {
	if score >= 0 {
		return "positive"
	} else {
		return "negative"
	}
}
