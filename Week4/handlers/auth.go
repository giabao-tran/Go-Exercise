package handlers

import (
	"fmt"
	"jwt-authentication-golang/initializers"
	"jwt-authentication-golang/models"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required"})
		return
	}

	file, err := c.FormFile("userProfile")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error retrieving the file: " + err.Error()})
		return
	}

	filename := generateUniqueFileName(file.Filename)
	uploadPath := filepath.Join("uploads", filename)

	if err := ensureUploadDir(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory: " + err.Error()})
		return
	}

	if err := c.SaveUploadedFile(file, uploadPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the file: " + err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password: " + err.Error()})
		return
	}

	user := models.User{
		Username:    username,
		Password:    string(hashedPassword),
		UserProfile: filename,
	}

	if err := initializers.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"profile":  filename,
		},
	})
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var user models.User
	if err := initializers.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func ViewProfile(c *gin.Context) {
	user, _ := c.Get("user")
	currentUser := user.(models.User)

	c.JSON(http.StatusOK, gin.H{
		"id":       currentUser.ID,
		"username": currentUser.Username,
		"profile":  currentUser.UserProfile,
	})
}

func EditProfile(c *gin.Context) {
	user, _ := c.Get("user")
	currentUser := user.(models.User)

	newUsername := c.PostForm("username")
	if newUsername != "" {
		currentUser.Username = newUsername
	}

	file, err := c.FormFile("userProfile")
	if err == nil {
		filename := generateUniqueFileName(file.Filename)
		uploadPath := filepath.Join("uploads", filename)

		if err := ensureUploadDir(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory: " + err.Error()})
			return
		}

		if err := c.SaveUploadedFile(file, uploadPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the file: " + err.Error()})
			return
		}

		currentUser.UserProfile = filename
	}

	if err := initializers.DB.Save(&currentUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile updated successfully",
		"user": gin.H{
			"id":       currentUser.ID,
			"username": currentUser.Username,
			"profile":  currentUser.UserProfile,
		},
	})
}

func ViewOtherProfile(c *gin.Context) {
	username := c.Param("username")

	var user models.User
	if err := initializers.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"profile":  user.UserProfile,
	})
}

func generateUniqueFileName(originalFilename string) string {
	return fmt.Sprintf("%s%s", uuid.New().String(), filepath.Ext(originalFilename))
}

func ensureUploadDir() error {
	return os.MkdirAll("uploads", os.ModePerm)
}
