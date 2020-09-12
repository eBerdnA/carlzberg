package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

var db *sql.DB

func main() {
	viper.SetDefault("dbPath", "data/goshort.db")
	viper.SetDefault("port", 8080)

	if !viper.IsSet("dbPath") {
		log.Fatal("No database path (dbPath) is configured.")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", CatchAllHandler)
	r.HandleFunc("/default", DefaultHandler).Methods(http.MethodGet)

	addr := ":" + strconv.Itoa(viper.GetInt("port"))
	fmt.Println("Listening to " + addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

func CatchAllHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/default", http.StatusTemporaryRedirect)
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	type row struct {
		Slug string
		Url  string
		Hits int
	}
	// var list []row
	// entry := row{Slug: "demo", Url: "http://heise.de",Hits: 0}
	// list = append(list, entry)
	defaultTemplate.Execute(w, nil)

}
