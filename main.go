package main

import (
	"os"

	"github.com/bluepeppers/danckelmann/config"
	"github.com/bluepeppers/danckelmann/display"
	"github.com/bluepeppers/cotta/game"
	"github.com/bluepeppers/cotta/data"
)

const (
	CONFIG_LOCATION = "${HOME}/.cottaconfig"
)

func main() {
	os.Chdir(data.GetDataDir())

	display.InitializeAllegro()
	conf := config.LoadUserConfig(CONFIG_LOCATION)

	ge := game.CreateGameEngine(conf)
	de := display.CreateDisplayEngine(conf, ge)

	go ge.MainLoop()
	de.Run()
}