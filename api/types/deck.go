package types

type Card struct {
	Topic       string `json:"topic"`
	Description string `json:"description"`
	ID          int    `json:"id"`
}
type DeckMetaData struct {
	ID    int
	Title string
}
