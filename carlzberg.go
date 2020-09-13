package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
)

var db *sql.DB

func main() {
	viper.SetDefault("dbPath", "data/goshort.db")
	viper.SetDefault("port", 8080)

	if !viper.IsSet("dbPath") {
		log.Fatal("No database path (dbPath) is configured.")
	}

	var err error
	db, err = sql.Open("sqlite3", viper.GetString("dbPath"))
	if err != nil {
		log.Fatal(err)
	}

	migrateDatabase()

	defer func() {
		_ = db.Close()
	}()

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
	type post_type struct {
		Id           int
		Post_title   string
		Post_content string
	}
	var post post_type
	err := db.QueryRow("SELECT id, post_title, post_content FROM posts WHERE post_title = ?", "Hello World from Carlzberg").Scan(&post.Id, &post.Post_title, &post.Post_content)
	if err != nil {
		http.NotFound(w, r)
		log.Fatal(err)
		return
	}

	err = defaultTemplate.Execute(w, &post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
