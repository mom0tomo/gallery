package gallery

import (
	"image"
	"time"

	_"image/jpeg"
)

type Image struct {
	ID        int64     `datastore:"-"`
	Image image.RGBA
	CreatedAt time.Time `datastore:"createdAt"`
}