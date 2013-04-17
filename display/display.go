package display

import (
	"log"
	"os"
	"sync"
	"fmt"
)

import "github.com/bluepeppers/allegro"

import "github.com/bluepeppers/cotta/config"

const (
	DEFAULT_WIDTH     = 600
	DEFAULT_HEIGHT    = 400
	DEFAULT_FONT_SIZE = 12
	DEFAULT_TILEMAP   = "tilemap.png"
	// Because I'm lazy; this should be dynamic
	TILE_WIDTH = 8
	TILE_HEIGHT = 8
	TILE_HORIZONTAL = 16
	TILE_VERTICAL = 16
)

type Display struct {
	disp    *allegro.Display
	font    *allegro.Font
	tilemap *TileMap
}

type TileMap struct {
	bmp *allegro.Bitmap
	tileWidth int
	tileHeight int
	subBmpCache [](*allegro.Bitmap)
	glyphPositions map[string][2]int
}

func InitalizeAllegro() {
	allegro.Init()
	allegro.InstallKeyboard()
	allegro.InstallMouse()
	allegro.InitFont()
	allegro.InitImage()
	allegro.InitTTF()
}

func CreateDisplay(conf *allegro.Config) *Display {
	var adisp *allegro.Display
	var afont *allegro.Font
	var atm *TileMap

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		adisp = createDisp(conf)
		wg.Done()
	}()
	go func() {
		afont = loadFont(conf)
		wg.Done()
	}()
	go func() {
		atm = loadTileMap(conf)
		wg.Done()
	}()
	wg.Wait()
	return &Display{adisp, afont, atm}
}

func (d *Display) Destroy() {
	d.disp.Destroy()
	d.font.Destroy()
	d.tilemap.Destroy()
}

func createDisp(conf *allegro.Config) *allegro.Display {
	width := config.GetInt(conf, "display", "width", DEFAULT_WIDTH)
	height := config.GetInt(conf, "display", "height", DEFAULT_HEIGHT)

	flags := allegro.RESIZABLE
	switch config.GetString(conf, "display", "windowed", "windowed") {
	case "fullscreen":
		flags |= allegro.FULLSCREEN
	case "fullscreenwindow":
		flags |= allegro.FULLSCREEN_WINDOW
	default:
		log.Printf("display.windowed not one of \"fullscreen\", \"fullscreenwindow\", or \"windowed\"")
		log.Printf("Defaulting to display.windowed=\"windowed\"")
		fallthrough
	case "windowed":
		flags |= allegro.WINDOWED
	}

	disp := allegro.CreateDisplay(width, height, flags)
	if disp == nil {
		log.Fatalf("Could not create display")
	}
	return disp
}

func loadFont(conf *allegro.Config) *allegro.Font {
	fname := config.GetString(conf, "display", "font", "")
	size := config.GetInt(conf, "display", "fontsize", DEFAULT_FONT_SIZE)
	if fname != "" {
		f := allegro.LoadFont(fname, size, 0)
		if f != nil {
			return f
		}
	}
	return allegro.CreateBuiltinFont()
}

func exists(fname string) bool {
	_, err := os.Open(fname)
	return !os.IsNotExist(err)
}

func loadTileMap(conf *allegro.Config) *TileMap {
	fname := config.GetString(conf, "display", "tilemap", DEFAULT_TILEMAP)
	if fname != DEFAULT_TILEMAP && !exists(fname) {
		log.Printf("display.tilemap=%q not a file", fname)
		log.Printf("Defaulting to display.tilemap=%q", DEFAULT_TILEMAP)
		fname = DEFAULT_TILEMAP
	}
	if !exists(fname) {
		log.Fatalf("display.tilemap=%q is not a file", fname)
	}
	
	bmp := allegro.LoadBitmap(fname)
	if bmp == nil {
		log.Fatalf("display.tilemap=%q could not be loaded as a bitmap")
	}
	
	tm := new(TileMap)
	tm.bmp = bmp
	tm.subBmpCache = make([]*allegro.Bitmap, TILE_HORIZONTAL * TILE_VERTICAL)
	tm.glyphPositions = getGlyphPositions()
	tm.tileWidth = TILE_WIDTH
	tm.tileHeight = TILE_HEIGHT

	minx := tm.tileWidth * TILE_HORIZONTAL
	miny := tm.tileHeight * TILE_VERTICAL
	width, height := tm.bmp.GetDimensions()
	if minx > width || miny > height {
		log.Fatalf("display.tilemap=%q smaller than minimum (%v, %v) < (%v, %v)",
			minx, miny, width, height)
	}
	
	return tm
}

func (tm *TileMap) Destroy() {
	for _, sub := range tm.subBmpCache {
		if sub != nil {
			sub.Destroy()
		}
	}
	tm.bmp.Destroy()
}

func (tm *TileMap) GetTile(x, y int) *allegro.Bitmap {
	if x < 0 || x > TILE_HORIZONTAL ||
		y < 0 || y > TILE_VERTICAL {
		panic(fmt.Sprintf("Can't access tile out of range: (%v, %v)", x, y))
	}

	p := x * TILE_HORIZONTAL + y
	if tm.subBmpCache[p] == nil {
		xp := x * tm.tileWidth
		yp := y * tm.tileHeight
		sub := tm.bmp.CreateSubBitmap(xp, yp, tm.tileWidth, tm.tileHeight)
		if sub == nil {
			panic(fmt.Sprintf("Can't create subbitmap (%v, %v, %v, %v)", 
				xp, yp, tm.tileWidth, tm.tileHeight))
		}
		tm.subBmpCache[p] = sub
	}
	
	return tm.subBmpCache[p]
}

func (tm *TileMap) GetGlyph(code string) *allegro.Bitmap {
	pos, ok := tm.glyphPositions[code]
	if !ok {
		panic(fmt.Sprintf("Can't access glyph not in tilesheet: %v", code))
	}
	
	return tm.GetTile(pos[0], pos[1])
}