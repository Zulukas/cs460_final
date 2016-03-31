package main

import (
	"fmt"
	"strconv"
)

/*
	Helper function to perform a frequently used set of checks
*/
func validateWhiteJump(checkers Checkers, src Location, jmp Location, dst Location) bool {
	if isValid(src) && isValid(jmp) && isValid(dst) {
		if isWhite(getPiece(checkers, src)) && isRed(getPiece(checkers, jmp)) && isEmpty(getPiece(checkers, dst)) {
			return true
		}
	}

	return false
}

/*
	Helper function to check the diagonals at a location to see if
	it is a possible jumping location for a white piece.
*/
func whitePieceHasJump(checkers Checkers, loc Location) bool {
	if !isValid(loc) {
		return false
	}

	//White pieces move upward (Decrementing on the grid)
	if isKinged(checkers.grid[loc.y][loc.x]) {
		//up left
		jmp := Location{loc.x - 1, loc.y - 1}
		dst := Location{loc.x - 2, loc.y - 2}

		if validateWhiteJump(checkers, loc, jmp, dst) { return true }

		//up right
		jmp = Location{loc.x + 1, loc.y - 1}
		dst = Location{loc.x + 2, loc.y - 2}

		if validateWhiteJump(checkers, loc, jmp, dst) { return true }
	}

	//down left
	jmp := Location{loc.x - 1, loc.y + 1}
	dst := Location{loc.x - 2, loc.y + 2}

	if validateWhiteJump(checkers, loc, jmp, dst) { return true }

	//down right
	jmp = Location{loc.x + 1, loc.y + 1}
	dst = Location{loc.x + 2, loc.y + 2}

	if validateWhiteJump(checkers, loc, jmp, dst) { return true }

	return false
}

/*
	Scan the checker board for possible white jumps and return a 
	list of locations from where a player may jump.
*/
func getWhiteForcedJumps(checkers Checkers) []Location {
	var possibleJumps []Location
	//Iterate through all the pieces on the board...
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			if isWhite(checkers.grid[row][col]) {
				src := Location{int8(col), int8(row)}

				if whitePieceHasJump(checkers, src) {
					possibleJumps = append(possibleJumps, src)
				}
			}			
		}
	}

	return possibleJumps
}

/*
	Game logic for handling a white piece move whether they're kinged or not.
*/
func whiteMove(checkers *Checkers, src Location, dst Location) bool {

	//Ensure it is white player's turn
	if checkers.whitePlayerTurn == false {
		fmt.Println("Not white player turn!")
		return false
	}

	//Make sure the src and dst are within legal bounds
	if !isValid(src) || !isValid(dst) {
		fmt.Println("src or dst is invalid!")
		return false
	}

	//Grab the pieces for convenient use
	srcPiece := getPiece(*checkers, src)
	dstPiece := getPiece(*checkers, dst)

	srcIsKinged := isKinged(srcPiece)

	//Make sure the src and dst are legal for this context
	if !isWhite(srcPiece) || !isEmpty(dstPiece) {
		fmt.Println("src is not white or dst is not empty!")
		return false
	}

	//Check the delta...
	xDelta, yDelta := getDelta(src, dst)

	legalX := (abs(xDelta) == 1)
	legalY := (abs(yDelta) == 1)

	//Now check the movement
	if srcIsKinged && (!legalX || !legalY) {
		fmt.Println("Invalid kinged vector!")
		return false
	} else if !srcIsKinged && (!legalX || yDelta != 1) {
		fmt.Println("Invalid non-kinged vector!")
		fmt.Println("\t" + strconv.Itoa(int(xDelta)) + ", " + strconv.Itoa(int(yDelta)))
		return false
	}

	var empty Piece
	empty.mask = 0

	//If it passed the gauntlet, the move is legal
	checkers.grid[dst.y][dst.x] = Piece{1}
	checkers.grid[src.y][src.x] = Piece{0}
	checkers.whitePlayerTurn = false

	if dst.y == 7 {
		checkers.grid[dst.y][dst.x] = setKinged(checkers.grid[dst.y][dst.x])
	}

	fmt.Println("success")

	return true
}

/*
	Game logic to handle white jumps whether they're kinged or not.
*/
func whiteJump(checkers *Checkers, src Location, dst Location) bool {

	//Must be white's turn to make a white jump
	if checkers.whitePlayerTurn == false {
		return false
	}

	//Validate src and dst
	if !isValid(src) || !isValid(dst) {
		return false
	}

	//Get the location between src and dst, and check if they're valid for this action
	jmp, valid := getLocationBetween(src, dst)

	//If not valid, return.
	if !valid {
		return false
	}

	//Get the relevant pieces
	srcPiece := getPiece(*checkers, src)
	jmpPiece := getPiece(*checkers, jmp)
	dstPiece := getPiece(*checkers, dst)

	srcIsKinged := isKinged(srcPiece)
	xDelta, yDelta := getDelta(src, dst)

	//Ensure each piece is the correct type
	if !isWhite(srcPiece) || !isRed(jmpPiece) || !isEmpty(dstPiece) {
		return false
	}

	//Check the vectors
	//If kinged, piece must move +/- 2 in the Y AND +/- 2 in the X.
	if srcIsKinged && (abs(yDelta) != 2 || abs(xDelta) != 2) {
		return false
	} else if !srcIsKinged && (abs(xDelta) != 2 || yDelta == -2) {
		//If not kinged, must move toward red side 2 squares and +/- 2 in the X
		return false
	}

	//If it passes the gauntlet, make the jump
	checkers.grid[dst.y][dst.x] = Piece{1}
	checkers.grid[jmp.y][jmp.x] = Piece{0}
	checkers.grid[src.y][src.x] = Piece{0}

	if dst.y == 7 {
		checkers.grid[dst.y][dst.x] = setKinged(checkers.grid[dst.y][dst.x])
	}

	return true
}