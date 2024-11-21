package views

import (
	"follocoball/store"
	"follocoball/templa"
	"net/http"

	"github.com/go-fuego/fuego"
)

func playerPage(c fuego.ContextNoBody) (fuego.Templ, error) {
	return templa.Player(), nil
}

func addPlayerPage(c fuego.ContextNoBody) (fuego.Templ, error) {
	return templa.PlayerNew(), nil
}

func createPlayer(c *fuego.ContextWithBody[store.Player]) (any, error) {
	c.Response().Header().Set("HX-Trigger", "entity-updated")
	return c.Redirect(http.StatusSeeOther, "/player")
}
