package main

import "fmt"

func main() {

	showIntro()
	showMenu()
	command := seeCommand()

	switch command {
	case 1:
	case 2:
	case 0:
	default:
	}
}

func showIntro() {
	name := "Wileke"
	version := 1.1
	fmt.Println("Hello Mr.", name)
	fmt.Println("This program is on version", version)
}

func showMenu() {
	fmt.Println("1- Start monitoring")
	fmt.Println("2- Exhibit Logs")
	fmt.Println("0- Exit")
}

func seeCommand() int {
	var commandRead int
	fmt.Scan(&commandRead)

	return commandRead
}
