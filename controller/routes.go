package controller

import (
	"follocoball/repository/player"
	"time"

	"github.com/rs/cors"

	"github.com/go-fuego/fuego"
	"gorm.io/gorm"
)

type Resource struct {
	playerQuries *player.Repository
	Session      *gorm.DB
	ExternalAPI  interface{}            // External API
	Cache        map[string]interface{} // Some cache
	Now          func() time.Time       // Function to get the current time. Mocked in tests.
	Security     fuego.Security         // Security configuration
}

func (rs Resource) MountRoutes(s *fuego.Server) {
	fuego.Use(s, cors.AllowAll().Handler)

	playerResource{
		PlayerRepository: rs.playerQuries,
		Session:          rs.Session,
	}.MountRoutes(s)
}
