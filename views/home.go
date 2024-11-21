package views

import (
	"follocoball/templa"

	"github.com/go-fuego/fuego"
)

func homePage(c fuego.ContextNoBody) (fuego.Templ, error) {

	return templa.Home(templa.HomeProps{}), nil
}
