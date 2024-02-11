package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)          // All event
	server.GET("/events/:id", getEvent)       // Single event
	server.POST("/events", createEvents)      // Create event
	server.PUT("/events/:id", updateEvent)    // Update event
	server.DELETE("/events/:id", deleteEvent) // Delete event
	server.POST("/signup", signup)
	server.POST("/login", login)
}
