package exercise

// a Player is a composition of User with a GameId
type Player struct {
	*User
	GameId int
}

func NewPlayer(id int, name string, location string, gameId int) *Player {
	return &Player{User: &User{id, name, location}, GameId: gameId}
}
