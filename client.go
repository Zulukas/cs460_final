package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func displayGrid(data string) {	
	pieces := strings.Split(data, ",")

	fmt.Println("    A   B   C   D   E   F   G   H")
	fmt.Println("  +---+---+---+---+---+---+---+---+")

	for row := 0; row < 8; row++ {
		fmt.Print(row + 1)
		fmt.Print(" ")

		for col := 0; col < 8; col++ {
			fmt.Print("| " + pieces[row * 8 + col] + " ")
		}

		fmt.Println("|")
		fmt.Println("  +---+---+---+---+---+---+---+---+")
	}
}

func main() { 
	//Connect to the server...
	conn, _ := net.Dial("tcp", "127.0.0.1:6789")

	sReader := bufio.NewReader(conn)
	cReader := bufio.NewReader(os.Stdin)

	message, _ := sReader.ReadString('\n')

	if message == "White\n" {
		fmt.Println("You are the white player.\n")
	} else if message == "Red\n" {
		fmt.Println("You are the red player.\n")
	} else {
		fmt.Println("The game is full, exiting.\n")
		return
	}

	for {
		//Get the grid
		message, _ = sReader.ReadString('\n')
		message = message[:len(message) - 1]
		displayGrid(message)

		//Get any error messages
		message, _ = sReader.ReadString('\n')
		message = message[:len(message) - 1]

		if message != "..." {
			fmt.Println(message)
		}

		//Get the prompt
		message, _ = sReader.ReadString('\n')
		message = message[:len(message) - 1]

		fmt.Println(message)
		
		//Get the player's action and send it.
		input, _ := cReader.ReadString('\n')
		conn.Write([]byte(input + "\n"))

		//Receive confirmation
	}	
}