package main

import "fmt"

// User type
type User struct {
	ID             int
	Name, Location string
}

// Player type uses composition of User type
type Player struct {
	User
	GameID int
}

func main() {
	p := Player{
		User{ID: 42, Name: "Player1", Location: "LA"},
		90004,
	}

	fmt.Printf("ID: %d, Name: %s, Location: %s, Game id: %d\n",
		p.ID, p.Name, p.Location, p.GameID)

	p.ID = 11

	fmt.Printf("%+v", p)
}
