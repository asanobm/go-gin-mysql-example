package repo

import "github.com/google/uuid"

type Todo struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Title       string    `json:"title" gorm:"type:varchar(100) not null"`
	Description string    `json:"description" gorm:"type:blob not null"`
}
