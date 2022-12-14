package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

const port string = ":8080"

// Router
type Router struct {
	Engine *gin.Engine
}

// NewRouter ...
func NewRouter() *Router {
	engine := gin.New()
	return &Router{
		Engine: engine,
	}
}

// Run ...
func (r *Router) Run(port string) error {
	ctx := context.Background()
	c := NewClient(ctx)
	pingHandler := NewPingHandler()
	nyuwaizeHandler := NewNyuwaizeHandler(c)
	analyzeSentimentHandler := NewAnalyzeSentimentHandler(c)
	analyzeEntityHandler := NewAnalyzeEntityHandler(c)
	analyzeEntitySentimentHandler := NewAnalyzeEntitySentimentHandler(c)
	api := r.Engine.Group("")
	api.GET("/ping", pingHandler.Handle)
	// api.POST("/analyze-syntax/:text")
	api.GET("/nyuwaize/:text", nyuwaizeHandler.Handle)
	api.GET("analyze-sentiment/:text", analyzeSentimentHandler.Handle)
	api.GET("analyze-entity/:text", analyzeEntityHandler.Handle)
	api.GET("analyze-entity-sentiment/:text", analyzeEntitySentimentHandler.Handle)
	return r.Engine.Run(port)
}

func main() {
	log.Println(NewRouter().Run(port))
}
