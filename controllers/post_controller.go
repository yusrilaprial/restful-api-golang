package controllers

import (
	"errors"
	"net/http"
	"yusrilaprial/backend-api/db"
	"yusrilaprial/backend-api/models"
	"yusrilaprial/backend-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func FindPosts(c *gin.Context) {
	var posts []models.Post
	db.DB.Find(&posts)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Lists Data Posts",
		"data": posts,
	})
}

func StorePost(c *gin.Context) {
	var input models.ValidatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]models.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = models.ErrorMsg{Field: fe.Field(), Message: utils.GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
	}
	db.DB.Create(&post)

	c.JSON(201, gin.H{
		"success": true,
		"message": "Post Created Successfully",
		"data":    post,
	})
}
