package models

type Game struct {
	SoloGame bool
	IsDealer bool
	Player1  Player
	Player2  Player
	Deck     []Card
}

type Card struct {
	Value int
	Suit  string
}

type Player struct {
	Name   string
	Hand   []Card
	Points int
}
