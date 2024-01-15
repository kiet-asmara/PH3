package controllers

import (
	"net/http"
	"rest-review/helpers"
	"rest-review/models"
	"rest-review/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	UserService *models.UserService
}

// register controller
func (u Users) HandleRegister(c *gin.Context) {
	var input models.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorMessage(c, &utils.ErrBadInput, err)
		return
	}

	user, err := u.UserService.Create(input)
	if err != nil {
		utils.ErrorMessage(c, &utils.ErrInternalServer, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "register successful",
		"user":    user,
	})
}

// login controller
func (u Users) HandleLogin(c *gin.Context) {
	var input models.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorMessage(c, &utils.ErrBadInput, err)
		return
	}

	user, err := u.UserService.Read(input)
	if err != nil {
		utils.ErrorMessage(c, &utils.ErrInternalServer, err)
		return
	}
	// TODO: need to separate business logic
	if user.UserID == 0 {
		utils.ErrorMessage(c, &utils.ErrDataNotFound, err)
		return
	}

	// check password
	passCheck := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if passCheck != nil {
		utils.ErrorMessage(c, &utils.ErrUnauthorized, err)
		return
	}

	// generate JWT
	token, err := helpers.GenerateJWT(user.UserID)
	if err != nil {
		utils.ErrorMessage(c, &utils.ErrInternalServer, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "login successful",
		"token":   token,
	})
}
