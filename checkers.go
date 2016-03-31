package main

import (
	"fmt"
	// "strconv"
)

type Checkers struct {
	whitePlayerTurn bool
	grid            [8][8]Piece
}

func fillSpot(row int, col int, checkers *Checkers) {
	if row == 0 || row == 2 {
		if col%2 == 1 {
			checkers.grid[row][col].mask = 0
		} else {
			checkers.grid[row][col].mask = 1
		}
	} else if row == 6 {
		if col%2 == 1 {
			checkers.grid[row][col].mask = 0
		} else {
			checkers.grid[row][col].mask = 2
		}
	} else if row == 1 {
		if col%2 == 0 {
			checkers.grid[row][col].mask = 0
		} else {
			checkers.grid[row][col].mask = 1
		}
	} else if row == 5 || row == 7 {
		if col%2 == 0 {
			checkers.grid[row][col].mask = 0
		} else {
			checkers.grid[row][col].mask = 2
		}
	} else {
		checkers.grid[row][col].mask = 0
	}
}

func initCheckers() Checkers {
	var checkers Checkers

	checkers.whitePlayerTurn = true

	for row := 0; row < 8; row++ {

		for col := 0; col < 8; col++ {

			fillSpot(row, col, &checkers)
		}
	}

	return checkers
}

func displayCheckers(checkers Checkers) {
	fmt.Println("    A   B   C   D   E   F   G   H")
	fmt.Println("  +---+---+---+---+---+---+---+---+")

	for row := 0; row < 8; row++ {
		fmt.Print(row + 1)
		fmt.Print(" ")

		for col := 0; col < 8; col++ {
			fmt.Print("| " + getPieceLetter(checkers.grid[row][col]) + " ")
		}

		fmt.Println("|")
		fmt.Println("  +---+---+---+---+---+---+---+---+")
	}

}

func getPiece(checkers Checkers, loc Location) Piece {
	return checkers.grid[loc.y][loc.x]
}

func countPieces(checkers Checkers) (int8, int8) {
	var numWhite int8 = 0
	var numRed int8 = 0

	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			if isWhite(checkers.grid[row][col]) {
				numWhite++
			}
			if isRed(checkers.grid[row][col]) {
				numRed++
			}
		}
	}

	return numWhite, numRed
}

func jumperHasAnotherJump(checkers Checkers, loc Location) bool {
	p := getPiece(checkers, loc)

	if isWhite(p) {
		if isKinged(p) {
			//up left
			jmpLoc := Location{loc.x - 1, loc.y - 1}
			jmp := getPiece(checkers, jmpLoc)

			dstLoc := Location{loc.x - 2, loc.y - 2}
			dst := getPiece(checkers, dstLoc)

			if isValid(jmpLoc) && isValid(dstLoc) {
				if isRed(jmp) && isEmpty(dst) {
					return true
				}
			}

			//up right
			jmpLoc = Location{loc.x + 1, loc.y - 1}
			jmp = getPiece(checkers, jmpLoc)
			dstLoc = Location{loc.x + 2, loc.y - 1}
			dst = getPiece(checkers, dstLoc)

			if isValid(jmpLoc) && isValid(dstLoc) {
				if isRed(jmp) && isEmpty(dst) {
					return true
				}
			}
		}

		//down left
		jmpLoc := Location{loc.x - 1, loc.y + 1}
		jmp := getPiece(checkers, jmpLoc)
		dstLoc := Location{loc.x - 2, loc.y + 2}
		dst := getPiece(checkers, dstLoc)

		if isValid(jmpLoc) && isValid(dstLoc) {
			if isRed(jmp) && isEmpty(dst) {
				return true
			}
		}

		//down right
		jmpLoc = Location{loc.x + 1, loc.y + 1}
		jmp = getPiece(checkers, jmpLoc)
		dstLoc = Location{loc.x + 2, loc.y + 2}
		dst = getPiece(checkers, dstLoc)

		if isValid(jmpLoc) && isValid(dstLoc) {
			if isRed(jmp) && isEmpty(dst) {
				return true
			}
		}

		return false
	} else if isRed(p) {
		if isKinged(p) {
			//up left
			jmpLoc := Location{loc.x - 1, loc.y + 1}
			jmp := getPiece(checkers, jmpLoc)
			dstLoc := Location{loc.x - 2, loc.y + 2}
			dst := getPiece(checkers, dstLoc)

			if isValid(jmpLoc) && isValid(dstLoc) {
				if isWhite(jmp) && isEmpty(dst) {
					return true
				}
			}

			//up right
			jmpLoc = Location{loc.x + 1, loc.y + 1}
			jmp = getPiece(checkers, jmpLoc)
			dstLoc = Location{loc.x + 2, loc.y + 2}
			dst = getPiece(checkers, dstLoc)

			if isValid(jmpLoc) && isValid(dstLoc) {
				if isRed(jmp) && isEmpty(dst) {
					return true
				}
			}
		}

		//down left
		jmpLoc := Location{loc.x -1, loc.y - 1}
		jmp := getPiece(checkers, jmpLoc)
		dstLoc := Location{loc.x - 2, loc.y - 2}
		dst := getPiece(checkers, dstLoc)

		if isValid(jmpLoc) && isValid(dstLoc) {
			if isWhite(jmp) && isEmpty(dst) {
				return true
			}
		}

		//down right
		jmpLoc = Location{loc.x + 1, loc.y - 1}
		jmp = getPiece(checkers, jmpLoc)
		dstLoc = Location{loc.x + 2, loc.y - 2}
		dst = getPiece(checkers, dstLoc)

		if isValid(jmpLoc) && isValid(dstLoc) {
			if isWhite(jmp) && isEmpty(dst) {
				return true
			}
		}

		return false
	} else {
		return false
	}
}