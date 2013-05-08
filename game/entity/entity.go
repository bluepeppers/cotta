package entity


import (
	"github.com/bluepeppers/allegro"

	"github.com/bluepeppers/danckelmann/resources"
)

type Entity interface {
	GetSprites(*resources.ResourceManager) []*allegro.Bitmap
}