package routes

import (
	"encoding/json"
	"net/http"

	app_context "github.com/GrantCanty/flashcards/appContext"
)

func GetDeckTitles(ac *app_context.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")

		deckNames := make(map[int]string)
		for name := range ac.Decks {
			deckNames[name.ID] = name.Title
		}

		json.NewEncoder(w).Encode(deckNames)
	}
}
