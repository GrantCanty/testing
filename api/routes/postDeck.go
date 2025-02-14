package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	app_context "github.com/GrantCanty/testing/api/appContext"
    "github.com/GrantCanty/testing/api/types"
	"github.com/gorilla/mux"
)

func EditDeck(ac *app_context.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		deckName := r.URL.Query().Get("deckName")
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Println(err)
		}

		var deck []types.Card
		_ = json.NewDecoder(r.Body).Decode(&deck)

		if _, found := ac.Decks[types.DeckMetaData{ID: id, Title: deckName}]; found {
			ac.Decks[types.DeckMetaData{ID: id, Title: deckName}] = deck
		}
	}
}
