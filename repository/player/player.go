package player

import (
	"context"
	"fmt"
	"follocoball/store"

	"gorm.io/gorm"
)

type Repository struct {
	Session *gorm.DB
}

func New(session *gorm.DB) *Repository {
	return &Repository{
		Session: session,
	}
}

func (r *Repository) GetPlayers(ctx context.Context) ([]store.Player, error) {
	var players []store.Player
	result := r.Session.Find(&players)

	if result.Error != nil {
		return nil, fmt.Errorf("retrieving all Players failed: %w", result.Error)
	}

	return players, nil
}
