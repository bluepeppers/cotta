package main

import (
	"log"
	"os"
	"runtime"

	"github.com/bluepeppers/cotta/data"
	"github.com/bluepeppers/cotta/game"
	"github.com/bluepeppers/danckelmann/config"
	"github.com/bluepeppers/danckelmann/display"
)

const (
	CONFIG_LOCATION = "${HOME}/.cottaconfig"
)

func main() {
	runtime.GOMAXPROCS(10)

	dataDir := data.GetDataDir()
	os.Chdir(dataDir)

	display.InitializeAllegro()
	conf := config.LoadUserConfig(CONFIG_LOCATION)
	if conf == nil {
		log.Panicf("Could not load user config (%q) or create empty", CONFIG_LOCATION)
	}

	defaultResourceDir := "resources"
	resourceDir := config.GetString(conf, "resources", "dir", defaultResourceDir)

	ge := game.CreateGameEngine(conf)
	de := display.CreateDisplayEngine(resourceDir, conf, ge)

	go ge.MainLoop()
	de.Run()
}
