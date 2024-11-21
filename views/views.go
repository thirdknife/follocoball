package views

import (
	"github.com/go-fuego/fuego"
)

func (rs PlayerResource) MountRoutes(s *fuego.Server) {
	// Public Pages
	fuego.Get(s, "/", homePage)
	fuego.Get(s, "/player", rs.playerPage)
	fuego.Get(s, "/player/add", rs.addPlayerPage)
	fuego.Post(s, "/player/add", rs.createPlayer)
}
