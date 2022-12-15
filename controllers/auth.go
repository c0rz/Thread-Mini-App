package controllers

import (
	"Blog/config"
	"Blog/middleware"
	"Blog/models"
	"Blog/repository"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var (
		input models.RegisterAccount
		user  models.Account
		count int
		ve    validator.ValidationErrors
	)

	if err := c.ShouldBindJSON(&input); err != nil {
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

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 16)

	user.Email = input.Email
	user.Full_Name = input.Name
	user.Level = "Member"
	user.Password = string(hashedPassword)

	repository.GetbyIDAccount(config.Database(), "SELECT COUNT(*) FROM account WHERE email = '"+input.Email+"'").Scan(&count)

	if count == 0 {
		err := repository.InsertAccount(config.Database(), user)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "Registration success",
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"code":    http.StatusBadRequest,
		"message": "Email already registered",
	})
}

func Login(c *gin.Context) {
	var (
		input   models.LoginInput
		ve      validator.ValidationErrors
		account models.Account
	)

	if err := c.ShouldBindJSON(&input); err != nil {
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

	repository.GetbyIDAccount(config.Database(), "SELECT * FROM account WHERE email = '"+input.Email+"';").Scan(&account.ID, &account.Full_Name, &account.Email, &account.Password, &account.Level, &account.CreatedAt)

	if account.Password != "" {
		err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(input.Password))
		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Password wrong.",
			})
			return
		}
		token, err := middleware.GenerateToken(account)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"code":  http.StatusOK,
			"token": "Bearer " + token,
		})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"code":    http.StatusUnauthorized,
		"message": "Email not found.",
	})
}
