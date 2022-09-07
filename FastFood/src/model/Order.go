package model

type Order struct {
	Id        string
	Command   int
	StartDate int
	Item      []Item
}