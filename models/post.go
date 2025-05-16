package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Post struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Title     string         `gorm:"not null" json:"title"`
	Slug      string         `gorm:"uniqueIndex;not null" json:"slug"`
	Content   string         `gorm:"type:text" json:"content"`
	Excerpt   string         `gorm:"type:text" json:"excerpt"`
	Published bool           `gorm:"default:false" json:"published"`
	Tags      pq.StringArray `gorm:"type:text[]" json:"tags"`
	PostType  string         `gorm:"not null;default:'post'" json:"post_type"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
