package controller

import (
	"context"
	"follocoball/repository/player"
	"follocoball/store"

	"github.com/go-fuego/fuego"
	"gorm.io/gorm"
)

type playerResource struct {
	PlayerRepository *player.Repository
	Session          *gorm.DB
}

func (ps playerResource) MountRoutes(s *fuego.Server) {
	ps.PlayerRepository = player.New(ps.Session)
	PlayerGroup := fuego.Group(s, "/players")

	fuego.Get(PlayerGroup, "/players", ps.getAllPlayers).
		Summary("Get all Players").Description("Get all players")
}

func (ps playerResource) getAllPlayers(c fuego.ContextNoBody) ([]store.Player, error) {
	players, err := ps.PlayerRepository.GetPlayers(c.Context())
	if err != nil {
		return nil, err
	}

	return players, nil
}

type PlayerRepository interface {
	GetPlayers(ctx context.Context) ([]store.Player, error)
}
