package models

import "auth-service/db"

type User struct {
	Id        string `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
}

func CreateUser(user User) error {

	db := db.DB
	query := `INSERT INTO users (email, password) 
	VALUES ($1, $2) RETURNING id, created_at`

	var id string
	var createdAt string

	err := db.QueryRow(query, user.Email, user.Password).Scan(&id, &createdAt)
	if err != nil {
		return err
	}
	return nil
}
