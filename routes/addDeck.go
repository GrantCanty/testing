package routes

import (
	"net/http"

	app_context "github.com/GrantCanty/testing"
	"github.com/GrantCanty/testing/types"
)

func AddDeck(ac *app_context.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		deckName := r.URL.Query().Get("deckName")

		var deck []types.Card
		deck = append(deck, types.Card{Topic: "", Description: ""})

		ac.Decks[types.DeckMetaData{ID: ac.DeckCounter.New(), Title: deckName}] = deck
	}
}
