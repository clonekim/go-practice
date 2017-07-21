package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
	"strconv"
)

func feedAttachments(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	attachments, err := SelectAttachmentsByRefId(id)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	ToJSON(writer, attachments)
}

func upload(writer http.ResponseWriter, request *http.Request) {

	request.ParseMultipartForm(50 << 20)

	file, handler, err := request.FormFile("upload")

	userid := request.FormValue("userid")

	fmt.Printf("Form Data[userid]: %s\n", userid)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	defer file.Close()

	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	defer f.Close()
	size, err := io.Copy(f, file)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	attach := Attachment{
		FileName:    handler.Filename,
		FileSize:    size,
		DiskName:    handler.Filename,
		ContentType: handler.Header.Get("Content-Type"),
		Deleted:     "0",
	}

	ToJSON(writer, attach)
}
