package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/Nandhagopiv/paper-social-go/models"
    "github.com/Nandhagopiv/paper-social-go/utils"
)

// CreateTweet handles creating a new tweet
func CreateTweet(c *gin.Context) {
    var input struct {
        UserID  uint   `json:"user_id"`
        Content string `json:"content"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    tweet := models.Tweet{Content: input.Content, UserID: input.UserID}
    if err := utils.DB.Create(&tweet).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tweet"})
        return
    }

    c.JSON(http.StatusOK, tweet)
}

// GetTweets handles fetching all tweets
func GetTweets(c *gin.Context) {
    var tweets []models.Tweet
    if err := utils.DB.Preload("User").Order("created_at desc").Find(&tweets).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tweets"})
        return
    }

    c.JSON(http.StatusOK, tweets)
}
