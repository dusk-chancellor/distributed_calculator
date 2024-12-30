package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	Email    string    `json:"email" db:"email"`
	Password []byte    `json:"password" db:"password"`
	Role     string    `json:"role" db:"role"`
}
