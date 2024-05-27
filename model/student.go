package model

import (
	"github.com/google/uuid"
	"time"
)

// Student represents the student entity
type Student struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Class     string    `json:"class"`
	Subject   string    `json:"subject"`
	Deleted   bool      `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

// GenerateUUID generates a new UUID for student ID
func GenerateUUID() string {
	return uuid.New().String()
}
