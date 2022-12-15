package controllers

import (
	"Blog/config"
	"Blog/models"
	"Blog/repository"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateCategory(c *gin.Context) {
	var (
		CategoryInput models.CategoryInput
		ve        validator.ValidationErrors
		category      models.Category
		db        = config.Database()
	)

	defer db.Close()

	if err := c.ShouldBindJSON(&CategoryInput); err != nil {
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

	category.Name = CategoryInput.Name

	err := repository.InsertCategory(db, category)
	if err != nil {
		c.JSON(http.StatusPartialContent, gin.H{
			"code":  http.StatusPartialContent,
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"response": "Success insert category.",
	})
}

func DeleteCategory(c *gin.Context) {
	var (
		db    = config.Database()
	)

	defer db.Close()

	pid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":     http.StatusBadRequest,
			"response": "Error Bad Request.",
		})
		return
	}
	
	err = repository.DeleteCategory(db, pid)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"response": "Success delete category.",
	})
}