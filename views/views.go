package views

import (
	"github.com/go-fuego/fuego"
)

func (rs PlayerResource) MountRoutes(s *fuego.Server) {
	// Public Pages
	fuego.Get(s, "/", homePage)
	fuego.Get(s, "/player", rs.playerPage)
	fuego.Get(s, "/player/edit/{id}", rs.editPlayer)
	fuego.Get(s, "/player/add", rs.addPlayerPage)
	fuego.Post(s, "/player/add", rs.createPlayer)
	fuego.Delete(s, "/player/delete/{id}", rs.deletePlayer)
	fuego.Put(s, "/player/edit/{id}", rs.updatePlayer)
}
