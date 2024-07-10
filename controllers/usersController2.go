package controllers

import (
	"gin-backend/initializers"
	"gin-backend/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup2(c *gin.Context) {
	var UserBody struct {
		Email    string
		Password string
	}

	if err := c.Bind(&UserBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read the body",
		})
		return
	}

	//Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(UserBody.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash the password",
		})
		return
	}

	user := models.User{Email: UserBody.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

}

func Login1(c *gin.Context) {

	var UserBody struct {
		Email    string
		Password string
	}

	if err := c.Bind(&UserBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read the body",
		})
		return
	}

	//check if user exists or not
	var user models.User

	initializers.DB.First(&user, "email = ?", UserBody.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User Does not exists",
		})
		return
	}

	//check password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(UserBody.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong Password",
		})
		return
	}

	//Generate jwt

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Token creation failed",
			"error":   err,
		})
		return
	}

	log.Println(err)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})

}
