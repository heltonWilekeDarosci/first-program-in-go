package main

import "fmt"

func main() {
	name := "Wileke"
	version := 1.1
	fmt.Println("Hello Mr.", name)
	fmt.Println("This program is on version", version)

	fmt.Println("1- Start monitoring")
	fmt.Println("2- Exhibit Logs")
	fmt.Println("0- Exit")

	var command int
	fmt.Scan(&command)
}
