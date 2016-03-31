package main

func validateRedJump(checkers Checkers, src Location, jmp Location, dst Location) bool {
	if isValid(src) && isValid(jmp) && isValid(dst) {
		if isRed(getPiece(checkers, src)) && isWhite(getPiece(checkers, jmp)) && isEmpty(getPiece(checkers, dst)) {
			return true
		}
	}

	return false
}

func getRedForcedJumps(checkers Checkers) []Location {
	var possibleJumps []Location
	//Iterate through all the pieces on the board...
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			if isRed(checkers.grid[row][col]) {
				srcCheckLoc := Location{int8(row), int8(col)}
				piece := getPiece(checkers, srcCheckLoc)

				//Is this piece kinged?
				if isKinged(piece) {
					//Check the four adjacents

					//up-left
					jmpCheckLoc := Location{int8(col - 1), int8(row + 1)}
					dstCheckLoc := Location{int8(col - 2), int8(row + 2)}

					if validateRedJump(checkers, srcCheckLoc, jmpCheckLoc, dstCheckLoc) {
						possibleJumps = append(possibleJumps, srcCheckLoc)
						continue
					}

					//up-right
					jmpCheckLoc.x = int8(col + 1)
					dstCheckLoc.x = int8(col + 2)

					if validateRedJump(checkers, srcCheckLoc, jmpCheckLoc, dstCheckLoc) {
						possibleJumps = append(possibleJumps, srcCheckLoc)
						continue
					}
				}

				//Check the two downward adjacents
				//down-left

				jmpCheckLoc := Location{int8(col - 1), int8(row - 1)}
				dstCheckLoc := Location{int8(col - 2), int8(row - 2)}

				if validateRedJump(checkers, srcCheckLoc, jmpCheckLoc, dstCheckLoc) {
					possibleJumps = append(possibleJumps, srcCheckLoc)
					continue
				}

				jmpCheckLoc.x = int8(col + 1)
				dstCheckLoc.x = int8(col + 2)

				if validateRedJump(checkers, srcCheckLoc, jmpCheckLoc, dstCheckLoc) {
					possibleJumps = append(possibleJumps, srcCheckLoc)
					continue
					
				}
			}
		}
	}

	return possibleJumps
}

func redMove(checkers *Checkers, src Location, dst Location) bool {
	//Ensure it is red player's turn
	if checkers.whitePlayerTurn == true {
		return false
	}

	//Make sure the src and dst are within legal bounds
	if !isValid(src) || !isValid(dst) {
		return false
	}

	//Grab the pieces for convenient use
	srcPiece := checkers.grid[src.y][src.x]
	dstPiece := checkers.grid[dst.y][dst.x]

	srcIsKinged := isKinged(srcPiece)

	//Make sure the src and dst are legal for this context
	if !isRed(srcPiece) || !isEmpty(dstPiece) {
		return false
	}

	//Check the delta...
	xDelta, yDelta := getDelta(src, dst)

	legalX := (xDelta == 1 || xDelta == -1)
	legalY := (yDelta == 1 || yDelta == -1)

	//Now check the movement
	if srcIsKinged && (!legalX || !legalY) {
		return false
	} else if !legalX || yDelta != -1 {
		return false
	}

	//If it passed the gauntlet, the move is legal
	checkers.grid[dst.y][dst.x] = Piece{2}
	checkers.grid[src.y][src.x] = Piece{0}
	checkers.whitePlayerTurn = !checkers.whitePlayerTurn

	return true
}

func redJump(checkers *Checkers, src Location, dst Location) bool {

	//Must be white's turn to make a white jump
	if checkers.whitePlayerTurn == true {
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
	if !isRed(srcPiece) || !isWhite(jmpPiece) || !isEmpty(dstPiece) {
		return false
	}

	//Check the vectors
	//If kinged, piece must move +/- 2 in the Y AND +/- 2 in the X.
	if srcIsKinged && (abs(yDelta) != 2 || abs(xDelta) != 2) {
		return false
	} else if !srcIsKinged && (abs(xDelta) != 2 || yDelta == 2) {
		//If not kinged, must move toward red side 2 squares and +/- 2 in the X
		return false
	}

	//If it passes the gauntlet, make the jump
	checkers.grid[dst.y][dst.x] = Piece{2}
	checkers.grid[jmp.y][jmp.x] = Piece{0}
	checkers.grid[src.y][src.x] = Piece{0}

	return true
}