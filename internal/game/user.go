package game

type User struct {
	ID       string
	UserName string
	Color    Color
	IsPlayer bool
	IsViewer bool
}

func NewUser(id string, username string, color Color, isPlayer, isViewer bool) User {
	return User{
		ID:       id,
		UserName: username,
		Color:    color,
		IsPlayer: isPlayer,
		IsViewer: isViewer,
	}
}
