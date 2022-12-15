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

func GetAllThread(c *gin.Context) {
	var (
		db = config.Database()
	)

	defer db.Close()

	err, blog := repository.GetAllPost(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"result": blog,
	})
}

func SubmitThread(c *gin.Context) {
	var (
		BlogInput models.BlogInput
		ve        validator.ValidationErrors
		blog      models.Post
		db        = config.Database()
	)

	defer db.Close()

	if err := c.ShouldBindJSON(&BlogInput); err != nil {
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

	blog.Text = BlogInput.Text
	blog.Title = BlogInput.Title
	blog.ID_User = strconv.Itoa(int(id))
	blog.ID_Category = BlogInput.CategoryID

	idPost, err := repository.InsertPost(db, blog)
	if err != nil {
		c.JSON(http.StatusPartialContent, gin.H{
			"code":  http.StatusPartialContent,
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"id_post":  idPost,
		"response": "Success publish.",
	})
}

func DeleteThread(c *gin.Context) {
	var (
		db    = config.Database()
		count int
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

	id, _, err := middleware.ExtractTokenID(c)
	if err != nil {
		panic(err)
	}

	repository.GetbyIDPost(db, "SELECT COUNT(*) FROM post WHERE id = "+strconv.Itoa(pid)+" AND id_user = "+strconv.Itoa(int(id))+";").
		Scan(&count)

	if count == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":     http.StatusUnauthorized,
			"response": "Not found post.",
		})
		return
	}

	err = repository.DeletePost(db, pid)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"response": "Success delete post.",
	})
}

func UpdateThread(c *gin.Context) {
	var (
		BlogInput models.BlogInput
		ve        validator.ValidationErrors
		blog      models.Post
		db        = config.Database()
		count     int
	)

	pid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":     http.StatusBadRequest,
			"response": "Error Bad Request.",
		})
		return
	}

	defer db.Close()

	if err := c.ShouldBindJSON(&BlogInput); err != nil {
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

	repository.GetbyIDPost(db, "SELECT COUNT(*) FROM post WHERE id = "+strconv.Itoa(pid)+" AND id_user = "+strconv.Itoa(int(id))+";").
		Scan(&count)

	if count == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":     http.StatusUnauthorized,
			"response": "Not found post.",
		})
		return
	}

	blog.ID = uint64(pid)
	blog.Text = BlogInput.Text
	blog.Title = BlogInput.Title
	blog.ID_Category = BlogInput.CategoryID

	err = repository.UpdatePost(db, blog)
	if err != nil {
		c.JSON(http.StatusPartialContent, gin.H{
			"code":  http.StatusPartialContent,
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"response": "Update success.",
	})
}

func GetThread(c *gin.Context) {
	var (
		post  models.Post
		db    = config.Database()
	)

	pid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":     http.StatusBadRequest,
			"response": "Error Bad Request.",
		})
		return
	}

	defer db.Close()

	repository.GetbyIDPost(db, `SELECT post.title, category.name, text, account.full_name, post.created_at
	FROM post
	JOIN account ON account.id = post.id_user
	JOIN category ON category.id = post.id_category
	WHERE post.id = '`+strconv.Itoa(pid)+`';`).Scan(&post.Title, &post.ID_Category, &post.Text, &post.ID_User, &post.CreatedAt)

	if post.Title == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":     http.StatusUnauthorized,
			"response": "Not found post.",
		})
		return
	}

	post.Comment, err = repository.GetComment(db, strconv.Itoa(pid))
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"response": post,
	})
}
