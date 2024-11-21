package controller

import (
	"follocoball/repository/player"

	"github.com/rs/cors"

	"github.com/go-fuego/fuego"
	"gorm.io/gorm"
)

type Resource struct {
	playerQuries *player.Repository
	Session      *gorm.DB
}

func (rs Resource) MountRoutes(s *fuego.Server) {
	fuego.Use(s, cors.AllowAll().Handler)

	playerResource{
		PlayerRepository: rs.playerQuries,
		Session:          rs.Session,
	}.MountRoutes(s)
}
