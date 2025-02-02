package main

import (
	"net/http"
	"strings"

	ac "github.com/GrantCanty/testing/api/appContext"
	"github.com/GrantCanty/testing/api/routes"
)

var ctx ac.AppContext

func init() {
	ctx = ac.NewAppContext()
}

func main() {
	http.HandleFunc("/api/decks", handleDecks)
	http.HandleFunc("/api/deck/", handleDeck)
	http.HandleFunc("/api/deckcount", handleDeckCount)
	http.HandleFunc("/api/user", handleUser)
	http.HandleFunc("/api/occupations", handleOccupations)

	http.ListenAndServe(":3000", nil)
}

func handleDecks(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/")

	if path == "decks" && r.Method == "GET" {
		routes.GetDeckTitles(&ctx)(w, r)
	}
}

func handleDeck(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/")
	
	if strings.HasPrefix(path, "deck/") && r.Method == "GET" {
		routes.GetDeck(&ctx)(w, r)
	} else if strings.HasPrefix(path, "deck/") && r.Method == "POST" {
		routes.EditDeck(&ctx)(w, r)
	} else if path == "deck/" && r.Method == "POST" {
		routes.AddDeck(&ctx)(w, r)
	}
}

func handleDeckCount(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		routes.GetDeckLength(&ctx)(w, r)
	}
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		routes.GetProfileData(&ctx)(w, r)
	}
}

func handleOccupations(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		routes.GetOccupations(&ctx)(w, r)
	}
}
