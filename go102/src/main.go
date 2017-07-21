package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"os"
	"time"
)

type Config struct {
	Address string
	Static  string
	DBHost  string
	DBName  string
	LogFile string
}

func init() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}

	decorder := json.NewDecoder(file)
	err = decorder.Decode(&config)

	if err != nil {
		log.Fatal(err)
	}

	logfile, err := os.OpenFile(config.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal(err)
	}

	logger = log.New(logfile, "[cms-go]\tINFO\t", log.Ldate|log.Ltime|log.Lshortfile)

	dbs := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "root", "root", config.DBHost, config.DBName)
	log.Printf("Connecting database --- %s/%s\n", config.DBHost, config.DBName)

	DB, err = sql.Open("mysql", dbs)

	if err != nil {
		log.Fatal(err)
	}

}

var DB *sql.DB
var logger *log.Logger
var config Config = Config{}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/feeds", feedList).Methods("GET")
	r.HandleFunc("/api/feeds/{id:[0-9]+}", feedDetail).Methods("GET")
	r.HandleFunc("/api/feeds/{id:[0-9]+}/attachments", feedAttachments).Methods("GET")
	r.HandleFunc("/api/upload", upload).Methods("POST")
	n := negroni.New()
	n.Use(negroni.NewStatic(http.Dir(config.Static)))
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())
	n.UseHandler(r)

	s := &http.Server{
		Handler:      n,
		Addr:         config.Address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting server...")
	log.Fatal(s.ListenAndServe())
}
