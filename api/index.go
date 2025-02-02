package handler

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

// Change function signature to return http.HandlerFunc
func Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
}
