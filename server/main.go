package handler // Change package from main to handler

import (
	"net/http"

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

var router *mux.Router
var ctx app_context.AppContext
var finalHandler http.Handler

func init() {
	ctx = app_context.NewAppContext()
	router = mux.NewRouter()
	router.Use(CorsMiddleware)

	router.HandleFunc("/api/decks", routes.GetDeckTitles(&ctx)).Methods("GET")
	router.HandleFunc("/api/deck/{id}", routes.GetDeck(&ctx)).Methods("GET")
	router.HandleFunc("/api/deck/{id}", routes.EditDeck(&ctx)).Methods("POST")
	router.HandleFunc("/api/deck/", routes.AddDeck(&ctx)).Methods("POST")
	router.HandleFunc("/api/deckcount", routes.GetDeckLength(&ctx)).Methods("GET")
	router.HandleFunc("/api/user", routes.GetProfileData(&ctx)).Methods("GET")
	router.HandleFunc("/api/occupations", routes.GetOccupations(&ctx)).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	finalHandler = c.Handler(router)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	finalHandler.ServeHTTP(w, r)
}