package views

import (
	"github.com/go-fuego/fuego"
)

type Resource struct {
}

func (rs Resource) MountRoutes(s *fuego.Server) {
	// Public Pages
	fuego.Get(s, "/", homePage)
	fuego.Get(s, "/player", playerPage)
	fuego.Get(s, "/player/add", addPlayerPage)
	fuego.Post(s, "/player/add", createPlayer)
}
