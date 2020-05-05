package gallery

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
)

var indexTmpl = template.Must(template.ParseFiles("./view/index.html"))

// IndexTemplate is a structure of index template.
type IndexTemplate struct {
	Images    []*Image
}

func Index(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, "wwgt-codelabs")
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	var imgs []*Images
	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	keys, err := client.GetAll(ctx, wc, &img)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	for i := 0; i < len(img); i++ {
		img[i].ID = keys[i].ID
	}

	idxt := &IndexTemplate{
		Image:     image,
		Images:    imgs,
	}

	if err := indexTmpl.Execute(w, idxt); err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
	}
}