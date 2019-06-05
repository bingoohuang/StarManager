package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Star struct {
	Name        string     `gorm:"primary_key" json:"name"`
	Description string     `json:"description"`
	URL         string     `json:"url"`
	Day         *time.Time `json:"day"`
}

func main() {
	a := &App{}
	 a.Initialize("sqlite3", "test.db")
	//a.Initialize("mysql", "root:111111@tcp(192.168.136.90:3307)/typhon")
	defer a.db.Close()

	r := mux.NewRouter()
	http.Handle("/", r)

	r.HandleFunc("/db/stats", a.DBStatsHandler).Methods("GET")
	r.HandleFunc("/stars", a.ListHandler).Methods("GET")
	r.HandleFunc("/stars/{name:.+}", a.ViewHandler).Methods("GET")
	r.HandleFunc("/stars", a.CreateHandler).Methods("POST")
	r.HandleFunc("/stars/{name:.+}", a.UpdateHandler).Methods("PUT")
	r.HandleFunc("/stars/{name:.+}", a.DeleteHandler).Methods("DELETE")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
