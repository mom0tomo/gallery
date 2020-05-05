package gallery

import (
	"image"
	"time"

	_ "image/jpeg"
)

// Message is a structure of messages to be posted to guest book.
type Image struct {
	ID        int64     `datastore:"-"`
	Image      image.RGBA    `datastore:"image"`
	CreatedAt time.Time `datastore:"createdAt"`
}