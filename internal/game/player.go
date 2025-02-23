package game

type Player struct {
	ID       string
	Color    Color
	UserName string
}

func NewPlayer(id string, color Color, username string) Player {
	return Player{
		ID:       id,
		Color:    color,
		UserName: username,
	}
}
