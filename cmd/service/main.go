package main

import (
	"go-gin-demo/internal/handlers"
	"go-gin-demo/internal/middlewares"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	authRoutes := r.Group("/").Use(middlewares.AuthMiddleware())

	authRoutes.GET("/api/equityPositions", handlers.GetEquityPositionsHandler)

	err := r.Run("localhost:8080")
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
