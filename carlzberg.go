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
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *sql.DB
var _viper *viper.Viper

func init() {
	log.Print("init of main")
	_viper := viper.GetViper()
	_viper.SetDefault("dbPath", "data/goshort.db")
	_viper.SetDefault("port", 8080)
	_viper.SetDefault("templatePath", "templates")

	if !_viper.IsSet("dbPath") {
		log.Fatal("No database path (dbPath) is configured.")
	}

	if !_viper.IsSet("templatePath") {
		log.Fatal("No template path (templatePath) is configured.")
	}

	log.Print("viper initialized")
}

func main() {
	log.Print("starting")

	var err error
	// db, err = sql.Open("sqlite3", viper.GetString("dbPath"))
	db, err := gorm.Open(sqlite.Open(viper.GetString("dbPath")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Post{})

	// defer func() {
	// 	_ = db.Close()
	// }()

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

	var post Post
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
