package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trickqz/user-auth-token/models"
)

func GetUserProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id não encontrado"})
		return
	}

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ID":       user.ID,
		"Email":    user.Email,
		"Username": user.Username,
		"Avatar":   user.Avatar,
	})
}

func UpdateUserProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id não encontrado"})
		return
	}

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	var updatedData struct {
		Username string `json:"username"`
		Avatar   string `json:"avatar"`
	}
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	if updatedData.Username != "" {
		user.Username = updatedData.Username
	}
	if updatedData.Avatar != "" {
		user.Avatar = updatedData.Avatar
	}

	if err := db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar perfil"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
