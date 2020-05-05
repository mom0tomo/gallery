package gallery

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"cloud.google.com/go/datastore"
)

func Post(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, "wwgt-codelabs")
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	images := r.FormValue("images")
	if images == "" {
		images = "NO MESSAGE"
	}

	img := &Image{
		Image:   image,
		CreatedAt: time.Now(),
	}

	key := datastore.IncompleteKey(r.Host, nil)
	if _, err := client.Put(ctx, key, img); err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}