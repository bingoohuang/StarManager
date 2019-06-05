package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/now"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	db *gorm.DB
}

func (a *App) Initialize(dbDriver, dbURI string) {
	isMySQLDriver := IsMySQLDriver(dbDriver)
	if isMySQLDriver {
		dbURI = FixMySQLURIParameters(dbURI)
	}

	db, err := gorm.Open(dbDriver, dbURI)
	if err != nil {
		panic(err)
	}
	if isMySQLDriver {
		db = FixMySQLTableOptions(db)
	}

	SetConnectionPool(db.DB())

	a.db = db
	// Migrate the schema.
	a.db.AutoMigrate(&Star{})
}

func (a *App) DBStatsHandler(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(a.db.DB().Stats())
}

func (a *App) ListHandler(w http.ResponseWriter, r *http.Request) {
	var stars []Star

	// Select all stars and convert to JSON.
	a.db.Find(&stars)
	starsJSON, _ := json.Marshal(stars)

	// Write to HTTP response.
	w.WriteHeader(200)
	_, _ = w.Write(starsJSON)
}

func (a *App) ViewHandler(w http.ResponseWriter, r *http.Request) {
	var star Star
	vars := mux.Vars(r)

	// Select the star with the given name, and convert to JSON.
	a.db.First(&star, "name = ?", vars["name"])
	starJSON, _ := json.Marshal(star)

	// Write to HTTP response.
	w.WriteHeader(200)
	_, _ = w.Write(starJSON)
}

func (a *App) CreateHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the POST body to populate r.PostForm.
	if err := r.ParseForm(); err != nil {
		panic("failed in ParseForm() call")
	}

	// Create a new star from the request body.
	dayStr := r.PostFormValue("day")
	var day *time.Time
	if dayStr != "" {
		d := now.MustParse(dayStr)
		day = &d
	}
	star := &Star{
		Name:        r.PostFormValue("name"),
		Description: r.PostFormValue("description"),
		URL:         r.PostFormValue("url"),
		Day:         day,
	}
	a.db.Create(star)

	// Form the URL of the newly created star.
	u, err := url.Parse(fmt.Sprintf("/stars/%s", star.Name))
	if err != nil {
		panic("failed to form new Star URL")
	}
	base, err := url.Parse(r.URL.String())
	if err != nil {
		panic("failed to parse request URL")
	}

	// Write to HTTP response.
	w.Header().Set("Location", base.ResolveReference(u).String())
	w.WriteHeader(201)
}

func (a *App) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Parse the POST body to populate r.PostForm.
	if err := r.ParseForm(); err != nil {
		panic("failed in ParseForm() call")
	}

	dayStr := r.PostFormValue("day")
	var day *time.Time
	if dayStr != "" {
		d := now.MustParse(dayStr)
		day = &d
	}
	// Set new star values from the request body.
	star := &Star{
		Name:        r.PostFormValue("name"),
		Description: r.PostFormValue("description"),
		URL:         r.PostFormValue("url"),
		Day:         day,
	}

	// Update the star with the given name.
	a.db.Model(&star).Where("name = ?", vars["name"]).Updates(&star)

	// Write to HTTP response.
	w.WriteHeader(204)
}

func (a *App) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Delete the star with the given name.
	a.db.Where("name = ?", vars["name"]).Delete(Star{})

	// Write to HTTP response.
	w.WriteHeader(204)
}
