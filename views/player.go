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

func (rs PlayerResource) editPlayer(c fuego.ContextNoBody) (fuego.Templ, error) {
	id := c.PathParam("id")

	var player store.Player
	result := rs.Session.Where("id = ?", id).Find(&player)

	if result.Error != nil {
		return nil, fmt.Errorf("retrieving Player failed: %w", result.Error)
	}

	return templa.EditPlayer(templa.EditPlayerProps{Player: player}), nil
}

func (rs PlayerResource) updatePlayer(c *fuego.ContextWithBody[store.Player]) (any, error) {
	updatePlayerArgs, _ := c.Body()
	fmt.Print(updatePlayerArgs.Name)

	player := store.Player{
		Name: updatePlayerArgs.Name,
	}

	_ = rs.Session.Save(&player) // pass pointer of data to Create

	return c.Redirect(http.StatusMovedPermanently, "/player")
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

	_ = rs.Session.Create(&player) // pass pointer of data to Create

	c.Response().Header().Set("HX-Trigger", "entity-updated")
	return c.Redirect(http.StatusSeeOther, "/player")
}

func (rs PlayerResource) deletePlayer(c fuego.ContextNoBody) (any, error) {
	id := c.PathParam("id")

	var player store.Player
	result := rs.Session.Where("id = ?", id).Delete(&player)
	fmt.Print(result.RowsAffected)

	if result.Error != nil {
		return nil, fmt.Errorf("deleting Player failed: %w", result.Error)
	}

	return rs.playerPage(c)
}
