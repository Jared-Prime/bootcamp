package exercise

// a Player is a composition of User with a GameId
type Player struct {
	*User
	GameId int
}
