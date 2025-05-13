package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/Nandhagopiv/paper-social-go/models"
    "github.com/Nandhagopiv/paper-social-go/utils"
    "golang.org/x/crypto/bcrypt"
    "net/http"
)

func Register(c *gin.Context) {
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
    user := models.User{Username: input.Username, Email: input.Email, Password: string(hashedPassword)}
    result := utils.DB.Create(&user)

    if result.Error != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User registered"})
}