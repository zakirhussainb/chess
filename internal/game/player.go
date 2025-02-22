package game

type Player struct {
	ID       string
	Color    string
	UserName string
}

func NewPlayer(id, color, username string) Player {
	return Player{
		ID:       id,
		Color:    color,
		UserName: username,
	}
}
