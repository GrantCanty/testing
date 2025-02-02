package handler

import (
	"net/http"
	"strings"

	app_context "github.com/GrantCanty/flashcards/appContext"
	"github.com/GrantCanty/flashcards/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json;charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

var ctx app_context.AppContext

func init() {
	ctx = app_context.NewAppContext()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Extract the path and remove the /api prefix
	path := strings.TrimPrefix(r.URL.Path, "/api/")
	
	// Handle each route manually
	switch {
	case path == "decks" && r.Method == "GET":
		routes.GetDeckTitles(&ctx)(w, r)
	case strings.HasPrefix(path, "deck/") && r.Method == "GET":
		routes.GetDeck(&ctx)(w, r)
	case strings.HasPrefix(path, "deck/") && r.Method == "POST":
		routes.EditDeck(&ctx)(w, r)
	case path == "deck/" && r.Method == "POST":
		routes.AddDeck(&ctx)(w, r)
	case path == "deckcount" && r.Method == "GET":
		routes.GetDeckLength(&ctx)(w, r)
	case path == "user" && r.Method == "GET":
		routes.GetProfileData(&ctx)(w, r)
	case path == "occupations" && r.Method == "GET":
		routes.GetOccupations(&ctx)(w, r)
	default:
		http.NotFound(w, r)
		return
	}
}