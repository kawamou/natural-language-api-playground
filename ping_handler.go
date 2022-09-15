package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingHandler ...
type PingHandler struct{}

// PingHandlerRequest ...
type PingHandlerRequest struct{}

// PingHandlerResponse ...
type PingHandlerResponse struct {
	Message string `json:"message"`
}

// NewPingHandler ...
func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

// Handle ...
func (p *PingHandler) Handle(c *gin.Context) {
	response := PingHandlerResponse{
		Message: "pong",
	}
	c.JSON(http.StatusOK, response)
	return
}
