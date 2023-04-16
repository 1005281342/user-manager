package db

import (
	"database/sql"
	"fmt"
	"github.com/1005281342/user-manager/models"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect() error {
	var err error

	db, err = sql.Open("postgres", "postgres://postgres:123456@192.168.8.42/user_manager?sslmode=disable")
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	return nil
}

func GetDB() *sql.DB {
	return db
}

func GetAllUsers(perPage int, offset int) ([]models.User, error) {
	var users []models.User

	rows, err := GetDB().Query("SELECT id, first_name, last_name, email FROM users LIMIT $1 OFFSET $2", perPage, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func SearchUsers(keyword string) ([]models.User, error) {
	var users []models.User

	rows, err := GetDB().Query(`SELECT id, first_name, last_name, email FROM users WHERE 
                                first_name LIKE '%' || $1 || '%' OR 
                                last_name LIKE '%' || $1 || '%' OR 
                                email LIKE '%' || $1 || '%'`, keyword)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func SearchUsersPerPage(keyword string, perPage int, page int) ([]models.User, error) {
	var users []models.User

	offset := (page - 1) * perPage

	rows, err := GetDB().Query(`SELECT id, first_name, last_name, email FROM users WHERE
                               first_name LIKE '%' || $1 || '%' OR
                               last_name LIKE '%' || $1 || '%' OR
                               email LIKE '%' || $1 || '%'
                               LIMIT $2 OFFSET $3`, keyword, perPage, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetAllUsersPerPage(perPage int, page int) ([]models.User, error) {
	var users []models.User

	offset := (page - 1) * perPage

	rows, err := GetDB().Query("SELECT id, first_name, last_name, email FROM users LIMIT $1 OFFSET $2", perPage, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	row := GetDB().QueryRow("SELECT id, first_name, last_name, email, password FROM users WHERE email=$1", email)

	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
