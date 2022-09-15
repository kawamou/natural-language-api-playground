package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AnalyzeEntityHandler ...
type AnalyzeEntityHandler struct {
	c *Client
}

// AnalyzeAnalyzeEntityHandlerResponse ...
type AnalyzeEntityHandlerResponse struct {
	Entities Entities `json:"entities"`
}

// NewAnalyzeEntityHandler ...
func NewAnalyzeEntityHandler(c *Client) *AnalyzeEntityHandler {
	return &AnalyzeEntityHandler{
		c: c,
	}
}

// Handle ...
func (a *AnalyzeEntityHandler) Handle(c *gin.Context) {
	ctx := c.Request.Context()
	text := c.Param("text")
	ret, err := a.c.AnalyzeEntity(ctx, text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	response := &AnalyzeEntityHandlerResponse{
		Entities: ret.Entities,
	}
	c.JSON(http.StatusOK, response)
}
