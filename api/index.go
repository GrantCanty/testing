package handler

import (
	"net/http"
	"strings"

	ac "github.com/GrantCanty/testing/appContext"
	"github.com/GrantCanty/testing/types"
	"github.com/GrantCanty/testing/routes"
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

var ctx ac.AppContext

func init() {
	ctx = app_context.NewAppContext()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Extract the path and remove the /api prefix
	path := strings.TrimPrefix(r.URL.Path, "/api/")
	
	// Handle each route manually
	switch {
	case path == "decks" && r.Method == "GET":
		GetDeckTitles(&ctx)(w, r)
	case strings.HasPrefix(path, "deck/") && r.Method == "GET":
		GetDeck(&ctx)(w, r)
	case strings.HasPrefix(path, "deck/") && r.Method == "POST":
		EditDeck(&ctx)(w, r)
	case path == "deck/" && r.Method == "POST":
		AddDeck(&ctx)(w, r)
	case path == "deckcount" && r.Method == "GET":
		GetDeckLength(&ctx)(w, r)
	case path == "user" && r.Method == "GET":
		GetProfileData(&ctx)(w, r)
	case path == "occupations" && r.Method == "GET":
		GetOccupations(&ctx)(w, r)
	default:
		http.NotFound(w, r)
		return
	}
}