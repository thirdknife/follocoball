package store

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	ID uuid.UUID `gorm:"primaryKey;type:uuid"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (bm *BaseModel) BeforeCreate(tx *gorm.DB) error {
	bm.ID = uuid.New()
	return nil
}

type Player struct {
	BaseModel

	// The full name of the player. Keeping it in a single string allows any input, which is better
	// than trying to deal with the intricacies of separating first and last names, nicknames, etc.
	Name string
}
