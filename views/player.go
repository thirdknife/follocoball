package views

import (
	"fmt"
	"follocoball/store"
	"follocoball/templa"
	"net/http"

	"github.com/go-fuego/fuego"
	"gorm.io/gorm"
)

type PlayerResource struct {
	Session *gorm.DB
}

func (rs PlayerResource) playerPage(c fuego.ContextNoBody) (fuego.Templ, error) {
	var players []store.Player
	result := rs.Session.Find(&players)

	if result.Error != nil {
		return nil, fmt.Errorf("retrieving all Players failed: %w", result.Error)
	}

	return templa.Player(templa.PlayerProps{Players: players}), nil
}

func (rs PlayerResource) addPlayerPage(c fuego.ContextNoBody) (fuego.Templ, error) {
	return templa.PlayerNew(), nil
}

func (rs PlayerResource) createPlayer(c *fuego.ContextWithBody[store.Player]) (any, error) {
	body, _ := c.Body()

	player := store.Player{
		Name: body.Name,
		// Add other required fields here, for example:
		// BaseModel: store.BaseModel{ID: 0, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	result := rs.Session.Create(&player) // pass pointer of data to Create

	fmt.Println(result)

	c.Response().Header().Set("HX-Trigger", "entity-updated")
	return c.Redirect(http.StatusSeeOther, "/player")
}
