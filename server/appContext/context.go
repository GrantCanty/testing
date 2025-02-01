package app_context

import (
	"github.com/GrantCanty/flashcards/types.go"
)

type AppContext struct {
	Decks       map[types.DeckMetaData][]types.Card
	User        types.User
	DeckCounter types.Counter
	CardCounter types.Counter
	Occupations []string
}

func NewAppContext() AppContext {
	c := AppContext{}
	c.DeckCounter.Init()
	c.CardCounter.Init()

	c.Decks = make(map[types.DeckMetaData][]types.Card)
	c.Decks[types.DeckMetaData{ID: c.DeckCounter.New(), Title: "French 3250"}] = []types.Card{{Topic: "Avoir", Description: "To have", ID: c.CardCounter.New()}, {Topic: "Être", Description: "To be", ID: c.CardCounter.New()}, {Topic: "Dire", Description: "To say", ID: c.CardCounter.New()}, {Topic: "Lire", Description: "To read", ID: c.CardCounter.New()}}
	c.Decks[types.DeckMetaData{ID: c.DeckCounter.New(), Title: "CS 2400"}] = []types.Card{{Topic: "+", Description: "Addition", ID: c.CardCounter.New()}, {Topic: "==", Description: "Comparison", ID: c.CardCounter.New()}, {Topic: "-", Description: "Subtraction", ID: c.CardCounter.New()}, {Topic: "float", Description: "number with ability to have decimals", ID: c.CardCounter.New()}}
	c.Decks[types.DeckMetaData{ID: c.DeckCounter.New(), Title: "Econ 3200"}] = []types.Card{{Topic: "Supply", Description: "Amount of a product or service", ID: c.CardCounter.New()}, {Topic: "Demand", Description: "Willingness and ability to buy a product or service", ID: c.CardCounter.New()}, {Topic: "Inflation", Description: "The rate at which prices change", ID: c.CardCounter.New()}, {Topic: "GDP", Description: "Gross Domestic Product. Economic output or a country", ID: c.CardCounter.New()}}
	c.Decks[types.DeckMetaData{ID: c.DeckCounter.New(), Title: "CS 2401"}] = []types.Card{{Topic: "Class", Description: "Data type that holds variables and has member functions. Vars and functions can be public or private", ID: c.CardCounter.New()}, {Topic: "Linked List", Description: "Data type that uses nodes to hold data", ID: c.CardCounter.New()}, {Topic: "Operator Overloading", Description: "Gives special meaning to operators", ID: c.CardCounter.New()}, {Topic: "Inheritance", Description: "Defining classes in a hierarchy", ID: c.CardCounter.New()}}

	c.User.Name = "Grant Canty"
	c.User.About = "I enjoy coding, snowboarding, and playing N64 in my free time"
	c.User.Occupation = "Working"
	c.User.CreateImage()

	c.Occupations = []string{"Student", "Employed", "Retired"}

	return c
}
