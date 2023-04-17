package routes

import (
	"github.com/1005281342/user-manager/auth"
	"github.com/1005281342/user-manager/cache"
	"github.com/1005281342/user-manager/db"
	"github.com/1005281342/user-manager/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

func SetupUserRoutes(r *gin.Engine, auth auth.Auth) {
	user := r.Group("/user")

	user.GET("/:id", GetUser)
	user.POST("/", CreateUser)
	user.PUT("/:id", authMiddleware(auth), UpdateUser)
	user.DELETE("/:id", authMiddleware(auth), DeleteUser)

	r.GET("/users", ListUsers)
	r.GET("/users/search", SearchUsers)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	redisCache := cache.NewRedisCache()
	defer redisCache.Close()

	// Check if the user is already in the cache
	var user User
	err := redisCache.Get("user:"+id, &user)
	if err == nil {
		c.JSON(http.StatusOK, user)
		return
	} else if err != db.ErrNoResult {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// If not, query the database
	query := "SELECT * FROM users WHERE id = $1"
	row := db.GetDB().QueryRow(query, id)
	err = row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Add the user to the cache for future requests
	err = redisCache.Set("user:"+id, user, 5*time.Minute)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Check if user exists
	var existingUser User
	query := "SELECT * FROM users WHERE id = $1"
	row := db.GetDB().QueryRow(query, id)
	if err := row.Scan(&existingUser.ID, &existingUser.FirstName, &existingUser.LastName, &existingUser.Email, &existingUser.Password); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Update user
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	err = updateUserInDBAndCache(id, user.FirstName, user.LastName, user.Email, string(hash))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func updateUserInDBAndCache(id int, firstName string, lastName string, email string, password string) error {
	tx, err := db.GetDB().Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	query := "UPDATE users SET first_name = $1, last_name = $2, email = $3, password = $4 WHERE id = $5"
	_, err = tx.Exec(query, firstName, lastName, email, password, id)
	if err != nil {
		return err
	}

	// Update the user in cache
	redisCache := cache.NewRedisCache()
	defer redisCache.Close()
	err = redisCache.Set("user:"+strconv.Itoa(id), User{ID: id, FirstName: firstName, LastName: lastName, Email: email, Password: password}, 5*time.Minute)
	if err != nil {
		return err
	}

	return nil
}

func CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	query := "INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4) RETURNING id"
	row := db.GetDB().QueryRow(query, user.FirstName, user.LastName, user.Email, string(hash))
	var id int
	if err := row.Scan(&id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.ID = id

	c.JSON(http.StatusCreated, user)
}

// api/routes/user.go

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	// Check if user exists
	var existingUser User
	query := "SELECT * FROM users WHERE id = $1"
	row := db.GetDB().QueryRow(query, id)
	if err := row.Scan(&existingUser.ID, &existingUser.FirstName, &existingUser.LastName, &existingUser.Email, &existingUser.Password); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Delete user
	query = "DELETE FROM users WHERE id = $1"
	_, err := db.GetDB().Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func ListUsers(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}
	perPage, err := strconv.Atoi(c.Query("per_page"))
	if err != nil || perPage < 1 {
		perPage = 10
	}
	offset := (page - 1) * perPage

	users, err := db.GetAllUsers(perPage, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get users",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func SearchUsers(c *gin.Context) {
	keyword := c.Query("keyword")
	perPage, err := strconv.Atoi(c.Query("per_page"))
	if err != nil {
		perPage = 20 // 默认每页显示20条数据
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1 // 默认显示第一页
	}

	var users []models.User

	if keyword == "" {
		users, err = db.GetAllUsersPerPage(perPage, page)
	} else {
		users, err = db.SearchUsersPerPage(keyword, perPage, page)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to search users",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

type User models.User
