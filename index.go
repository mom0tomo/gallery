package gallery

import (
	"bytes"
	//"context"
	//"fmt"

	"encoding/base64"
	"html/template"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"

	//"cloud.google.com/go/datastore"
)

var indexTmpl = template.Must(template.ParseFiles("./view/index.html"))


// IndexTemplate is a structure of index template.
type IndexTemplate struct {
	Images    []*Image
}

func Index(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./images/sample.jpg")
	defer file.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	img, _, err := image.Decode(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeImageWithTemplate(w, "index", &img)
}

func writeImageWithTemplate(w http.ResponseWriter, tmpl string, img *image.Image) {
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Fatalln("Unable to encode image.")
	}
	str := base64.StdEncoding.EncodeToString(buffer.Bytes())
	data := map[string]interface{}{"Title": tmpl, "Image": str}
	if err :=indexTmpl.ExecuteTemplate(w, tmpl+".html", data); err != nil {
		log.Fatalln("Unable to execute template.")
	}
}