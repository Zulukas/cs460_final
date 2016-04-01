package main

import (
	"bufio"
	"fmt"
	"net"	
	"strings"
	"time"
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

func Write(conn net.Conn, msg string) {
	conn.Write([]byte(msg + "\n"))
}

func Read(conn net.Conn) string {
	reader := bufio.NewReader(conn)

	newmsg, _ := reader.ReadString('\n')
	newmsg = newmsg[:len(newmsg) - 1]	//Strip the new line			
		
	return newmsg
}

func acceptConnections(ln net.Listener, p1* net.Conn, p2* net.Conn) {
	white := false
	red := false

	for {
		if white == false {
			conn, _ := ln.Accept()
			conn.Write([]byte("White\n"))
			*p1 = conn
			white = true

			fmt.Println("White player has been registered")			
		} else if red == false {
			conn, _ := ln.Accept()
			conn.Write([]byte("Red\n"))
			*p2 = conn
			red = true

			fmt.Println("Red player has been registered")			
		} else {
			conn, _ := ln.Accept()
			conn.Write([]byte("Game is full :(\n"))			
		}		
	}
}

func main() {
	checkers := initCheckers()

	ln, _ := net.Listen("tcp", ":6789")

	var p1 net.Conn
	var p2 net.Conn

	go acceptConnections(ln, &p1, &p2)

	for (p1 == nil || p2 == nil) {
		time.Sleep(50 * time.Millisecond)
	}

	fmt.Println("White and Red players registered.")

	var goAgain bool
	var jumper Location

	errorMsg := ""
	error := false

	for {
		msg := ""

		/* TO PLAYER */

		if checkers.whitePlayerTurn {
			Write(p1, getGridString(checkers))

			if error {
				Write(p1, errorMsg)
				error = false
			} else {
				Write(p1, "...")
			}

			Write(p1, "White Player: ")
			msg = Read(p1)

			fmt.Println(msg)
		} else {
			Write(p2, getGridString(checkers))

			if error {
				Write(p2, errorMsg)
			} else {
				Write(p2, "...")
			}

			Write(p2, "Red Player: ")
			msg = Read(p2)
			fmt.Println(msg)
		}

		/* END TO PLAYER */

		/* HANDLE PLAYER INPUT */

		action, src, dst := interpretInput(msg)

		fmt.Println(action)
		display(src)
		fmt.Println("")
		display(dst)

		if action == "quit" {
			Write(p1, "quit")
			go p1.Close()

			Write(p2, "quit")
			go p2.Close()

			break
		}

		if goAgain == false {
			var jumps []Location

			if checkers.whitePlayerTurn {
				jumps = getWhiteForcedJumps(checkers)
				fmt.Println("White player has " + string(len(jumps)) + " jumps")
			} else {
				jumps = getRedForcedJumps(checkers)
				fmt.Println("Red player has " + string(len(jumps)) + " jumps")
			}

			if action == "invalid" {
				errorMsg = "Invalid action, please try again."
				error = true
			} else if action == "move" {
				if len(jumps) == 0 {
					if checkers.whitePlayerTurn {
						if !whiteMove(&checkers, src, dst) {
							errorMsg = "Invalid move, please try again."
							error = true
						}
					} else {
						if !redMove(&checkers, src, dst) {
							errorMsg = "Invalid move, please try again."
							error = true
						}
					}
				} else {
						errorMsg = "Must make a jump if a jump is available."
						error = true
				}
			} else if action == "jump" {
				if checkers.whitePlayerTurn {
					if whiteJump(&checkers, src, dst) {
						goAgain = true
						jumper = dst
					} else {
						errorMsg = "Invalid jump, please try again."
						error = true
					}
				} else {
					if redJump(&checkers, src, dst) {
						goAgain = true					
						jumper = dst
					} else {
						errorMsg = "Invalid jump, please try again."
						error = true
					}
				}
			}	
		} else {
			if jumperHasAnotherJump(checkers, jumper) {
				if action != "jump" {
					errorMsg =  "Must make another jump after jumping."
					error = true
				} else {
					if checkers.whitePlayerTurn {
						if whiteJump(&checkers, src, dst) {
							jumper = dst
						} else {
							errorMsg = "Invalid jump, please try again."
							error = true
						}
					} else {
						if redJump(&checkers, src, dst) {
							jumper = dst
						} else {
							errorMsg = "Invalid jump, please try again."
							error = true
						}
					}
				}
			} else {
				goAgain = false
				checkers.whitePlayerTurn = !checkers.whitePlayerTurn
			}
		}

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