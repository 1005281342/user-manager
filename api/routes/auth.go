package routes

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/1005281342/user-manager/auth"
	"github.com/1005281342/user-manager/db"
)

func SetupAuthRoutes(r *gin.Engine) {
	jwtAuth := auth.NewJWTAuth("secret-key")
	auth := r.Group("/auth")

	auth.POST("/login", func(c *gin.Context) {
		var login struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := db.GetUserByEmail(login.Email)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
			return
		}

		token, err := jwtAuth.CreateToken(fmt.Sprintf("%v", user.ID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	auth.POST("/logout", func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "token not found"})
			return
		}

		err := jwtAuth.DeleteToken(tokenString)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "logout success"})
	})
}
