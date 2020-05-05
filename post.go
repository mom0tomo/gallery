package gallery

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	_ "image/jpeg"

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

	file, _, err := r.FormFile("upload")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	f, err := os.Create("/tmp/test.jpg")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	io.Copy(f, file)
	http.Redirect(w, r, "/index", http.StatusFound)

	key := datastore.IncompleteKey(r.Host, nil)
	if _, err := client.Put(ctx, key, file); err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}