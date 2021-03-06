package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	IsDone      bool      `json:"is_done"`
	UserId      int       `json:"user_id"`
}
