package routes

import (
	"github.com/1005281342/user-manager/db"
	"github.com/1005281342/user-manager/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SetupUserRoutes(r *gin.Engine) {
	user := r.Group("/user")

	user.GET("/:id", GetUser)
	user.POST("/", CreateUser)
	user.PUT("/:id", UpdateUser)
	user.DELETE("/:id", DeleteUser)

	r.GET("/users", ListUsers)
	r.GET("/users/search", SearchUsers)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	var user User
	query := "SELECT * FROM users WHERE id = $1"
	row := db.GetDB().QueryRow(query, id)
	if err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "INSERT INTO users (first_name, last_name, email) VALUES ($1, $2, $3) RETURNING id"
	row := db.GetDB().QueryRow(query, user.FirstName, user.LastName, user.Email)
	var id int
	if err := row.Scan(&id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.ID = id

	c.JSON(http.StatusCreated, user)
}

func UpdateUser(c *gin.Context) {
	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")

	db := db.GetDB()

	result, err := db.Exec("UPDATE users SET first_name=$1, last_name=$2, email=$3 WHERE id=$4", user.FirstName, user.LastName, user.Email, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
}

// api/routes/user.go

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	// Check if user exists
	var user User
	err := db.GetDB().QueryRow("SELECT id, first_name, last_name, email FROM users WHERE id=$1", id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	// Delete user
	err = db.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
	})
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
