package main

type Piece struct {
	mask byte
}

func isRed(p Piece) bool {
	return (p.mask & 2) == 2
}

func isWhite(p Piece) bool {
	return (p.mask & 1) == 1
}

func isEmpty(p Piece) bool {
	return p.mask == 0
}

func isKinged(p Piece) bool {
	return (p.mask & 4) == 4
}

func setKinged(p Piece) Piece {
	p.mask |= 4

	return p
}

func getPieceLetter(p Piece) string {
	if isEmpty(p) {
		return " "
	}
	if isKinged(p) {
		if isWhite(p) {
			return "W"
		} else if isRed(p) {
			return "R"
		} else {
			return "!"
		}
	} else {
		if isWhite(p) {
			return "w"
		} else if isRed(p) {
			return "r"
		} else {
			return "$"
		}
	}
}