package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/Nandhagopiv/paper-social-go/controllers"
)

func SetupRoutes(r *gin.Engine) {
    r.POST("/register", controllers.Register)
	r.POST("/tweet", controllers.CreateTweet)
    r.GET("/tweets", controllers.GetTweets)
}
