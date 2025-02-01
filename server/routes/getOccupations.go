package routes

import (
	"encoding/json"
	"net/http"

	app_context "github.com/GrantCanty/flashcards/appContext"
)

func GetOccupations(ac *app_context.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")

		json.NewEncoder(w).Encode(ac.Occupations)
	}
}
