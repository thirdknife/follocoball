package views

import (
	"github.com/go-fuego/fuego"
)

type Resource struct {
}

func (rs Resource) MountRoutes(s *fuego.Server) {
	// Public Pages
	fuego.Get(s, "/player", playerPage)
}
