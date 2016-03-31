package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func interpretInput(input string) (string, Location, Location) {
	parts := strings.Split(input, " ")

	numParts := len(parts)

	action := ""

	if numParts == 3 {
		if parts[0] == "move" || parts[0] == "jump" {
			action = parts[0]
		}

		src := interpretLocation(parts[1])
		dst := interpretLocation(parts[2])

		// display(src)
		// display(dst)
		fmt.Println(src)
		fmt.Println(dst)

		fmt.Println("")

		return action, src, dst
	} else {
		if input == "quit" {
			action = "quit"
		} else {
			action = "invalid"
		}

		return action, Location{-1,-1}, Location{-1,-1}
	}
}

func handleWhitePlayer(checkers* Checkers) {
	reader := bufio.NewReader(os.Stdin)
	
	goAgain := false
	jumper := Location{int8(-1), int8(-1)}

	for {
		if goAgain == false {
			//Check to see if white has any available jumps...
			jumps := getWhiteForcedJumps(*checkers)
				
			fmt.Print("Num jumps: " + string(strconv.Itoa(len(jumps))) + " ")
			fmt.Print(jumps)
			fmt.Print("White player: ")

			input, _ := reader.ReadString('\n')
			action, src, dst := interpretInput(input)
			
			if len(jumps) == 0 {
				if action == "invalid" {
					continue
				} else if action == "move" {
					whiteMove(checkers, src, dst)
					return
				} else if action == "jump" {
					fmt.Println("No possible jumps available")
				}
			} else {
				if action == "invalid" {
					continue
				} else if action == "move" {
					fmt.Println("Unable to make a move if there are available jumps")
					continue
				} else if action == "jump" {
					whiteJump(checkers, src, dst)					
					jumper = dst
					goAgain = true
				}
			}
		} else {
			if jumperHasAnotherJump(*checkers, jumper) {
				fmt.Print("White player: ")
				input, _ := reader.ReadString('\n')
				action, src, dst := interpretInput(input)

				if action != "jump" {
					fmt.Println("Must make another jump after making a previous jump.")
					continue	
				}

				if jumper == src {
					whiteJump(checkers, src, dst)
					jumper = dst
				} else {
					fmt.Println("Must make another jump with the piece that previously jumped.")
				}
			} else {
				return
			}
		}

	}
}

func handleRedPlayer(checkers* Checkers) {
	reader := bufio.NewReader(os.Stdin)
	
	goAgain := false
	jumper := Location{int8(-1), int8(-1)}

	for {
		if goAgain == false {
			//Check to see if white has any available jumps...
			jumps := getRedForcedJumps(*checkers)
				
			fmt.Print("Num jumps: " + string(strconv.Itoa(len(jumps))) + " ")
			fmt.Print(jumps)
			fmt.Print("Red player: ")

			input, _ := reader.ReadString('\n')
			action, src, dst := interpretInput(input)
			
			if len(jumps) == 0 {
				if action == "invalid" {
					continue
				} else if action == "move" {
					redMove(checkers, src, dst)
					return
				} else if action == "jump" {
					fmt.Println("No possible jumps available")
				}
			} else {
				if action == "invalid" {
					continue
				} else if action == "move" {
					fmt.Println("Unable to make a move if there are available jumps")
					continue
				} else if action == "jump" {
					redJump(checkers, src, dst)
					goAgain = true
					jumper = dst
				}
			}
		} else {
			if jumperHasAnotherJump(*checkers, jumper) {
				fmt.Print("Red player: ")
				input, _ := reader.ReadString('\n')
				action, src, dst := interpretInput(input)

				if action != "jump" {
					fmt.Println("Must make another jump after making a previous jump.")
					continue	
				}

				if jumper == src {
					redJump(checkers, src, dst)
					jumper = dst
				} else {
					fmt.Println("Must make another jump with the piece that previously jumped.")
				}
			} else {
				return
			}
		}

	}
}

func main() {
	checkers := initCheckers()

	for {
		displayCheckers(checkers)


		if checkers.whitePlayerTurn {
			handleWhitePlayer(&checkers)
		} else {
			handleRedPlayer(&checkers)
		}

		// if checkers.whitePlayerTurn {
		// 	//Check to see if white has any available jumps...
		// 	jumps := getWhiteForcedJumps(checkers)

		// 	//Did white make a jump already this turn and has another jump?
		// 	if whiteAgain && !jumperHasAnotherJump(checkers, jumper) {
		// 		whiteAgain = false
		// 		checkers.whitePlayerTurn = false
		// 	}

		// 	fmt.Print("Num jumps: " + string(strconv.Itoa(len(jumps))) + " ")
		// 	fmt.Print(jumps)
		// 	fmt.Print("White player: ")
		// 	text, _ := reader.ReadString('\n')
		// 	action, src, dst := interpretInput(text)

		// 	if action == "invalid" {
		// 		continue
		// 	} else if action == "move" {
		// 		whiteMove(&checkers, src, dst)
		// 	} else if action == "jump" {
		// 		whiteJump(&checkers, src, dst)
		// 		whiteAgain = true
		// 		jumper = dst
		// 	}
		// } else {
		// 	jumps := getRedForcedJumps(checkers)

		// 	if redAgain && jumperHasAnotherJump(checkers, jumper) {
		// 		redAgain = false
		// 		checkers.whitePlayerTurn = true
		// 	}

		// 	fmt.Print("Num jumps: " + string(strconv.Itoa(len(jumps))) + " ")
		// 	fmt.Print("Red player: ")

		// 	text, _ := reader.ReadString('\n')
		// 	action, src, dst := interpretInput(text)

		// 	if action == "invalid" {
		// 		continue
		// 	} else if action == "move" {
		// 		redMove(&checkers, src, dst)
		// 	} else if action == "jump" {
		// 		redJump(&checkers, src, dst)
		// 		redAgain = true
		// 		jumper = dst
		// 	}
		// }
	}
}

/*
	Game logic is as follows:
		Player makes a move or a jump

		If a move, perform the move and the next player takes their turn

		If a jump, perform the jump and check to see if there are more jumps from 
		piece that just jumped
			Repeat forever until no more jumps are found with that particular piece.
*/