package controllers

import (
	"Blog/config"
	"Blog/middleware"
	"Blog/models"
	"Blog/repository"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func AddComment(c *gin.Context){
	var (
		InputC models.CommentInput
		ve        validator.ValidationErrors
		comment      models.Comment
		db        = config.Database()
	)

	defer db.Close()

	if err := c.ShouldBindJSON(&InputC); err != nil {
		if errors.As(err, &ve) {
			var errorMessage []string
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage = append(errorMessage, "Error on field "+e.Field()+", conditon: "+e.ActualTag())
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  http.StatusBadRequest,
				"error": errorMessage,
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Terjadi kesalahan pada inputan.",
		})
		return
	}

	id, _, err := middleware.ExtractTokenID(c)
	if err != nil {
		panic(err)
	}

	comment.ID_Post = uint64(InputC.IDPost)
	comment.ID_User = strconv.Itoa(int(id))
	comment.Comment = InputC.Comment

	err = repository.InsertComment(db, comment)
	if err != nil {
		c.JSON(http.StatusPartialContent, gin.H{
			"code":  http.StatusPartialContent,
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"response": "Insert comment success.",
	})
}