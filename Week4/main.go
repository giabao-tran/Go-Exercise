package main

import (
	"jwt-authentication-golang/handlers"
	"jwt-authentication-golang/initializers"
	"jwt-authentication-golang/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDb()
	initializers.SyncDb()
}

func main() {
	r := gin.Default()

	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.Login)

	authorized := r.Group("/")
	authorized.Use(middlewares.RequireAuth)
	{
		authorized.GET("/profile", handlers.ViewProfile)
		authorized.PUT("/profile", handlers.EditProfile)
		authorized.POST("/posts", handlers.CreatePost)
		authorized.GET("/posts/:id", handlers.GetPost)
		authorized.PUT("/posts/:id", handlers.UpdatePost)
		authorized.POST("/posts/:id/comments", handlers.AddComment)
		authorized.POST("/posts/:id/like", handlers.LikePost)
	}

	r.GET("/profile/:username", handlers.ViewOtherProfile)

	r.Run()
}
