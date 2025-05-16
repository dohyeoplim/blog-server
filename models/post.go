package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Post struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title     string         `gorm:"not null"`
	Slug      string         `gorm:"uniqueIndex;not null"`
	Content   string         `gorm:"type:text"`
	Excerpt   string         `gorm:"type:text"`
	Published bool           `gorm:"default:false"`
	Tags      pq.StringArray `gorm:"type:text[]"`
	PostType  string         `gorm:"not null;default:'post'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
