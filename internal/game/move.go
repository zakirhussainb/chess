package game

type Move struct {
	Player Player
	Piece  PIECE
	From   string
	To     string
}

func NewMove(player Player, piece PIECE, from, to string) Move {
	return Move{
		Player: player,
		Piece:  piece,
		From:   from,
		To:     to,
	}
}

func validateMove(move Move) bool {
	switch move.Piece {
	case PAWN:
		return validatePawnMove(move.From, move.To, move.Player.Color)
	// case KNIGHT:
	// 	return validateKnightMove(move.From, move.To, move.Player.Color)
	// case BISHOP:
	// 	return validateBishopMove(move.From, move.To, move.Player.Color)
	// case ROOK:
	// 	return validateRookMove(move.From, move.To, move.Player.Color)
	// case QUEEN:
	// 	return validateQueenMove(move.From, move.To, move.Player.Color)
	// case KING:
	// 	return validateKingMove(move.From, move.To, move.Player.Color)
	default:
		return false
	}
}

func validatePawnMove(from, to string, color Color) bool {
	// a2 -> a3
	fromFile := from[0] // a
	fromRank := from[1] // 2
	toFile := to[0]     // a
	toRank := to[1]     // 3

	// When pawns move forward one square
	if fromFile == toFile {
		if color == White {
			return toRank == fromRank+1
		} else {
			return toRank == fromRank-1
		}
	}

	// When pawns move forward two squares from their starting position
	if fromFile == toFile {
		if color == White {
			return fromRank == '2' && toRank == '4'
		} else {
			return fromRank == '7' && toRank == '5'
		}
	}

	// When pawns capture a piece diagonally
	if fromFile == toFile+1 || fromFile == toFile-1 {
		if color == White {
			return toRank == fromRank+1
		} else {
			return toRank == fromRank-1
		}
	}
	return false
}

// func validateKnightMove(from, to string, color Color) bool {
// 	// g1 -> f3
// 	// g1 -> h3
// 	// g1 -> e2
// 	fromFile := from[0] // a
// 	fromRank := from[1] // a
// 	toFile := to[0]     // a
// 	toRank := to[1]     // a

// 	if
// }
