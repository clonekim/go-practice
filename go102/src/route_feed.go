package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func feedList(writer http.ResponseWriter, request *http.Request) {
	feeds, err := Feeds()

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInsufficientStorage)
		return
	}

	ToJSON(writer, feeds)
}

func feedDetail(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInsufficientStorage)
		return
	}

	feed, err := SelectFeedById(id)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInsufficientStorage)
		return
	}

	ToJSON(writer, feed)
}
