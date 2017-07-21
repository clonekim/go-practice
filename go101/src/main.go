package main

import (
	"bytes"
	"encoding/json"
	"github.com/disintegration/imaging"
	"github.com/gorilla/mux"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

const (
	PORT         = ":8000"
	GALLERY_PATH = "static/photos"
)

type Gallery struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Filename string `json:"filename"`
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	file := vars["name"]
	fullpath := filepath.Join("static", file)

	log.Println("static file -> " + fullpath)
	_, err := os.Stat(fullpath)

	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	http.ServeFile(w, r, fullpath)
}

func thumbnailHandler(w http.ResponseWriter, r *http.Request) {

	file := mux.Vars(r)["name"]
	fullpath := filepath.Join(GALLERY_PATH, file)
	size, err1 := strconv.Atoi(r.URL.Query().Get("size"))

	if err1 != nil {
		size = 100
	}

	if size > 800 {
		size = 800
	}

	_, err := os.Stat(fullpath)

	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	log.Println("generating thumbnail....")

	src, _ := imaging.Open(fullpath)
	src = imaging.Resize(src, size, 0, imaging.Lanczos)

	buf := new(bytes.Buffer)
	jpeg.Encode(buf, src, nil)

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buf.Bytes())))
	if _, err := w.Write(buf.Bytes()); err != nil {
		log.Println("unable to write image")
	}

	defer func() {
		buf = nil
		src = nil
	}()

}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	raw, err := ioutil.ReadFile("./photo.json")

	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		log.Println("Cannot find photo.json")
		return
	}

	var response []Gallery
	json.Unmarshal(raw, &response)
	json.NewEncoder(w).Encode(response)

}

func main() {

	log.Println("starting server on port " + PORT)
	router := mux.NewRouter()
	router.HandleFunc("/static/{name:.+}", staticHandler)
	router.HandleFunc("/thumbnails/{name:.+}", thumbnailHandler)
	router.HandleFunc("/api", pageHandler)
	router.HandleFunc("/", staticHandler)

	http.Handle("/", router)
	http.ListenAndServe(PORT, nil)
}
