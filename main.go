package main

import (
	"os"

	"github.com/bluepeppers/cotta/config"
	"github.com/bluepeppers/cotta/display"
	"github.com/bluepeppers/cotta/game"
	"github.com/bluepeppers/cotta/data"
)

func main() {
	os.Chdir(data.GetDataDir())

	display.InitalizeAllegro()
	conf := config.LoadUserConfig()
	window := display.CreateDisplay(conf)
	defer window.Destroy()
	g := game.CreateGame(conf)

	g.MainLoop()
}
