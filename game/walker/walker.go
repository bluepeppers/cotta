package walker

import (
	"github.com/bluepeppers/allegro"

	"github.com/bluepeppers/danckelmann/resources"
)

type Resource struct {
	Name string
}

type Walker interface {
	GetResource() Resource
	GetSprites(*resources.ResourceManager) []*allegro.Bitmap
}